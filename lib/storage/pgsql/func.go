package pgsql

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"sungora/lib/typ"
)

var sqlArgsSearch = regexp.MustCompile(`\$\d{1,2}`)

const sqlTempSeparator = "@@@"

func sqlIn(query string, args ...interface{}) (queryNew string, argsNew []interface{}, err error) {
	queryArgs := sqlArgsSearch.FindAllString(query, -1)
	checkCountArgs := make(map[string]bool)
	for i := range queryArgs {
		if _, ok := checkCountArgs[queryArgs[i]]; !ok {
			checkCountArgs[queryArgs[i]] = true
		}
	}
	if len(args) != len(checkCountArgs) {
		return "", nil, errors.New("args count invalid")
	}

	var index, indexShift, num int
	var replace = make(map[string]string)
	var argsRes = make([]interface{}, 0, len(args))
	var qu []string
	var ar []interface{}

	for i := range args {
		index++
		indexShift++

		switch arg := args[i].(type) {
		default:
			replace["$"+strconv.Itoa(index)] = sqlTempSeparator + strconv.Itoa(indexShift)
			argsRes = append(argsRes, arg)
		case []string:
			qu = make([]string, len(arg))
			ar = make([]interface{}, len(argsRes), len(argsRes)+len(arg))
			copy(ar, argsRes)
			for num = range arg {
				qu[num] = sqlTempSeparator + strconv.Itoa(indexShift+num)
				ar = append(ar, arg[num])
			}
			argsRes = ar
			replace["$"+strconv.Itoa(index)] = strings.Join(qu, ",")
			indexShift += num
		case []int:
			qu = make([]string, len(arg))
			ar = make([]interface{}, len(argsRes), len(argsRes)+len(arg))
			copy(ar, argsRes)
			for num = range arg {
				qu[num] = sqlTempSeparator + strconv.Itoa(indexShift+num)
				ar = append(ar, arg[num])
			}
			argsRes = ar
			replace["$"+strconv.Itoa(index)] = strings.Join(qu, ",")
			indexShift += num
		case []decimal.Decimal:
			qu = make([]string, len(arg))
			ar = make([]interface{}, len(argsRes), len(argsRes)+len(arg))
			copy(ar, argsRes)
			for num = range arg {
				qu[num] = sqlTempSeparator + strconv.Itoa(indexShift+num)
				ar = append(ar, arg[num])
			}
			argsRes = ar
			replace["$"+strconv.Itoa(index)] = strings.Join(qu, ",")
			indexShift += num
		case []typ.UUID:
			qu = make([]string, len(arg))
			ar = make([]interface{}, len(argsRes), len(argsRes)+len(arg))
			copy(ar, argsRes)
			for num = range arg {
				qu[num] = sqlTempSeparator + strconv.Itoa(indexShift+num)
				ar = append(ar, arg[num])
			}
			argsRes = ar
			replace["$"+strconv.Itoa(index)] = strings.Join(qu, ",")
			indexShift += num
		}
	}

	for i := range replace {
		query = strings.ReplaceAll(query, i, replace[i])
	}
	query = strings.ReplaceAll(query, sqlTempSeparator, "$")

	return query, argsRes, nil
}
