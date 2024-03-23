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

// Comment is an object representing the database table.
type Comment struct {
	ID   null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Book int64      `boil:"book" json:"book" toml:"book" yaml:"book"`
	Text string     `boil:"text" json:"text" toml:"text" yaml:"text"`

	R *commentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommentColumns = struct {
	ID   string
	Book string
	Text string
}{
	ID:   "id",
	Book: "book",
	Text: "text",
}

var CommentTableColumns = struct {
	ID   string
	Book string
	Text string
}{
	ID:   "comments.id",
	Book: "comments.book",
	Text: "comments.text",
}

// Generated where

var CommentWhere = struct {
	ID   whereHelpernull_Int64
	Book whereHelperint64
	Text whereHelperstring
}{
	ID:   whereHelpernull_Int64{field: "\"comments\".\"id\""},
	Book: whereHelperint64{field: "\"comments\".\"book\""},
	Text: whereHelperstring{field: "\"comments\".\"text\""},
}

// CommentRels is where relationship names are stored.
var CommentRels = struct {
}{}

// commentR is where relationships are stored.
type commentR struct {
}

// NewStruct creates a new relationship struct
func (*commentR) NewStruct() *commentR {
	return &commentR{}
}

// commentL is where Load methods for each relationship are stored.
type commentL struct{}

var (
	commentAllColumns            = []string{"id", "book", "text"}
	commentColumnsWithoutDefault = []string{"book", "text"}
	commentColumnsWithDefault    = []string{"id"}
	commentPrimaryKeyColumns     = []string{"id"}
	commentGeneratedColumns      = []string{"id"}
)

type (
	// CommentSlice is an alias for a slice of pointers to Comment.
	// This should almost always be used instead of []Comment.
	CommentSlice []*Comment
	// CommentHook is the signature for custom Comment hook methods
	CommentHook func(context.Context, boil.ContextExecutor, *Comment) error

	commentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commentType                 = reflect.TypeOf(&Comment{})
	commentMapping              = queries.MakeStructMapping(commentType)
	commentPrimaryKeyMapping, _ = queries.BindMapping(commentType, commentMapping, commentPrimaryKeyColumns)
	commentInsertCacheMut       sync.RWMutex
	commentInsertCache          = make(map[string]insertCache)
	commentUpdateCacheMut       sync.RWMutex
	commentUpdateCache          = make(map[string]updateCache)
	commentUpsertCacheMut       sync.RWMutex
	commentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var commentAfterSelectMu sync.Mutex
var commentAfterSelectHooks []CommentHook

var commentBeforeInsertMu sync.Mutex
var commentBeforeInsertHooks []CommentHook
var commentAfterInsertMu sync.Mutex
var commentAfterInsertHooks []CommentHook

var commentBeforeUpdateMu sync.Mutex
var commentBeforeUpdateHooks []CommentHook
var commentAfterUpdateMu sync.Mutex
var commentAfterUpdateHooks []CommentHook

var commentBeforeDeleteMu sync.Mutex
var commentBeforeDeleteHooks []CommentHook
var commentAfterDeleteMu sync.Mutex
var commentAfterDeleteHooks []CommentHook

var commentBeforeUpsertMu sync.Mutex
var commentBeforeUpsertHooks []CommentHook
var commentAfterUpsertMu sync.Mutex
var commentAfterUpsertHooks []CommentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Comment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Comment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Comment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Comment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Comment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Comment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Comment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Comment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Comment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCommentHook registers your hook function for all future operations.
func AddCommentHook(hookPoint boil.HookPoint, commentHook CommentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		commentAfterSelectMu.Lock()
		commentAfterSelectHooks = append(commentAfterSelectHooks, commentHook)
		commentAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		commentBeforeInsertMu.Lock()
		commentBeforeInsertHooks = append(commentBeforeInsertHooks, commentHook)
		commentBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		commentAfterInsertMu.Lock()
		commentAfterInsertHooks = append(commentAfterInsertHooks, commentHook)
		commentAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		commentBeforeUpdateMu.Lock()
		commentBeforeUpdateHooks = append(commentBeforeUpdateHooks, commentHook)
		commentBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		commentAfterUpdateMu.Lock()
		commentAfterUpdateHooks = append(commentAfterUpdateHooks, commentHook)
		commentAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		commentBeforeDeleteMu.Lock()
		commentBeforeDeleteHooks = append(commentBeforeDeleteHooks, commentHook)
		commentBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		commentAfterDeleteMu.Lock()
		commentAfterDeleteHooks = append(commentAfterDeleteHooks, commentHook)
		commentAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		commentBeforeUpsertMu.Lock()
		commentBeforeUpsertHooks = append(commentBeforeUpsertHooks, commentHook)
		commentBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		commentAfterUpsertMu.Lock()
		commentAfterUpsertHooks = append(commentAfterUpsertHooks, commentHook)
		commentAfterUpsertMu.Unlock()
	}
}

// OneG returns a single comment record from the query using the global executor.
func (q commentQuery) OneG(ctx context.Context) (*Comment, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single comment record from the query.
func (q commentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Comment, error) {
	o := &Comment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for comments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Comment records from the query using the global executor.
func (q commentQuery) AllG(ctx context.Context) (CommentSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Comment records from the query.
func (q commentQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommentSlice, error) {
	var o []*Comment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Comment slice")
	}

	if len(commentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Comment records in the query using the global executor
func (q commentQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Comment records in the query.
func (q commentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count comments rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q commentQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q commentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if comments exists")
	}

	return count > 0, nil
}

// Comments retrieves all the records using an executor.
func Comments(mods ...qm.QueryMod) commentQuery {
	mods = append(mods, qm.From("\"comments\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"comments\".*"})
	}

	return commentQuery{q}
}

// FindCommentG retrieves a single record by ID.
func FindCommentG(ctx context.Context, iD null.Int64, selectCols ...string) (*Comment, error) {
	return FindComment(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindComment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindComment(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Comment, error) {
	commentObj := &Comment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"comments\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, commentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from comments")
	}

	if err = commentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return commentObj, err
	}

	return commentObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Comment) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Comment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comments provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commentInsertCacheMut.RLock()
	cache, cached := commentInsertCache[key]
	commentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, commentGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commentType, commentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"comments\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"comments\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into comments")
	}

	if !cached {
		commentInsertCacheMut.Lock()
		commentInsertCache[key] = cache
		commentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Comment record using the global executor.
// See Update for more documentation.
func (o *Comment) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Comment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Comment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	commentUpdateCacheMut.RLock()
	cache, cached := commentUpdateCache[key]
	commentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, commentGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update comments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"comments\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, commentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, append(wl, commentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update comments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for comments")
	}

	if !cached {
		commentUpdateCacheMut.Lock()
		commentUpdateCache[key] = cache
		commentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q commentQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q commentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for comments")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CommentSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"comments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all comment")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Comment) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Comment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comments provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)

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

	commentUpsertCacheMut.RLock()
	cache, cached := commentUpsertCache[key]
	commentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert comments, could not build update column list")
		}

		ret := strmangle.SetComplement(commentAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commentPrimaryKeyColumns))
			copy(conflict, commentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"comments\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commentType, commentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert comments")
	}

	if !cached {
		commentUpsertCacheMut.Lock()
		commentUpsertCache[key] = cache
		commentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Comment record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Comment) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Comment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Comment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Comment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commentPrimaryKeyMapping)
	sql := "DELETE FROM \"comments\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for comments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q commentQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q commentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comments")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CommentSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(commentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"comments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comments")
	}

	if len(commentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Comment) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Comment provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Comment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindComment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommentSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CommentSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"comments\".* FROM \"comments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommentSlice")
	}

	*o = slice

	return nil
}

// CommentExistsG checks if the Comment row exists.
func CommentExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return CommentExists(ctx, boil.GetContextDB(), iD)
}

// CommentExists checks if the Comment row exists.
func CommentExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"comments\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if comments exists")
	}

	return exists, nil
}

// Exists checks if the Comment row exists.
func (o *Comment) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CommentExists(ctx, exec, o.ID)
}
