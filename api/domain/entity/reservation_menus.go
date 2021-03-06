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

// ReservationMenu is an object representing the database table.
type ReservationMenu struct {
	ReservationID    int64     `boil:"reservation_id" json:"reservation_id" toml:"reservation_id" yaml:"reservation_id"`
	BeauticianMenuID int64     `boil:"beautician_menu_id" json:"beautician_menu_id" toml:"beautician_menu_id" yaml:"beautician_menu_id"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *reservationMenuR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reservationMenuL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReservationMenuColumns = struct {
	ReservationID    string
	BeauticianMenuID string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ReservationID:    "reservation_id",
	BeauticianMenuID: "beautician_menu_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// Generated where

var ReservationMenuWhere = struct {
	ReservationID    whereHelperint64
	BeauticianMenuID whereHelperint64
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
	DeletedAt        whereHelpernull_Time
}{
	ReservationID:    whereHelperint64{field: "`reservation_menus`.`reservation_id`"},
	BeauticianMenuID: whereHelperint64{field: "`reservation_menus`.`beautician_menu_id`"},
	CreatedAt:        whereHelpertime_Time{field: "`reservation_menus`.`created_at`"},
	UpdatedAt:        whereHelpertime_Time{field: "`reservation_menus`.`updated_at`"},
	DeletedAt:        whereHelpernull_Time{field: "`reservation_menus`.`deleted_at`"},
}

// ReservationMenuRels is where relationship names are stored.
var ReservationMenuRels = struct {
	BeauticianMenu string
	Reservation    string
}{
	BeauticianMenu: "BeauticianMenu",
	Reservation:    "Reservation",
}

// reservationMenuR is where relationships are stored.
type reservationMenuR struct {
	BeauticianMenu *BeauticianMenu
	Reservation    *Reservation
}

// NewStruct creates a new relationship struct
func (*reservationMenuR) NewStruct() *reservationMenuR {
	return &reservationMenuR{}
}

// reservationMenuL is where Load methods for each relationship are stored.
type reservationMenuL struct{}

var (
	reservationMenuAllColumns            = []string{"reservation_id", "beautician_menu_id", "created_at", "updated_at", "deleted_at"}
	reservationMenuColumnsWithoutDefault = []string{"reservation_id", "beautician_menu_id", "created_at", "updated_at", "deleted_at"}
	reservationMenuColumnsWithDefault    = []string{}
	reservationMenuPrimaryKeyColumns     = []string{"reservation_id", "beautician_menu_id"}
)

type (
	// ReservationMenuSlice is an alias for a slice of pointers to ReservationMenu.
	// This should generally be used opposed to []ReservationMenu.
	ReservationMenuSlice []*ReservationMenu
	// ReservationMenuHook is the signature for custom ReservationMenu hook methods
	ReservationMenuHook func(context.Context, boil.ContextExecutor, *ReservationMenu) error

	reservationMenuQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reservationMenuType                 = reflect.TypeOf(&ReservationMenu{})
	reservationMenuMapping              = queries.MakeStructMapping(reservationMenuType)
	reservationMenuPrimaryKeyMapping, _ = queries.BindMapping(reservationMenuType, reservationMenuMapping, reservationMenuPrimaryKeyColumns)
	reservationMenuInsertCacheMut       sync.RWMutex
	reservationMenuInsertCache          = make(map[string]insertCache)
	reservationMenuUpdateCacheMut       sync.RWMutex
	reservationMenuUpdateCache          = make(map[string]updateCache)
	reservationMenuUpsertCacheMut       sync.RWMutex
	reservationMenuUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var reservationMenuBeforeInsertHooks []ReservationMenuHook
var reservationMenuBeforeUpdateHooks []ReservationMenuHook
var reservationMenuBeforeDeleteHooks []ReservationMenuHook
var reservationMenuBeforeUpsertHooks []ReservationMenuHook

var reservationMenuAfterInsertHooks []ReservationMenuHook
var reservationMenuAfterSelectHooks []ReservationMenuHook
var reservationMenuAfterUpdateHooks []ReservationMenuHook
var reservationMenuAfterDeleteHooks []ReservationMenuHook
var reservationMenuAfterUpsertHooks []ReservationMenuHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ReservationMenu) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ReservationMenu) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ReservationMenu) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ReservationMenu) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ReservationMenu) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ReservationMenu) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ReservationMenu) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ReservationMenu) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ReservationMenu) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reservationMenuAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReservationMenuHook registers your hook function for all future operations.
func AddReservationMenuHook(hookPoint boil.HookPoint, reservationMenuHook ReservationMenuHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		reservationMenuBeforeInsertHooks = append(reservationMenuBeforeInsertHooks, reservationMenuHook)
	case boil.BeforeUpdateHook:
		reservationMenuBeforeUpdateHooks = append(reservationMenuBeforeUpdateHooks, reservationMenuHook)
	case boil.BeforeDeleteHook:
		reservationMenuBeforeDeleteHooks = append(reservationMenuBeforeDeleteHooks, reservationMenuHook)
	case boil.BeforeUpsertHook:
		reservationMenuBeforeUpsertHooks = append(reservationMenuBeforeUpsertHooks, reservationMenuHook)
	case boil.AfterInsertHook:
		reservationMenuAfterInsertHooks = append(reservationMenuAfterInsertHooks, reservationMenuHook)
	case boil.AfterSelectHook:
		reservationMenuAfterSelectHooks = append(reservationMenuAfterSelectHooks, reservationMenuHook)
	case boil.AfterUpdateHook:
		reservationMenuAfterUpdateHooks = append(reservationMenuAfterUpdateHooks, reservationMenuHook)
	case boil.AfterDeleteHook:
		reservationMenuAfterDeleteHooks = append(reservationMenuAfterDeleteHooks, reservationMenuHook)
	case boil.AfterUpsertHook:
		reservationMenuAfterUpsertHooks = append(reservationMenuAfterUpsertHooks, reservationMenuHook)
	}
}

// One returns a single reservationMenu record from the query.
func (q reservationMenuQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ReservationMenu, error) {
	o := &ReservationMenu{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for reservation_menus")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ReservationMenu records from the query.
func (q reservationMenuQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReservationMenuSlice, error) {
	var o []*ReservationMenu

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to ReservationMenu slice")
	}

	if len(reservationMenuAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ReservationMenu records in the query.
func (q reservationMenuQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count reservation_menus rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q reservationMenuQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if reservation_menus exists")
	}

	return count > 0, nil
}

// BeauticianMenu pointed to by the foreign key.
func (o *ReservationMenu) BeauticianMenu(mods ...qm.QueryMod) beauticianMenuQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BeauticianMenuID),
	}

	queryMods = append(queryMods, mods...)

	query := BeauticianMenus(queryMods...)
	queries.SetFrom(query.Query, "`beautician_menus`")

	return query
}

// Reservation pointed to by the foreign key.
func (o *ReservationMenu) Reservation(mods ...qm.QueryMod) reservationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ReservationID),
	}

	queryMods = append(queryMods, mods...)

	query := Reservations(queryMods...)
	queries.SetFrom(query.Query, "`reservations`")

	return query
}

// LoadBeauticianMenu allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reservationMenuL) LoadBeauticianMenu(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReservationMenu interface{}, mods queries.Applicator) error {
	var slice []*ReservationMenu
	var object *ReservationMenu

	if singular {
		object = maybeReservationMenu.(*ReservationMenu)
	} else {
		slice = *maybeReservationMenu.(*[]*ReservationMenu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reservationMenuR{}
		}
		args = append(args, object.BeauticianMenuID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reservationMenuR{}
			}

			for _, a := range args {
				if a == obj.BeauticianMenuID {
					continue Outer
				}
			}

			args = append(args, obj.BeauticianMenuID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`beautician_menus`), qm.WhereIn(`beautician_menus.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BeauticianMenu")
	}

	var resultSlice []*BeauticianMenu
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BeauticianMenu")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for beautician_menus")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for beautician_menus")
	}

	if len(reservationMenuAfterSelectHooks) != 0 {
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
		object.R.BeauticianMenu = foreign
		if foreign.R == nil {
			foreign.R = &beauticianMenuR{}
		}
		foreign.R.ReservationMenus = append(foreign.R.ReservationMenus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BeauticianMenuID == foreign.ID {
				local.R.BeauticianMenu = foreign
				if foreign.R == nil {
					foreign.R = &beauticianMenuR{}
				}
				foreign.R.ReservationMenus = append(foreign.R.ReservationMenus, local)
				break
			}
		}
	}

	return nil
}

// LoadReservation allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reservationMenuL) LoadReservation(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReservationMenu interface{}, mods queries.Applicator) error {
	var slice []*ReservationMenu
	var object *ReservationMenu

	if singular {
		object = maybeReservationMenu.(*ReservationMenu)
	} else {
		slice = *maybeReservationMenu.(*[]*ReservationMenu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reservationMenuR{}
		}
		args = append(args, object.ReservationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reservationMenuR{}
			}

			for _, a := range args {
				if a == obj.ReservationID {
					continue Outer
				}
			}

			args = append(args, obj.ReservationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`reservations`), qm.WhereIn(`reservations.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Reservation")
	}

	var resultSlice []*Reservation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Reservation")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for reservations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reservations")
	}

	if len(reservationMenuAfterSelectHooks) != 0 {
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
		object.R.Reservation = foreign
		if foreign.R == nil {
			foreign.R = &reservationR{}
		}
		foreign.R.ReservationMenus = append(foreign.R.ReservationMenus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ReservationID == foreign.ID {
				local.R.Reservation = foreign
				if foreign.R == nil {
					foreign.R = &reservationR{}
				}
				foreign.R.ReservationMenus = append(foreign.R.ReservationMenus, local)
				break
			}
		}
	}

	return nil
}

// SetBeauticianMenu of the reservationMenu to the related item.
// Sets o.R.BeauticianMenu to related.
// Adds o to related.R.ReservationMenus.
func (o *ReservationMenu) SetBeauticianMenu(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BeauticianMenu) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `reservation_menus` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"beautician_menu_id"}),
		strmangle.WhereClause("`", "`", 0, reservationMenuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ReservationID, o.BeauticianMenuID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BeauticianMenuID = related.ID
	if o.R == nil {
		o.R = &reservationMenuR{
			BeauticianMenu: related,
		}
	} else {
		o.R.BeauticianMenu = related
	}

	if related.R == nil {
		related.R = &beauticianMenuR{
			ReservationMenus: ReservationMenuSlice{o},
		}
	} else {
		related.R.ReservationMenus = append(related.R.ReservationMenus, o)
	}

	return nil
}

// SetReservation of the reservationMenu to the related item.
// Sets o.R.Reservation to related.
// Adds o to related.R.ReservationMenus.
func (o *ReservationMenu) SetReservation(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Reservation) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `reservation_menus` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"reservation_id"}),
		strmangle.WhereClause("`", "`", 0, reservationMenuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ReservationID, o.BeauticianMenuID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ReservationID = related.ID
	if o.R == nil {
		o.R = &reservationMenuR{
			Reservation: related,
		}
	} else {
		o.R.Reservation = related
	}

	if related.R == nil {
		related.R = &reservationR{
			ReservationMenus: ReservationMenuSlice{o},
		}
	} else {
		related.R.ReservationMenus = append(related.R.ReservationMenus, o)
	}

	return nil
}

// ReservationMenus retrieves all the records using an executor.
func ReservationMenus(mods ...qm.QueryMod) reservationMenuQuery {
	mods = append(mods, qm.From("`reservation_menus`"))
	return reservationMenuQuery{NewQuery(mods...)}
}

// FindReservationMenu retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReservationMenu(ctx context.Context, exec boil.ContextExecutor, reservationID int64, beauticianMenuID int64, selectCols ...string) (*ReservationMenu, error) {
	reservationMenuObj := &ReservationMenu{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `reservation_menus` where `reservation_id`=? AND `beautician_menu_id`=?", sel,
	)

	q := queries.Raw(query, reservationID, beauticianMenuID)

	err := q.Bind(ctx, exec, reservationMenuObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from reservation_menus")
	}

	return reservationMenuObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReservationMenu) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no reservation_menus provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(reservationMenuColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reservationMenuInsertCacheMut.RLock()
	cache, cached := reservationMenuInsertCache[key]
	reservationMenuInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reservationMenuAllColumns,
			reservationMenuColumnsWithDefault,
			reservationMenuColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `reservation_menus` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `reservation_menus` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `reservation_menus` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, reservationMenuPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into reservation_menus")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ReservationID,
		o.BeauticianMenuID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for reservation_menus")
	}

CacheNoHooks:
	if !cached {
		reservationMenuInsertCacheMut.Lock()
		reservationMenuInsertCache[key] = cache
		reservationMenuInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ReservationMenu.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReservationMenu) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	reservationMenuUpdateCacheMut.RLock()
	cache, cached := reservationMenuUpdateCache[key]
	reservationMenuUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reservationMenuAllColumns,
			reservationMenuPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update reservation_menus, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `reservation_menus` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, reservationMenuPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, append(wl, reservationMenuPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update reservation_menus row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for reservation_menus")
	}

	if !cached {
		reservationMenuUpdateCacheMut.Lock()
		reservationMenuUpdateCache[key] = cache
		reservationMenuUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q reservationMenuQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for reservation_menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for reservation_menus")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReservationMenuSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reservationMenuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `reservation_menus` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reservationMenuPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in reservationMenu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all reservationMenu")
	}
	return rowsAff, nil
}

var mySQLReservationMenuUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ReservationMenu) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no reservation_menus provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(reservationMenuColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLReservationMenuUniqueColumns, o)

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

	reservationMenuUpsertCacheMut.RLock()
	cache, cached := reservationMenuUpsertCache[key]
	reservationMenuUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			reservationMenuAllColumns,
			reservationMenuColumnsWithDefault,
			reservationMenuColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			reservationMenuAllColumns,
			reservationMenuPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("entity: unable to upsert reservation_menus, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "reservation_menus", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `reservation_menus` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert for reservation_menus")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(reservationMenuType, reservationMenuMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for reservation_menus")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for reservation_menus")
	}

CacheNoHooks:
	if !cached {
		reservationMenuUpsertCacheMut.Lock()
		reservationMenuUpsertCache[key] = cache
		reservationMenuUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ReservationMenu record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReservationMenu) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no ReservationMenu provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reservationMenuPrimaryKeyMapping)
	sql := "DELETE FROM `reservation_menus` WHERE `reservation_id`=? AND `beautician_menu_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from reservation_menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for reservation_menus")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q reservationMenuQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no reservationMenuQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from reservation_menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for reservation_menus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReservationMenuSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(reservationMenuBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reservationMenuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `reservation_menus` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reservationMenuPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from reservationMenu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for reservation_menus")
	}

	if len(reservationMenuAfterDeleteHooks) != 0 {
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
func (o *ReservationMenu) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReservationMenu(ctx, exec, o.ReservationID, o.BeauticianMenuID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReservationMenuSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReservationMenuSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reservationMenuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `reservation_menus`.* FROM `reservation_menus` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reservationMenuPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in ReservationMenuSlice")
	}

	*o = slice

	return nil
}

// ReservationMenuExists checks if the ReservationMenu row exists.
func ReservationMenuExists(ctx context.Context, exec boil.ContextExecutor, reservationID int64, beauticianMenuID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `reservation_menus` where `reservation_id`=? AND `beautician_menu_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, reservationID, beauticianMenuID)
	}
	row := exec.QueryRowContext(ctx, sql, reservationID, beauticianMenuID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if reservation_menus exists")
	}

	return exists, nil
}
