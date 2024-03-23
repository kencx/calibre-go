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

// Identifier is an object representing the database table.
type Identifier struct {
	ID   null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Book int64      `boil:"book" json:"book" toml:"book" yaml:"book"`
	Type string     `boil:"type" json:"type" toml:"type" yaml:"type"`
	Val  string     `boil:"val" json:"val" toml:"val" yaml:"val"`

	R *identifierR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L identifierL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IdentifierColumns = struct {
	ID   string
	Book string
	Type string
	Val  string
}{
	ID:   "id",
	Book: "book",
	Type: "type",
	Val:  "val",
}

var IdentifierTableColumns = struct {
	ID   string
	Book string
	Type string
	Val  string
}{
	ID:   "identifiers.id",
	Book: "identifiers.book",
	Type: "identifiers.type",
	Val:  "identifiers.val",
}

// Generated where

var IdentifierWhere = struct {
	ID   whereHelpernull_Int64
	Book whereHelperint64
	Type whereHelperstring
	Val  whereHelperstring
}{
	ID:   whereHelpernull_Int64{field: "\"identifiers\".\"id\""},
	Book: whereHelperint64{field: "\"identifiers\".\"book\""},
	Type: whereHelperstring{field: "\"identifiers\".\"type\""},
	Val:  whereHelperstring{field: "\"identifiers\".\"val\""},
}

// IdentifierRels is where relationship names are stored.
var IdentifierRels = struct {
}{}

// identifierR is where relationships are stored.
type identifierR struct {
}

// NewStruct creates a new relationship struct
func (*identifierR) NewStruct() *identifierR {
	return &identifierR{}
}

// identifierL is where Load methods for each relationship are stored.
type identifierL struct{}

var (
	identifierAllColumns            = []string{"id", "book", "type", "val"}
	identifierColumnsWithoutDefault = []string{"book", "val"}
	identifierColumnsWithDefault    = []string{"id", "type"}
	identifierPrimaryKeyColumns     = []string{"id"}
	identifierGeneratedColumns      = []string{"id"}
)

type (
	// IdentifierSlice is an alias for a slice of pointers to Identifier.
	// This should almost always be used instead of []Identifier.
	IdentifierSlice []*Identifier
	// IdentifierHook is the signature for custom Identifier hook methods
	IdentifierHook func(context.Context, boil.ContextExecutor, *Identifier) error

	identifierQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	identifierType                 = reflect.TypeOf(&Identifier{})
	identifierMapping              = queries.MakeStructMapping(identifierType)
	identifierPrimaryKeyMapping, _ = queries.BindMapping(identifierType, identifierMapping, identifierPrimaryKeyColumns)
	identifierInsertCacheMut       sync.RWMutex
	identifierInsertCache          = make(map[string]insertCache)
	identifierUpdateCacheMut       sync.RWMutex
	identifierUpdateCache          = make(map[string]updateCache)
	identifierUpsertCacheMut       sync.RWMutex
	identifierUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var identifierAfterSelectMu sync.Mutex
var identifierAfterSelectHooks []IdentifierHook

var identifierBeforeInsertMu sync.Mutex
var identifierBeforeInsertHooks []IdentifierHook
var identifierAfterInsertMu sync.Mutex
var identifierAfterInsertHooks []IdentifierHook

var identifierBeforeUpdateMu sync.Mutex
var identifierBeforeUpdateHooks []IdentifierHook
var identifierAfterUpdateMu sync.Mutex
var identifierAfterUpdateHooks []IdentifierHook

var identifierBeforeDeleteMu sync.Mutex
var identifierBeforeDeleteHooks []IdentifierHook
var identifierAfterDeleteMu sync.Mutex
var identifierAfterDeleteHooks []IdentifierHook

var identifierBeforeUpsertMu sync.Mutex
var identifierBeforeUpsertHooks []IdentifierHook
var identifierAfterUpsertMu sync.Mutex
var identifierAfterUpsertHooks []IdentifierHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Identifier) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Identifier) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Identifier) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Identifier) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Identifier) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Identifier) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Identifier) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Identifier) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Identifier) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range identifierAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIdentifierHook registers your hook function for all future operations.
func AddIdentifierHook(hookPoint boil.HookPoint, identifierHook IdentifierHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		identifierAfterSelectMu.Lock()
		identifierAfterSelectHooks = append(identifierAfterSelectHooks, identifierHook)
		identifierAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		identifierBeforeInsertMu.Lock()
		identifierBeforeInsertHooks = append(identifierBeforeInsertHooks, identifierHook)
		identifierBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		identifierAfterInsertMu.Lock()
		identifierAfterInsertHooks = append(identifierAfterInsertHooks, identifierHook)
		identifierAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		identifierBeforeUpdateMu.Lock()
		identifierBeforeUpdateHooks = append(identifierBeforeUpdateHooks, identifierHook)
		identifierBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		identifierAfterUpdateMu.Lock()
		identifierAfterUpdateHooks = append(identifierAfterUpdateHooks, identifierHook)
		identifierAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		identifierBeforeDeleteMu.Lock()
		identifierBeforeDeleteHooks = append(identifierBeforeDeleteHooks, identifierHook)
		identifierBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		identifierAfterDeleteMu.Lock()
		identifierAfterDeleteHooks = append(identifierAfterDeleteHooks, identifierHook)
		identifierAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		identifierBeforeUpsertMu.Lock()
		identifierBeforeUpsertHooks = append(identifierBeforeUpsertHooks, identifierHook)
		identifierBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		identifierAfterUpsertMu.Lock()
		identifierAfterUpsertHooks = append(identifierAfterUpsertHooks, identifierHook)
		identifierAfterUpsertMu.Unlock()
	}
}

// OneG returns a single identifier record from the query using the global executor.
func (q identifierQuery) OneG(ctx context.Context) (*Identifier, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single identifier record from the query.
func (q identifierQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Identifier, error) {
	o := &Identifier{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for identifiers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Identifier records from the query using the global executor.
func (q identifierQuery) AllG(ctx context.Context) (IdentifierSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Identifier records from the query.
func (q identifierQuery) All(ctx context.Context, exec boil.ContextExecutor) (IdentifierSlice, error) {
	var o []*Identifier

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Identifier slice")
	}

	if len(identifierAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Identifier records in the query using the global executor
func (q identifierQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Identifier records in the query.
func (q identifierQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count identifiers rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q identifierQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q identifierQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if identifiers exists")
	}

	return count > 0, nil
}

// Identifiers retrieves all the records using an executor.
func Identifiers(mods ...qm.QueryMod) identifierQuery {
	mods = append(mods, qm.From("\"identifiers\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"identifiers\".*"})
	}

	return identifierQuery{q}
}

// FindIdentifierG retrieves a single record by ID.
func FindIdentifierG(ctx context.Context, iD null.Int64, selectCols ...string) (*Identifier, error) {
	return FindIdentifier(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindIdentifier retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIdentifier(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Identifier, error) {
	identifierObj := &Identifier{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"identifiers\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, identifierObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from identifiers")
	}

	if err = identifierObj.doAfterSelectHooks(ctx, exec); err != nil {
		return identifierObj, err
	}

	return identifierObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Identifier) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Identifier) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no identifiers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(identifierColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	identifierInsertCacheMut.RLock()
	cache, cached := identifierInsertCache[key]
	identifierInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			identifierAllColumns,
			identifierColumnsWithDefault,
			identifierColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, identifierGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(identifierType, identifierMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(identifierType, identifierMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"identifiers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"identifiers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into identifiers")
	}

	if !cached {
		identifierInsertCacheMut.Lock()
		identifierInsertCache[key] = cache
		identifierInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Identifier record using the global executor.
// See Update for more documentation.
func (o *Identifier) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Identifier.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Identifier) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	identifierUpdateCacheMut.RLock()
	cache, cached := identifierUpdateCache[key]
	identifierUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			identifierAllColumns,
			identifierPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, identifierGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update identifiers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"identifiers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, identifierPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(identifierType, identifierMapping, append(wl, identifierPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update identifiers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for identifiers")
	}

	if !cached {
		identifierUpdateCacheMut.Lock()
		identifierUpdateCache[key] = cache
		identifierUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q identifierQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q identifierQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for identifiers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for identifiers")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o IdentifierSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IdentifierSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), identifierPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"identifiers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, identifierPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in identifier slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all identifier")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Identifier) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Identifier) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no identifiers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(identifierColumnsWithDefault, o)

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

	identifierUpsertCacheMut.RLock()
	cache, cached := identifierUpsertCache[key]
	identifierUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			identifierAllColumns,
			identifierColumnsWithDefault,
			identifierColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			identifierAllColumns,
			identifierPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert identifiers, could not build update column list")
		}

		ret := strmangle.SetComplement(identifierAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(identifierPrimaryKeyColumns))
			copy(conflict, identifierPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"identifiers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(identifierType, identifierMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(identifierType, identifierMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert identifiers")
	}

	if !cached {
		identifierUpsertCacheMut.Lock()
		identifierUpsertCache[key] = cache
		identifierUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Identifier record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Identifier) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Identifier record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Identifier) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Identifier provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), identifierPrimaryKeyMapping)
	sql := "DELETE FROM \"identifiers\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from identifiers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for identifiers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q identifierQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q identifierQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no identifierQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from identifiers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for identifiers")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o IdentifierSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IdentifierSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(identifierBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), identifierPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"identifiers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, identifierPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from identifier slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for identifiers")
	}

	if len(identifierAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Identifier) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Identifier provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Identifier) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIdentifier(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IdentifierSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty IdentifierSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IdentifierSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IdentifierSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), identifierPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"identifiers\".* FROM \"identifiers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, identifierPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IdentifierSlice")
	}

	*o = slice

	return nil
}

// IdentifierExistsG checks if the Identifier row exists.
func IdentifierExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return IdentifierExists(ctx, boil.GetContextDB(), iD)
}

// IdentifierExists checks if the Identifier row exists.
func IdentifierExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"identifiers\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if identifiers exists")
	}

	return exists, nil
}

// Exists checks if the Identifier row exists.
func (o *Identifier) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return IdentifierExists(ctx, exec, o.ID)
}
