package protos

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func Generate3(dir, md string) {
	Tag(dir, md)
}

var typList []string

func Tag(dir, pkgName string) {
	fileList, err := os.ReadDir(dir + "/" + pkgName)
	if err != nil {
		log.Fatal(err)
	}
	// теги
	for i := range fileList {
		if _, ok := FileException[fileList[i].Name()]; ok {
			continue
		}
		tagFile(dir + "/" + pkgName + "/" + fileList[i].Name())
	}
	// типы для генерации в grpc
	data := make([]string, 0, len(typList))
	for i := range typList {
		data = append(data, "&"+pkgName+"."+typList[i]+"{},\n")
	}
	tpl := strings.ReplaceAll(tplConf, "PKGNAME", pkgName)
	tpl = strings.ReplaceAll(tpl, "PKGTYPES", strings.Join(data, ""))
	if err := os.WriteFile(dir+"/generate/protos/config_work.go", []byte(tpl), 0o600); err != nil {
		log.Fatal(err)
	}
}

var patternType = regexp.MustCompile("type (.+?) struct {")

func tagFile(filePath string) {
	d, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var flagType bool
	data := strings.Split(string(d), "\n")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
		tFind := patternType.FindStringSubmatch(data[i])
		if len(tFind) == 2 && tFind[1] == strings.Title(tFind[1]) {
			typList = append(typList, tFind[1])
			flagType = true
			continue
		}
		if flagType {
			if data[i] == "" || data[i] == "}" {
				flagType = false
				continue
			}
			data[i] = tagType(data[i])
		}
	}
	if err := os.WriteFile(filePath, []byte(strings.Join(data, "\n")), 0o600); err != nil {
		log.Fatal(err)
	}
}

var patternJson = regexp.MustCompile(` json:"(.+?)" `)

const (
	tagInt     = ` swaggertype:"number" example:"0"`
	tagDecimal = ` swaggertype:"number" example:"0.01"`
	tagString  = ` swaggertype:"string"`
	tagJson    = ` swaggertype:"string" example:"JSON"`
	tagBytes   = ` swaggertype:"string" example:"BYTES"`
	tagTime    = ` example:"2006-01-02T15:04:05Z"`
	tagUuid    = ` example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
)

// nolint gocyclo
func tagType(s string) string {
	switch {
	case strings.Contains(s, " null.Int64 ") || strings.Contains(s, " null.Int32 ") ||
		strings.Contains(s, " null.Int16 ") || strings.Contains(s, " null.Int8 ") ||
		strings.Contains(s, " null.Uint64 ") || strings.Contains(s, " null.Uint32 ") ||
		strings.Contains(s, " null.Uint16 ") || strings.Contains(s, " null.Uint8 ") ||
		strings.Contains(s, " null.Int ") || strings.Contains(s, " null.Uint "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagInt, s)
		if len(l) == 3 && !f {
			l[1] += tagInt
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " decimal.Decimal ") || strings.Contains(s, " decimal.NullDecimal ") ||
		strings.Contains(s, " null.Float32 ") || strings.Contains(s, " null.Float64 ") ||
		strings.Contains(s, " float32 ") || strings.Contains(s, " float64 "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagDecimal, s)
		if len(l) == 3 && !f {
			l[1] += tagDecimal
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " null.String "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagString, s)
		if len(l) == 3 && !f {
			l[1] += tagString
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " null.Bytes "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagBytes, s)
		if len(l) == 3 && !f {
			l[1] += tagBytes
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " null.JSON "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagJson, s)
		if len(l) == 3 && !f {
			l[1] += tagJson
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " time.Time ") || strings.Contains(s, " null.Time "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagTime, s)
		if len(l) == 3 && !f {
			l[1] += tagTime
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " uuid.UUID ") || strings.Contains(s, " typ.UUID "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagUuid, s)
		if len(l) == 3 && !f {
			l[1] += tagUuid
			s = strings.Join(l, "`")
		}
	case strings.Contains(s, " time.Duration "):
		l := strings.Split(s, "`")
		f, _ := regexp.MatchString(tagInt, s)
		if len(l) == 3 && !f {
			l[1] += tagInt
			s = strings.Join(l, "`")
		}
	}

	// from sqlx
	pFind := patternJson.FindStringSubmatch(s)
	if len(pFind) == 2 {
		db := ` db:"` + strings.Split(pFind[1], ",")[0] + `"`
		if !strings.Contains(s, db) {
			s = strings.ReplaceAll(s, pFind[0], db+pFind[0])
		}
	}

	return s
}

const tplConf = `
package protos

func init() {
	GenerateConfig["PKGNAME"] = []interface{}{
		PKGTYPES
	}
}
`
