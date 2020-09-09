// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Space is an object representing the database table.
type Space struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	SalonID   int64     `boil:"salon_id" json:"salon_id" toml:"salon_id" yaml:"salon_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *spaceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spaceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpaceColumns = struct {
	ID        string
	SalonID   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	SalonID:   "salon_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// Generated where

var SpaceWhere = struct {
	ID        whereHelperint64
	SalonID   whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint64{field: "`spaces`.`id`"},
	SalonID:   whereHelperint64{field: "`spaces`.`salon_id`"},
	CreatedAt: whereHelpertime_Time{field: "`spaces`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`spaces`.`updated_at`"},
	DeletedAt: whereHelpernull_Time{field: "`spaces`.`deleted_at`"},
}

// SpaceRels is where relationship names are stored.
var SpaceRels = struct {
	Salon        string
	Reservations string
}{
	Salon:        "Salon",
	Reservations: "Reservations",
}

// spaceR is where relationships are stored.
type spaceR struct {
	Salon        *Salon
	Reservations ReservationSlice
}

// NewStruct creates a new relationship struct
func (*spaceR) NewStruct() *spaceR {
	return &spaceR{}
}

// spaceL is where Load methods for each relationship are stored.
type spaceL struct{}

var (
	spaceAllColumns            = []string{"id", "salon_id", "created_at", "updated_at", "deleted_at"}
	spaceColumnsWithoutDefault = []string{"salon_id", "created_at", "updated_at", "deleted_at"}
	spaceColumnsWithDefault    = []string{"id"}
	spacePrimaryKeyColumns     = []string{"id"}
)

type (
	// SpaceSlice is an alias for a slice of pointers to Space.
	// This should generally be used opposed to []Space.
	SpaceSlice []*Space
	// SpaceHook is the signature for custom Space hook methods
	SpaceHook func(context.Context, boil.ContextExecutor, *Space) error

	spaceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spaceType                 = reflect.TypeOf(&Space{})
	spaceMapping              = queries.MakeStructMapping(spaceType)
	spacePrimaryKeyMapping, _ = queries.BindMapping(spaceType, spaceMapping, spacePrimaryKeyColumns)
	spaceInsertCacheMut       sync.RWMutex
	spaceInsertCache          = make(map[string]insertCache)
	spaceUpdateCacheMut       sync.RWMutex
	spaceUpdateCache          = make(map[string]updateCache)
	spaceUpsertCacheMut       sync.RWMutex
	spaceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var spaceBeforeInsertHooks []SpaceHook
var spaceBeforeUpdateHooks []SpaceHook
var spaceBeforeDeleteHooks []SpaceHook
var spaceBeforeUpsertHooks []SpaceHook

var spaceAfterInsertHooks []SpaceHook
var spaceAfterSelectHooks []SpaceHook
var spaceAfterUpdateHooks []SpaceHook
var spaceAfterDeleteHooks []SpaceHook
var spaceAfterUpsertHooks []SpaceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Space) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Space) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Space) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Space) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Space) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Space) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Space) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Space) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Space) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spaceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSpaceHook registers your hook function for all future operations.
func AddSpaceHook(hookPoint boil.HookPoint, spaceHook SpaceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		spaceBeforeInsertHooks = append(spaceBeforeInsertHooks, spaceHook)
	case boil.BeforeUpdateHook:
		spaceBeforeUpdateHooks = append(spaceBeforeUpdateHooks, spaceHook)
	case boil.BeforeDeleteHook:
		spaceBeforeDeleteHooks = append(spaceBeforeDeleteHooks, spaceHook)
	case boil.BeforeUpsertHook:
		spaceBeforeUpsertHooks = append(spaceBeforeUpsertHooks, spaceHook)
	case boil.AfterInsertHook:
		spaceAfterInsertHooks = append(spaceAfterInsertHooks, spaceHook)
	case boil.AfterSelectHook:
		spaceAfterSelectHooks = append(spaceAfterSelectHooks, spaceHook)
	case boil.AfterUpdateHook:
		spaceAfterUpdateHooks = append(spaceAfterUpdateHooks, spaceHook)
	case boil.AfterDeleteHook:
		spaceAfterDeleteHooks = append(spaceAfterDeleteHooks, spaceHook)
	case boil.AfterUpsertHook:
		spaceAfterUpsertHooks = append(spaceAfterUpsertHooks, spaceHook)
	}
}

// One returns a single space record from the query.
func (q spaceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Space, error) {
	o := &Space{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for spaces")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Space records from the query.
func (q spaceQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpaceSlice, error) {
	var o []*Space

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Space slice")
	}

	if len(spaceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Space records in the query.
func (q spaceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count spaces rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q spaceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if spaces exists")
	}

	return count > 0, nil
}

// Salon pointed to by the foreign key.
func (o *Space) Salon(mods ...qm.QueryMod) salonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.SalonID),
	}

	queryMods = append(queryMods, mods...)

	query := Salons(queryMods...)
	queries.SetFrom(query.Query, "`salons`")

	return query
}

// Reservations retrieves all the reservation's Reservations with an executor.
func (o *Space) Reservations(mods ...qm.QueryMod) reservationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`reservation`.`space_id`=?", o.ID),
	)

	query := Reservations(queryMods...)
	queries.SetFrom(query.Query, "`reservation`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`reservation`.*"})
	}

	return query
}

// LoadSalon allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spaceL) LoadSalon(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpace interface{}, mods queries.Applicator) error {
	var slice []*Space
	var object *Space

	if singular {
		object = maybeSpace.(*Space)
	} else {
		slice = *maybeSpace.(*[]*Space)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spaceR{}
		}
		args = append(args, object.SalonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spaceR{}
			}

			for _, a := range args {
				if a == obj.SalonID {
					continue Outer
				}
			}

			args = append(args, obj.SalonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`salons`), qm.WhereIn(`salons.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Salon")
	}

	var resultSlice []*Salon
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Salon")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for salons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for salons")
	}

	if len(spaceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Salon = foreign
		if foreign.R == nil {
			foreign.R = &salonR{}
		}
		foreign.R.Spaces = append(foreign.R.Spaces, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SalonID == foreign.ID {
				local.R.Salon = foreign
				if foreign.R == nil {
					foreign.R = &salonR{}
				}
				foreign.R.Spaces = append(foreign.R.Spaces, local)
				break
			}
		}
	}

	return nil
}

// LoadReservations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (spaceL) LoadReservations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpace interface{}, mods queries.Applicator) error {
	var slice []*Space
	var object *Space

	if singular {
		object = maybeSpace.(*Space)
	} else {
		slice = *maybeSpace.(*[]*Space)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spaceR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spaceR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`reservation`), qm.WhereIn(`reservation.space_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load reservation")
	}

	var resultSlice []*Reservation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice reservation")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on reservation")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reservation")
	}

	if len(reservationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Reservations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &reservationR{}
			}
			foreign.R.Space = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SpaceID {
				local.R.Reservations = append(local.R.Reservations, foreign)
				if foreign.R == nil {
					foreign.R = &reservationR{}
				}
				foreign.R.Space = local
				break
			}
		}
	}

	return nil
}

// SetSalon of the space to the related item.
// Sets o.R.Salon to related.
// Adds o to related.R.Spaces.
func (o *Space) SetSalon(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Salon) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `spaces` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"salon_id"}),
		strmangle.WhereClause("`", "`", 0, spacePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SalonID = related.ID
	if o.R == nil {
		o.R = &spaceR{
			Salon: related,
		}
	} else {
		o.R.Salon = related
	}

	if related.R == nil {
		related.R = &salonR{
			Spaces: SpaceSlice{o},
		}
	} else {
		related.R.Spaces = append(related.R.Spaces, o)
	}

	return nil
}

// AddReservations adds the given related objects to the existing relationships
// of the space, optionally inserting them as new records.
// Appends related to o.R.Reservations.
// Sets related.R.Space appropriately.
func (o *Space) AddReservations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Reservation) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SpaceID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `reservation` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"space_id"}),
				strmangle.WhereClause("`", "`", 0, reservationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SpaceID = o.ID
		}
	}

	if o.R == nil {
		o.R = &spaceR{
			Reservations: related,
		}
	} else {
		o.R.Reservations = append(o.R.Reservations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &reservationR{
				Space: o,
			}
		} else {
			rel.R.Space = o
		}
	}
	return nil
}

// Spaces retrieves all the records using an executor.
func Spaces(mods ...qm.QueryMod) spaceQuery {
	mods = append(mods, qm.From("`spaces`"))
	return spaceQuery{NewQuery(mods...)}
}

// FindSpace retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpace(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Space, error) {
	spaceObj := &Space{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `spaces` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, spaceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from spaces")
	}

	return spaceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Space) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no spaces provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spaceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spaceInsertCacheMut.RLock()
	cache, cached := spaceInsertCache[key]
	spaceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spaceAllColumns,
			spaceColumnsWithDefault,
			spaceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spaceType, spaceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spaceType, spaceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `spaces` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `spaces` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `spaces` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, spacePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into spaces")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == spaceMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for spaces")
	}

CacheNoHooks:
	if !cached {
		spaceInsertCacheMut.Lock()
		spaceInsertCache[key] = cache
		spaceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Space.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Space) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	spaceUpdateCacheMut.RLock()
	cache, cached := spaceUpdateCache[key]
	spaceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spaceAllColumns,
			spacePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update spaces, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `spaces` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, spacePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spaceType, spaceMapping, append(wl, spacePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update spaces row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for spaces")
	}

	if !cached {
		spaceUpdateCacheMut.Lock()
		spaceUpdateCache[key] = cache
		spaceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q spaceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for spaces")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SpaceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `spaces` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spacePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in space slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all space")
	}
	return rowsAff, nil
}

var mySQLSpaceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Space) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no spaces provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spaceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSpaceUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	spaceUpsertCacheMut.RLock()
	cache, cached := spaceUpsertCache[key]
	spaceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spaceAllColumns,
			spaceColumnsWithDefault,
			spaceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			spaceAllColumns,
			spacePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("entity: unable to upsert spaces, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "spaces", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `spaces` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(spaceType, spaceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spaceType, spaceMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert for spaces")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == spaceMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(spaceType, spaceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for spaces")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for spaces")
	}

CacheNoHooks:
	if !cached {
		spaceUpsertCacheMut.Lock()
		spaceUpsertCache[key] = cache
		spaceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Space record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Space) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Space provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spacePrimaryKeyMapping)
	sql := "DELETE FROM `spaces` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for spaces")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q spaceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no spaceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for spaces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SpaceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(spaceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `spaces` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spacePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from space slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for spaces")
	}

	if len(spaceAfterDeleteHooks) != 0 {
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
func (o *Space) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpace(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpaceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpaceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `spaces`.* FROM `spaces` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spacePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in SpaceSlice")
	}

	*o = slice

	return nil
}

// SpaceExists checks if the Space row exists.
func SpaceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `spaces` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if spaces exists")
	}

	return exists, nil
}