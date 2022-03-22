package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"sungora/lib/enum"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/typ"

	"github.com/go-chi/chi"
)

// Response Сопровождение входящего запроса и ответ на него
type Response struct {
	Request  *http.Request
	response http.ResponseWriter
	lg       logger.Logger
}

// New Функционал по работе с входящим запросом для формирования ответа
func New(r *http.Request, w http.ResponseWriter) *Response {
	var rw = &Response{
		response: w,
		lg:       logger.Get(r.Context()),
		Request:  r,
	}
	return rw
}

// CookieGet Получение куки.
func (rw *Response) CookieGet(name string) (c string) {
	for _, cookie := range rw.Request.Cookies() {
		if cookie.Name == name {
			c = cookie.Value
		}
	}
	if c == "" {
		for n, h := range rw.Request.Header {
			if strings.EqualFold(n, name) {
				c = h[0]
			}
		}
	}
	return c
}

// CookieSet Установка печенек. Если время не указано кука сессионная (пока открыт браузер).
func (rw *Response) CookieSet(name, value string, path []string, t ...time.Time) {
	for i := range path {
		var cookie = new(http.Cookie)
		cookie.HttpOnly = true
		cookie.Name = name
		cookie.Domain = strings.Split(rw.Request.Host, ":")[0]
		cookie.Path = path[i]
		if len(t) > 0 {
			cookie.Expires = t[0]
		}
		cookie.Value = value
		http.SetCookie(rw.response, cookie)
	}
}

// CookieRem Удаление печенек.
func (rw *Response) CookieRem(name string, path []string) {
	for i := range path {
		var cookie = new(http.Cookie)
		cookie.Name = name
		cookie.Domain = strings.Split(rw.Request.Host, ":")[0]
		cookie.Path = path[i]
		http.SetCookie(rw.response, cookie)
	}
}

// JSONBodyDecode декодирование полученного тела запроса в формате json в объект
func (rw *Response) JSONBodyDecode(object interface{}) error {
	return JSONBodyDecode(rw.Request, object)
}

// JSONBodyDecode декодирование полученного тела запроса в формате json в объект
func JSONBodyDecode(r *http.Request, object interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	if len(body) == 0 {
		return errs.NewBadRequest(errors.New("the request body is empty"))
	}
	err = json.Unmarshal(body, object)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	return nil
}

// JSON ответ в формате json
func (rw *Response) JSON(object interface{}) {
	var data []byte
	status := http.StatusOK
	switch tt := object.(type) {
	default:
		var err error
		data, err = json.Marshal(object)
		if err != nil {
			data = []byte(err.Error())
			status = http.StatusBadRequest
		}
	case string:
		data = []byte(`"` + tt + `"`)
	case int, int8, int16, int32, int64:
		data = []byte(`"` + strconv.Itoa(tt.(int)) + `"`)
	case Error:
		rw.lg.Error(tt.Error())
		for _, t := range tt.Trace() {
			rw.lg.Trace(t)
		}
		object = Data{
			Message: tt.Response(),
		}
		data, _ = json.Marshal(object)
		status = tt.HTTPCode()
	case error:
		rw.lg.Error(tt.Error())
		object = Data{
			Message: tt.Error(),
		}
		data, _ = json.Marshal(object)
		status = http.StatusBadRequest
	}
	rw.response.Header().Set("Content-Length", strconv.Itoa(len(data)))
	rw.response.Header().Set(logger.TraceID, rw.Request.Context().Value(logger.CtxTraceID).(string))
	rw.response.Header().Set("Content-Type", "application/json")
	rw.response.WriteHeader(status)
	_, _ = rw.response.Write(data)
}

// Static ответ - отдача статических данных
func (rw *Response) Static(fileName string) {
	fi, err := os.Stat(fileName)
	if err != nil {
		data := []byte(http.StatusText(http.StatusNotFound) + ": " + filepath.Base(fileName))
		rw.response.WriteHeader(http.StatusNotFound)
		_, _ = rw.response.Write(data)
		return
	}

	if fi.IsDir() {
		if rw.Request.URL.Path != "/" {
			fileName += string(os.PathSeparator)
		}

		fileName += IndexHtml

		if _, err = os.Stat(fileName); err != nil {
			data := []byte(http.StatusText(http.StatusNotFound) + ": " + filepath.Base(fileName))
			rw.response.WriteHeader(http.StatusNotFound)
			_, _ = rw.response.Write(data)
			return
		}
	}

	// content
	data, err := os.ReadFile(fileName)
	if err != nil {
		data = []byte(http.StatusText(http.StatusBadRequest) + ": " + filepath.Base(fileName))
		rw.response.WriteHeader(http.StatusBadRequest)
		_, _ = rw.response.Write(data)
		return
	}
	rw.generalHeaderSet(fileName, len(data), http.StatusOK)
	_, _ = rw.response.Write(data)
}

// Reader ответ
func (rw *Response) Reader(data io.Reader, fileName string, status int) {
	rw.generalHeaderSet(fileName, 0, status)
	_, _ = io.Copy(rw.response, data)
}

// Bytes ответ
func (rw *Response) Bytes(data []byte, fileName string, status int) {
	rw.generalHeaderSet(fileName, len(data), status)
	_, _ = rw.response.Write(data)
}

func (rw *Response) Redirect301(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusMovedPermanently)
}

func (rw *Response) Redirect302(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusFound)
}

// generalHeaderSet общие заголовки любого ответа
func (rw *Response) generalHeaderSet(fileName string, l, status int) {
	// размер и тип контента
	if l > 0 {
		rw.response.Header().Set("Content-Length", strconv.Itoa(l))
	}
	tp := `application/octet-stream`
	ext := strings.Split(fileName, ".")
	if m := mime.TypeByExtension("." + ext[len(ext)-1]); m != `` {
		tp = m
	}
	rw.response.Header().Set("Content-Type", tp)
	//
	if !strings.Contains(tp, `image`) &&
		!strings.Contains(tp, `text`) &&
		!strings.Contains(tp, `xml`) &&
		!strings.Contains(tp, `json`) {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(fileName))
	}
	// status
	rw.response.WriteHeader(status)
}

// UploadFiles загрузка файлов на сервер
// dir папка куда будут загружены все переданные в запросе файлы
func (rw *Response) UploadFiles(dir string) ([]string, error) {
	if err := os.MkdirAll(dir, 0o777); err != nil {
		return nil, errs.NewBadRequest(err, "ошибка создания хранилища")
	}

	mr, err := rw.Request.MultipartReader()
	if err != nil {
		return nil, errs.NewBadRequest(err, "ошибка получения информации о загрузке")
	}

	var result []string
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errs.NewBadRequest(err, "ошибка получения файла")
		}
		fileName := part.FileName()
		if fileName == "" {
			continue
		}

		l := strings.Split(fileName, ".")
		l[0] += time.Now().Format("_20060102150405")
		path := dir + "/" + strings.Join(l, ".")

		dst, err := os.Create(path)
		if err != nil {
			return nil, errs.NewBadRequest(err, "ошибка создания файла")
		}

		buffer := make([]byte, 100000)
		var read int
		for {
			n, err := part.Read(buffer)
			if err != nil && err != io.EOF {
				_ = dst.Close()
				_ = os.Remove(path)
				return nil, errs.NewBadRequest(err, "ошибка чтения файла")
			}
			if n == 0 {
				break
			}
			read += n
			if _, err = dst.Write(buffer[:n]); err != nil {
				_ = dst.Close()
				_ = os.Remove(path)
				return nil, errs.NewBadRequest(err, "ошибка записи в файл")
			}
		}
		_ = dst.Close()
		result = append(result, path)
	}
	if len(result) == 0 {
		return nil, errs.NewBadRequest(errors.New("request is empty"))
	}
	return result, nil
}

// UploadBuffer загрузка файлов на сервер
func (rw *Response) UploadBuffer() (fileData map[string]*bytes.Buffer, fileName []string, err error) {
	mr, err := rw.Request.MultipartReader()
	if err != nil {
		return nil, nil, errs.NewBadRequest(err, "ошибка получения информации о загрузке")
	}

	fileData = map[string]*bytes.Buffer{}
	fileName = make([]string, 0, 1)
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, errs.NewBadRequest(err, "ошибка получения файла")
		}
		fName := part.FileName()
		if fName == "" {
			continue
		}

		dst := &bytes.Buffer{}
		buffer := make([]byte, 100000)
		var read int
		for {
			n, err := part.Read(buffer)
			if err != nil && err != io.EOF {
				return nil, nil, errs.NewBadRequest(err, "ошибка чтения файла")
			}
			if n == 0 {
				break
			}
			read += n
			if _, err = dst.Write(buffer[:n]); err != nil {
				return nil, nil, errs.NewBadRequest(err, "ошибка записи в файл")
			}
		}
		fileData[fName] = dst
		fileName = append(fileName, fName)
	}
	if len(fileData) == 0 {
		return nil, nil, errs.NewBadRequest(errors.New("request is empty"))
	}
	return fileData, fileName, nil
}

func (rw *Response) GetUser() (*User, error) {
	us, ok := rw.Request.Context().Value(CtxUser).(*User)
	if !ok {
		return nil, errs.NewUnauthorized(errors.New("user is not context (middleware.Auth)"))
	}
	return us, nil
}

func (rw *Response) GetUserTest() (*User, error) {
	return &User{
		ID:    typ.UUIDNew(),
		Login: "test",
		Roles: []enum.Role{
			enum.Role_DEVELOP,
			enum.Role_ADMIN,
			enum.Role_MODERATOR,
		},
	}, nil
}

func (rw *Response) GetToken() (string, error) {
	token, ok := rw.Request.Context().Value(CtxToken).(string)
	if !ok {
		return "", errs.NewBadRequest(errors.New("token is not context (middleware.Auth)"))
	}
	return token, nil
}

func (rw *Response) GetUserAndID(r *http.Request) (*User, typ.UUID, error) {
	us, err := rw.GetUser()
	if err != nil {
		return nil, typ.UUID{}, err
	}
	ID, err := typ.UUIDParse(chi.URLParam(r, "id"))
	if err != nil {
		return nil, typ.UUID{}, err
	}
	return us, ID, nil
}

func (rw *Response) Access(roles ...enum.Role) bool {
	us, ok := rw.Request.Context().Value(CtxUser).(*User)
	if !ok {
		return false
	}
	for i := range us.Roles {
		for j := range roles {
			if us.Roles[i] == roles[j] {
				return true
			}
		}
	}
	return false
}
