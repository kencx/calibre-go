// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ABookSeries is an object representing the database table.
type ABookSeries struct {
	ID   null.Int64  `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Name string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Sort null.String `boil:"sort" json:"sort,omitempty" toml:"sort" yaml:"sort,omitempty"`
	Link string      `boil:"link" json:"link" toml:"link" yaml:"link"`

	R *aBookSeriesR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L aBookSeriesL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ABookSeriesColumns = struct {
	ID   string
	Name string
	Sort string
	Link string
}{
	ID:   "id",
	Name: "name",
	Sort: "sort",
	Link: "link",
}

var ABookSeriesTableColumns = struct {
	ID   string
	Name string
	Sort string
	Link string
}{
	ID:   "series.id",
	Name: "series.name",
	Sort: "series.sort",
	Link: "series.link",
}

// Generated where

var ABookSeriesWhere = struct {
	ID   whereHelpernull_Int64
	Name whereHelperstring
	Sort whereHelpernull_String
	Link whereHelperstring
}{
	ID:   whereHelpernull_Int64{field: "\"series\".\"id\""},
	Name: whereHelperstring{field: "\"series\".\"name\""},
	Sort: whereHelpernull_String{field: "\"series\".\"sort\""},
	Link: whereHelperstring{field: "\"series\".\"link\""},
}

// ABookSeriesRels is where relationship names are stored.
var ABookSeriesRels = struct {
}{}

// aBookSeriesR is where relationships are stored.
type aBookSeriesR struct {
}

// NewStruct creates a new relationship struct
func (*aBookSeriesR) NewStruct() *aBookSeriesR {
	return &aBookSeriesR{}
}

// aBookSeriesL is where Load methods for each relationship are stored.
type aBookSeriesL struct{}

var (
	aBookSeriesAllColumns            = []string{"id", "name", "sort", "link"}
	aBookSeriesColumnsWithoutDefault = []string{"name"}
	aBookSeriesColumnsWithDefault    = []string{"id", "sort", "link"}
	aBookSeriesPrimaryKeyColumns     = []string{"id"}
	aBookSeriesGeneratedColumns      = []string{"id"}
)

type (
	// ABookSeriesSlice is an alias for a slice of pointers to ABookSeries.
	// This should almost always be used instead of []ABookSeries.
	ABookSeriesSlice []*ABookSeries
	// ABookSeriesHook is the signature for custom ABookSeries hook methods
	ABookSeriesHook func(context.Context, boil.ContextExecutor, *ABookSeries) error

	aBookSeriesQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	aBookSeriesType                 = reflect.TypeOf(&ABookSeries{})
	aBookSeriesMapping              = queries.MakeStructMapping(aBookSeriesType)
	aBookSeriesPrimaryKeyMapping, _ = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, aBookSeriesPrimaryKeyColumns)
	aBookSeriesInsertCacheMut       sync.RWMutex
	aBookSeriesInsertCache          = make(map[string]insertCache)
	aBookSeriesUpdateCacheMut       sync.RWMutex
	aBookSeriesUpdateCache          = make(map[string]updateCache)
	aBookSeriesUpsertCacheMut       sync.RWMutex
	aBookSeriesUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var aBookSeriesAfterSelectMu sync.Mutex
var aBookSeriesAfterSelectHooks []ABookSeriesHook

var aBookSeriesBeforeInsertMu sync.Mutex
var aBookSeriesBeforeInsertHooks []ABookSeriesHook
var aBookSeriesAfterInsertMu sync.Mutex
var aBookSeriesAfterInsertHooks []ABookSeriesHook

var aBookSeriesBeforeUpdateMu sync.Mutex
var aBookSeriesBeforeUpdateHooks []ABookSeriesHook
var aBookSeriesAfterUpdateMu sync.Mutex
var aBookSeriesAfterUpdateHooks []ABookSeriesHook

var aBookSeriesBeforeDeleteMu sync.Mutex
var aBookSeriesBeforeDeleteHooks []ABookSeriesHook
var aBookSeriesAfterDeleteMu sync.Mutex
var aBookSeriesAfterDeleteHooks []ABookSeriesHook

var aBookSeriesBeforeUpsertMu sync.Mutex
var aBookSeriesBeforeUpsertHooks []ABookSeriesHook
var aBookSeriesAfterUpsertMu sync.Mutex
var aBookSeriesAfterUpsertHooks []ABookSeriesHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ABookSeries) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ABookSeries) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ABookSeries) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ABookSeries) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ABookSeries) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ABookSeries) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ABookSeries) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ABookSeries) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ABookSeries) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aBookSeriesAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddABookSeriesHook registers your hook function for all future operations.
func AddABookSeriesHook(hookPoint boil.HookPoint, aBookSeriesHook ABookSeriesHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		aBookSeriesAfterSelectMu.Lock()
		aBookSeriesAfterSelectHooks = append(aBookSeriesAfterSelectHooks, aBookSeriesHook)
		aBookSeriesAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		aBookSeriesBeforeInsertMu.Lock()
		aBookSeriesBeforeInsertHooks = append(aBookSeriesBeforeInsertHooks, aBookSeriesHook)
		aBookSeriesBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		aBookSeriesAfterInsertMu.Lock()
		aBookSeriesAfterInsertHooks = append(aBookSeriesAfterInsertHooks, aBookSeriesHook)
		aBookSeriesAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		aBookSeriesBeforeUpdateMu.Lock()
		aBookSeriesBeforeUpdateHooks = append(aBookSeriesBeforeUpdateHooks, aBookSeriesHook)
		aBookSeriesBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		aBookSeriesAfterUpdateMu.Lock()
		aBookSeriesAfterUpdateHooks = append(aBookSeriesAfterUpdateHooks, aBookSeriesHook)
		aBookSeriesAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		aBookSeriesBeforeDeleteMu.Lock()
		aBookSeriesBeforeDeleteHooks = append(aBookSeriesBeforeDeleteHooks, aBookSeriesHook)
		aBookSeriesBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		aBookSeriesAfterDeleteMu.Lock()
		aBookSeriesAfterDeleteHooks = append(aBookSeriesAfterDeleteHooks, aBookSeriesHook)
		aBookSeriesAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		aBookSeriesBeforeUpsertMu.Lock()
		aBookSeriesBeforeUpsertHooks = append(aBookSeriesBeforeUpsertHooks, aBookSeriesHook)
		aBookSeriesBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		aBookSeriesAfterUpsertMu.Lock()
		aBookSeriesAfterUpsertHooks = append(aBookSeriesAfterUpsertHooks, aBookSeriesHook)
		aBookSeriesAfterUpsertMu.Unlock()
	}
}

// OneG returns a single aBookSeries record from the query using the global executor.
func (q aBookSeriesQuery) OneG(ctx context.Context) (*ABookSeries, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single aBookSeries record from the query.
func (q aBookSeriesQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ABookSeries, error) {
	o := &ABookSeries{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for series")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ABookSeries records from the query using the global executor.
func (q aBookSeriesQuery) AllG(ctx context.Context) (ABookSeriesSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ABookSeries records from the query.
func (q aBookSeriesQuery) All(ctx context.Context, exec boil.ContextExecutor) (ABookSeriesSlice, error) {
	var o []*ABookSeries

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ABookSeries slice")
	}

	if len(aBookSeriesAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ABookSeries records in the query using the global executor
func (q aBookSeriesQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ABookSeries records in the query.
func (q aBookSeriesQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count series rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q aBookSeriesQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q aBookSeriesQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if series exists")
	}

	return count > 0, nil
}

// BookSeries retrieves all the records using an executor.
func BookSeries(mods ...qm.QueryMod) aBookSeriesQuery {
	mods = append(mods, qm.From("\"series\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"series\".*"})
	}

	return aBookSeriesQuery{q}
}

// FindABookSeriesG retrieves a single record by ID.
func FindABookSeriesG(ctx context.Context, iD null.Int64, selectCols ...string) (*ABookSeries, error) {
	return FindABookSeries(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindABookSeries retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindABookSeries(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*ABookSeries, error) {
	aBookSeriesObj := &ABookSeries{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"series\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, aBookSeriesObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from series")
	}

	if err = aBookSeriesObj.doAfterSelectHooks(ctx, exec); err != nil {
		return aBookSeriesObj, err
	}

	return aBookSeriesObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ABookSeries) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ABookSeries) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no series provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(aBookSeriesColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	aBookSeriesInsertCacheMut.RLock()
	cache, cached := aBookSeriesInsertCache[key]
	aBookSeriesInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			aBookSeriesAllColumns,
			aBookSeriesColumnsWithDefault,
			aBookSeriesColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, aBookSeriesGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"series\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"series\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into series")
	}

	if !cached {
		aBookSeriesInsertCacheMut.Lock()
		aBookSeriesInsertCache[key] = cache
		aBookSeriesInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ABookSeries record using the global executor.
// See Update for more documentation.
func (o *ABookSeries) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ABookSeries.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ABookSeries) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	aBookSeriesUpdateCacheMut.RLock()
	cache, cached := aBookSeriesUpdateCache[key]
	aBookSeriesUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			aBookSeriesAllColumns,
			aBookSeriesPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, aBookSeriesGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update series, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"series\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, aBookSeriesPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, append(wl, aBookSeriesPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update series row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for series")
	}

	if !cached {
		aBookSeriesUpdateCacheMut.Lock()
		aBookSeriesUpdateCache[key] = cache
		aBookSeriesUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q aBookSeriesQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q aBookSeriesQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for series")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for series")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ABookSeriesSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ABookSeriesSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aBookSeriesPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"series\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aBookSeriesPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in aBookSeries slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all aBookSeries")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ABookSeries) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ABookSeries) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no series provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(aBookSeriesColumnsWithDefault, o)

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

	aBookSeriesUpsertCacheMut.RLock()
	cache, cached := aBookSeriesUpsertCache[key]
	aBookSeriesUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			aBookSeriesAllColumns,
			aBookSeriesColumnsWithDefault,
			aBookSeriesColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			aBookSeriesAllColumns,
			aBookSeriesPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert series, could not build update column list")
		}

		ret := strmangle.SetComplement(aBookSeriesAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(aBookSeriesPrimaryKeyColumns))
			copy(conflict, aBookSeriesPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"series\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(aBookSeriesType, aBookSeriesMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert series")
	}

	if !cached {
		aBookSeriesUpsertCacheMut.Lock()
		aBookSeriesUpsertCache[key] = cache
		aBookSeriesUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ABookSeries record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ABookSeries) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ABookSeries record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ABookSeries) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ABookSeries provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), aBookSeriesPrimaryKeyMapping)
	sql := "DELETE FROM \"series\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from series")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for series")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q aBookSeriesQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q aBookSeriesQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no aBookSeriesQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from series")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for series")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ABookSeriesSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ABookSeriesSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(aBookSeriesBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aBookSeriesPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"series\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aBookSeriesPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from aBookSeries slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for series")
	}

	if len(aBookSeriesAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ABookSeries) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ABookSeries provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ABookSeries) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindABookSeries(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ABookSeriesSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ABookSeriesSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ABookSeriesSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ABookSeriesSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aBookSeriesPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"series\".* FROM \"series\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aBookSeriesPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ABookSeriesSlice")
	}

	*o = slice

	return nil
}

// ABookSeriesExistsG checks if the ABookSeries row exists.
func ABookSeriesExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return ABookSeriesExists(ctx, boil.GetContextDB(), iD)
}

// ABookSeriesExists checks if the ABookSeries row exists.
func ABookSeriesExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"series\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if series exists")
	}

	return exists, nil
}

// Exists checks if the ABookSeries row exists.
func (o *ABookSeries) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ABookSeriesExists(ctx, exec, o.ID)
}
