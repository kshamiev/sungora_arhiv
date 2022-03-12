// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mdsungora

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"sungora/lib/typ"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Minio is an object representing the database table.
type Minio struct { // ИД
	ID typ.UUID `boil:"id" db:"id" json:"id" toml:"id" yaml:"id" example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
	// папка хранения - тип объекта
	Bucket string `boil:"bucket" db:"bucket" json:"bucket" toml:"bucket" yaml:"bucket"`
	// файл хранения - ид объекта
	ObjectID typ.UUID `boil:"object_id" db:"object_id" json:"object_id" toml:"object_id" yaml:"object_id" example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
	// имя файла
	Name string `boil:"name" db:"name" json:"name" toml:"name" yaml:"name"`
	// тип файла
	FileType string `boil:"file_type" db:"file_type" json:"file_type" toml:"file_type" yaml:"file_type"`
	// размер файла
	FileSize int `boil:"file_size" db:"file_size" json:"file_size" toml:"file_size" yaml:"file_size"`
	// дополнительные параметры файла
	Label null.JSON `boil:"label" db:"label" json:"label,omitempty" toml:"label" yaml:"label,omitempty" swaggertype:"string" example:"JSON"`
	// пользователь
	UserLogin string `boil:"user_login" db:"user_login" json:"user_login" toml:"user_login" yaml:"user_login"`
	// дата и время создания
	CreatedAt time.Time `boil:"created_at" db:"created_at" json:"created_at" toml:"created_at" yaml:"created_at" example:"2006-01-02T15:04:05Z"`
	// подтверждение загрузки
	IsConfirm bool `boil:"is_confirm" db:"is_confirm" json:"is_confirm" toml:"is_confirm" yaml:"is_confirm"`

	R *minioR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L minioL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MinioColumns = struct {
	ID        string
	Bucket    string
	ObjectID  string
	Name      string
	FileType  string
	FileSize  string
	Label     string
	UserLogin string
	CreatedAt string
	IsConfirm string
}{
	ID:        "id",
	Bucket:    "bucket",
	ObjectID:  "object_id",
	Name:      "name",
	FileType:  "file_type",
	FileSize:  "file_size",
	Label:     "label",
	UserLogin: "user_login",
	CreatedAt: "created_at",
	IsConfirm: "is_confirm",
}

var MinioTableColumns = struct {
	ID        string
	Bucket    string
	ObjectID  string
	Name      string
	FileType  string
	FileSize  string
	Label     string
	UserLogin string
	CreatedAt string
	IsConfirm string
}{
	ID:        "minio.id",
	Bucket:    "minio.bucket",
	ObjectID:  "minio.object_id",
	Name:      "minio.name",
	FileType:  "minio.file_type",
	FileSize:  "minio.file_size",
	Label:     "minio.label",
	UserLogin: "minio.user_login",
	CreatedAt: "minio.created_at",
	IsConfirm: "minio.is_confirm",
}

// Generated where

type whereHelpertyp_UUID struct{ field string }

func (w whereHelpertyp_UUID) EQ(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertyp_UUID) NEQ(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertyp_UUID) LT(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertyp_UUID) LTE(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertyp_UUID) GT(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertyp_UUID) GTE(x typ.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MinioWhere = struct {
	ID        whereHelpertyp_UUID
	Bucket    whereHelperstring
	ObjectID  whereHelpertyp_UUID
	Name      whereHelperstring
	FileType  whereHelperstring
	FileSize  whereHelperint
	Label     whereHelpernull_JSON
	UserLogin whereHelperstring
	CreatedAt whereHelpertime_Time
	IsConfirm whereHelperbool
}{
	ID:        whereHelpertyp_UUID{field: "\"minio\".\"id\""},
	Bucket:    whereHelperstring{field: "\"minio\".\"bucket\""},
	ObjectID:  whereHelpertyp_UUID{field: "\"minio\".\"object_id\""},
	Name:      whereHelperstring{field: "\"minio\".\"name\""},
	FileType:  whereHelperstring{field: "\"minio\".\"file_type\""},
	FileSize:  whereHelperint{field: "\"minio\".\"file_size\""},
	Label:     whereHelpernull_JSON{field: "\"minio\".\"label\""},
	UserLogin: whereHelperstring{field: "\"minio\".\"user_login\""},
	CreatedAt: whereHelpertime_Time{field: "\"minio\".\"created_at\""},
	IsConfirm: whereHelperbool{field: "\"minio\".\"is_confirm\""},
}

// MinioRels is where relationship names are stored.
var MinioRels = struct {
}{}

// minioR is where relationships are stored.
type minioR struct {
}

// NewStruct creates a new relationship struct
func (*minioR) NewStruct() *minioR {
	return &minioR{}
}

// minioL is where Load methods for each relationship are stored.
type minioL struct{}

var (
	minioAllColumns            = []string{"id", "bucket", "object_id", "name", "file_type", "file_size", "label", "user_login", "created_at", "is_confirm"}
	minioColumnsWithoutDefault = []string{"bucket", "object_id", "name", "file_type", "user_login"}
	minioColumnsWithDefault    = []string{"id", "file_size", "label", "created_at", "is_confirm"}
	minioPrimaryKeyColumns     = []string{"id"}
	minioGeneratedColumns      = []string{}
)

type (
	// MinioSlice is an alias for a slice of pointers to Minio.
	// This should almost always be used instead of []Minio.
	MinioSlice []*Minio
	// MinioHook is the signature for custom Minio hook methods
	MinioHook func(context.Context, boil.ContextExecutor, *Minio) error

	minioQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	minioType                 = reflect.TypeOf(&Minio{})
	minioMapping              = queries.MakeStructMapping(minioType)
	minioPrimaryKeyMapping, _ = queries.BindMapping(minioType, minioMapping, minioPrimaryKeyColumns)
	minioInsertCacheMut       sync.RWMutex
	minioInsertCache          = make(map[string]insertCache)
	minioUpdateCacheMut       sync.RWMutex
	minioUpdateCache          = make(map[string]updateCache)
	minioUpsertCacheMut       sync.RWMutex
	minioUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var minioAfterSelectHooks []MinioHook

var minioBeforeInsertHooks []MinioHook
var minioAfterInsertHooks []MinioHook

var minioBeforeUpdateHooks []MinioHook
var minioAfterUpdateHooks []MinioHook

var minioBeforeDeleteHooks []MinioHook
var minioAfterDeleteHooks []MinioHook

var minioBeforeUpsertHooks []MinioHook
var minioAfterUpsertHooks []MinioHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Minio) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Minio) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Minio) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Minio) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Minio) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Minio) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Minio) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Minio) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Minio) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range minioAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMinioHook registers your hook function for all future operations.
func AddMinioHook(hookPoint boil.HookPoint, minioHook MinioHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		minioAfterSelectHooks = append(minioAfterSelectHooks, minioHook)
	case boil.BeforeInsertHook:
		minioBeforeInsertHooks = append(minioBeforeInsertHooks, minioHook)
	case boil.AfterInsertHook:
		minioAfterInsertHooks = append(minioAfterInsertHooks, minioHook)
	case boil.BeforeUpdateHook:
		minioBeforeUpdateHooks = append(minioBeforeUpdateHooks, minioHook)
	case boil.AfterUpdateHook:
		minioAfterUpdateHooks = append(minioAfterUpdateHooks, minioHook)
	case boil.BeforeDeleteHook:
		minioBeforeDeleteHooks = append(minioBeforeDeleteHooks, minioHook)
	case boil.AfterDeleteHook:
		minioAfterDeleteHooks = append(minioAfterDeleteHooks, minioHook)
	case boil.BeforeUpsertHook:
		minioBeforeUpsertHooks = append(minioBeforeUpsertHooks, minioHook)
	case boil.AfterUpsertHook:
		minioAfterUpsertHooks = append(minioAfterUpsertHooks, minioHook)
	}
}

// One returns a single minio record from the query.
func (q minioQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Minio, error) {
	o := &Minio{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mdsungora: failed to execute a one query for minio")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Minio records from the query.
func (q minioQuery) All(ctx context.Context, exec boil.ContextExecutor) (MinioSlice, error) {
	var o []*Minio

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mdsungora: failed to assign all query results to Minio slice")
	}

	if len(minioAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Minio records in the query.
func (q minioQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: failed to count minio rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q minioQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mdsungora: failed to check if minio exists")
	}

	return count > 0, nil
}

// Minios retrieves all the records using an executor.
func Minios(mods ...qm.QueryMod) minioQuery {
	mods = append(mods, qm.From("\"minio\""))
	return minioQuery{NewQuery(mods...)}
}

// FindMinio retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMinio(ctx context.Context, exec boil.ContextExecutor, iD typ.UUID, selectCols ...string) (*Minio, error) {
	minioObj := &Minio{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"minio\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, minioObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mdsungora: unable to select from minio")
	}

	if err = minioObj.doAfterSelectHooks(ctx, exec); err != nil {
		return minioObj, err
	}

	return minioObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Minio) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mdsungora: no minio provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(minioColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	minioInsertCacheMut.RLock()
	cache, cached := minioInsertCache[key]
	minioInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			minioAllColumns,
			minioColumnsWithDefault,
			minioColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(minioType, minioMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(minioType, minioMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"minio\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"minio\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "mdsungora: unable to insert into minio")
	}

	if !cached {
		minioInsertCacheMut.Lock()
		minioInsertCache[key] = cache
		minioInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Minio.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Minio) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	minioUpdateCacheMut.RLock()
	cache, cached := minioUpdateCache[key]
	minioUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			minioAllColumns,
			minioPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("mdsungora: unable to update minio, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"minio\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, minioPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(minioType, minioMapping, append(wl, minioPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to update minio row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: failed to get rows affected by update for minio")
	}

	if !cached {
		minioUpdateCacheMut.Lock()
		minioUpdateCache[key] = cache
		minioUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q minioQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to update all for minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to retrieve rows affected for minio")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MinioSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("mdsungora: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), minioPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"minio\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, minioPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to update all in minio slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to retrieve rows affected all in update all minio")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Minio) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mdsungora: no minio provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(minioColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	minioUpsertCacheMut.RLock()
	cache, cached := minioUpsertCache[key]
	minioUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			minioAllColumns,
			minioColumnsWithDefault,
			minioColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			minioAllColumns,
			minioPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("mdsungora: unable to upsert minio, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(minioPrimaryKeyColumns))
			copy(conflict, minioPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"minio\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(minioType, minioMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(minioType, minioMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "mdsungora: unable to upsert minio")
	}

	if !cached {
		minioUpsertCacheMut.Lock()
		minioUpsertCache[key] = cache
		minioUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Minio record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Minio) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mdsungora: no Minio provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), minioPrimaryKeyMapping)
	sql := "DELETE FROM \"minio\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to delete from minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: failed to get rows affected by delete for minio")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q minioQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mdsungora: no minioQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to delete all from minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: failed to get rows affected by deleteall for minio")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MinioSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(minioBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), minioPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"minio\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, minioPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: unable to delete all from minio slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mdsungora: failed to get rows affected by deleteall for minio")
	}

	if len(minioAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Minio) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMinio(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MinioSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MinioSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), minioPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"minio\".* FROM \"minio\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, minioPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mdsungora: unable to reload all in MinioSlice")
	}

	*o = slice

	return nil
}

// MinioExists checks if the Minio row exists.
func MinioExists(ctx context.Context, exec boil.ContextExecutor, iD typ.UUID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"minio\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mdsungora: unable to check if minio exists")
	}

	return exists, nil
}
