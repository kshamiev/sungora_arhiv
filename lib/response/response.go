package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"sungora/lib/enum"

	"github.com/go-chi/chi"

	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/typ"
)

// Структура для работы с входящим запросом
type Response struct {
	Request  *http.Request
	response http.ResponseWriter
	lg       logger.Logger
}

// New Функционал по работе с входящим запросом для формирования ответа
func New(r *http.Request, w http.ResponseWriter) *Response {
	var rw = &Response{
		response: w,
		lg:       logger.GetLogger(r.Context()),
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

// CookiesSet Установка нескольких кук. Если время не указано кука сессионная (пока открыт браузер).
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

// CookiesRem Удаление нескольких кук.
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
	body, err := ioutil.ReadAll(r.Body)
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

// JSONError ответ об ошибке в формате json
func (rw *Response) JSONError(err error) {
	if e, ok := err.(Error); ok {
		rw.lg.Error(e.Error())
		for _, t := range e.Trace() {
			rw.lg.Trace(t)
		}
		response := &Data{
			Code:    rw.Request.Context().Value(CtxTraceID).(string),
			Message: e.Response(),
		}
		rw.JSON(response, e.HTTPCode())
	} else {
		rw.lg.Error(err.Error())
		response := &Data{
			Code:    rw.Request.Context().Value(CtxTraceID).(string),
			Message: err.Error(),
		}
		rw.JSON(response, http.StatusBadRequest)
	}
}

// JSON ответ в формате json
func (rw *Response) JSON(object interface{}, status ...int) {
	data, err := json.Marshal(object)
	if err != nil {
		e := errs.NewBadRequest(err)
		rw.lg.Error(e.Error())
		for _, t := range e.Trace() {
			rw.lg.Trace(t)
		}
		// Заголовки
		rw.generalHeaderSet("application/json; charset=utf-8", int64(len(data)), http.StatusBadRequest)
		// Тело документа
		_, _ = rw.response.Write([]byte(e.Response()))

		return
	}
	// Статус ответа
	if len(status) == 0 {
		status = append(status, http.StatusOK)
	}

	// Заголовки
	rw.generalHeaderSet("application/json; charset=utf-8", int64(len(data)), status[0])
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Static ответ - отдача статических данных
func (rw *Response) Static(pathFile string) {
	fi, err := os.Stat(pathFile)
	if err != nil {
		data := []byte(http.StatusText(http.StatusNotFound) + ": " + filepath.Base(pathFile))
		rw.response.WriteHeader(http.StatusNotFound)
		_, _ = rw.response.Write(data)
		return
	}

	if fi.IsDir() {
		if rw.Request.URL.Path != "/" {
			pathFile += string(os.PathSeparator)
		}

		pathFile += "index.html"

		if _, err = os.Stat(pathFile); err != nil {
			data := []byte(http.StatusText(http.StatusNotFound) + ": " + filepath.Base(pathFile))
			rw.response.WriteHeader(http.StatusNotFound)
			_, _ = rw.response.Write(data)
			return
		}
	}

	// content
	data, err := ioutil.ReadFile(pathFile)
	if err != nil {
		data = []byte(http.StatusText(http.StatusBadRequest) + ": " + filepath.Base(pathFile))
		rw.response.WriteHeader(http.StatusBadRequest)
		_, _ = rw.response.Write(data)
		return
	}
	// type
	var tp = `application/octet-stream`

	l := strings.Split(pathFile, ".")
	fileExt := `.` + l[len(l)-1]

	if mimeType := mime.TypeByExtension(fileExt); mimeType != `` {
		tp = mimeType
	}
	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(tp, `image`) == -1 && strings.LastIndex(tp, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(pathFile))
	}
	// Заголовки
	rw.generalHeaderSet(tp, int64(len(data)), http.StatusOK)
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Reader ответ
func (rw *Response) Reader(data io.Reader, dataLen int64, fileName, mimeType string, status int) {
	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(mimeType, `image`) == -1 && strings.LastIndex(mimeType, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+fileName)
	}
	// Заголовки
	rw.generalHeaderSet(mimeType, dataLen, status)
	// Тело документа
	_, _ = io.Copy(rw.response, data)
}

// Bytes ответ
func (rw *Response) Bytes(data []byte, fileName string) {
	l := strings.Split(fileName, ".")
	mimeType := mime.TypeByExtension("." + l[len(l)-1])

	// Аттач если документ не картинка и не текстововой
	if strings.LastIndex(mimeType, `image`) == -1 && strings.LastIndex(mimeType, `text`) == -1 {
		rw.response.Header().Set("Content-Disposition", "attachment; filename = "+filepath.Base(fileName))
	}
	// Заголовки
	rw.generalHeaderSet(mimeType, int64(len(data)), http.StatusOK)
	// Тело документа
	_, _ = rw.response.Write(data)
}

// Redirect 301
func (rw *Response) Redirect301(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusMovedPermanently)
}

// Redirect 302
func (rw *Response) Redirect302(redirectURL string) {
	rw.response.Header().Set("Location", redirectURL)
	rw.response.WriteHeader(http.StatusFound)
}

// generalHeaderSet общие заголовки любого ответа
func (rw *Response) generalHeaderSet(contentTyp string, l int64, status int) {
	t := time.Now()
	// запрет кеширования
	rw.response.Header().Set("Cache-Control", "no-cache, must-revalidate")
	rw.response.Header().Set("Pragma", "no-cache")
	rw.response.Header().Set("Date", t.Format(time.RFC3339))
	rw.response.Header().Set("Last-Modified", t.Format(time.RFC3339))
	// размер и тип контента
	rw.response.Header().Set("Content-Type", contentTyp)

	if l > 0 {
		rw.response.Header().Set("Content-Length", fmt.Sprintf("%d", l))
	}

	// status
	rw.response.WriteHeader(status)
}

// UploadFiles загрузка файлов на сервер
// dir папка куда будут загружены все переданные в запросе файлы
func (rw *Response) UploadFiles(dir string) ([]string, error) {
	if err := os.MkdirAll(dir, 0777); err != nil {
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

		formName := part.FormName()
		if formName == "" {
			continue
		}

		fileName := part.FileName()
		if fileName == "" {
			continue
		}

		var read int64

		l := strings.Split(fileName, ".")
		l[0] += time.Now().Format("_20060102150405")
		path := dir + "/" + strings.Join(l, ".")

		dst, err := os.Create(path)
		if err != nil {
			return nil, errs.NewBadRequest(err, "ошибка создания файла")
		}

		buffer := make([]byte, 100000)
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

			read += int64(n)

			_, err = dst.Write(buffer[:n])
			if err != nil {
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

func (rw *Response) GetUser() (*User, error) {
	us, ok := rw.Request.Context().Value(CtxUser).(*User)
	if !ok {
		return nil, errs.NewUnauthorized(errors.New("user is not context (middleware.Auth)"))
	}
	return us, nil
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
