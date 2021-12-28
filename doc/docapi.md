### Документирование api

Для работы со свагером мы используем библиотеку: [swaggo](https://github.com/swaggo/swag#api-operation)

Описание документирования api:
<pre>
//+funcName godoc
//+@Tags tagName                                                    группировка api запросов
//+@Summary Авторизация пользователя по логину и паролю (ldap).     пишем кратко о чем речь и что принимает на входе
// @Description Возвращается токен авторизации и пользователья      пишем что возвращает и возможно подробности
// @Accept json                                                     тип принимаемых данных
// @Produce json                                                    тип возвращаемых данных
// @Param name TARGET TYPE_VALUE true "com"                         входящие параметры
// @Success 200 {TYPE_RESPONSE} TYPE_VALUE "com"                    положительный ответ
// @Failure 400 {TYPE_RESPONSE} TYPE_VALUE "com"                    отрицательный ответ
// @Failure 401 {TYPE_RESPONSE} TYPE_VALUE "user unauthorized"      пользователь не авторизован
// @Security ApiKeyAuth                                             запрос авторизованный по ключу или токену
// @Router /page/page [post]                                        относительный роутинг от базового и метод
// @Deprecated
</pre>

<pre>
+ Обязательные теги и теги по контексту (параметров может и не быть...)
TARGET          = header | path | query  | body | formData
TYPE_VALUE      = string | int  | number | bool | file | pkg.CustomStruct
TYPE_RESPONSE   = string | int  | number | bool | file | object | array
</pre>

@Accept & @Produce

    son                     application/json
    xml                     text/xml
    plain                   text/plain
    html                    text/html
    mpfd                    multipart/form-data
    x-www-form-urlencoded   application/x-www-form-urlencoded
    json-api                application/vnd.api+json
    json-stream             application/x-json-stream
    octet-stream            application/octet-stream
    png                     image/png
    jpeg                    image/jpeg
    gif                     image/gif

Пример:
<pre>
// Login авторизация пользователя по логину и паролю ldap
// @Tags Auth
// @Summary авторизация пользователя по логину и паролю (ldap).
// @Description возвращается токен авторизации
// @Accept json                                                    
// @Produce json                                                   
// @Param credentials body models.Credentials true "реквизиты доступа"
// @Success 200 {string} string "успешная авторизация"
// @Failure 400 {object} request.Error "operation error"
// @Failure 401 {object} request.Error "unauthorized"
// @Failure 403 {object} request.Error "forbidden"
// @Failure 404 {object} request.Error "not found"
// @Security ApiKeyAuth
// @Router /auth/login [post]
</pre>

Проблемы:

* Не умеет работать с алиасами в импортах.
* Типы slice, map не поддерживаются для входных параметров (нужно оборачивать в отдельные типы)

Принятые коды ответов:

- 200 Любой положительный ответ
- 301 Редирект (перманентный). Переход на другой запрос
- 302 Редирект (от логики). Переход на другой запрос
- 400 Ошибка работы приложения
- 401 Пользователь не авторизован
- 403 Отказано в операции за отсутствием прав
- 404 Данные по запросу не найдены

Формат принимаемых и отдаваемых данных для API:

- Данные передаются в формате JSON
