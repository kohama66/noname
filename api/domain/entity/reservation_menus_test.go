// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testReservationMenus(t *testing.T) {
	t.Parallel()

	query := ReservationMenus()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testReservationMenusDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReservationMenusQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ReservationMenus().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReservationMenusSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReservationMenuSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReservationMenusExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ReservationMenuExists(ctx, tx, o.ReservationID, o.BeauticianMenuID)
	if err != nil {
		t.Errorf("Unable to check if ReservationMenu exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReservationMenuExists to return true, but got false.")
	}
}

func testReservationMenusFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	reservationMenuFound, err := FindReservationMenu(ctx, tx, o.ReservationID, o.BeauticianMenuID)
	if err != nil {
		t.Error(err)
	}

	if reservationMenuFound == nil {
		t.Error("want a record, got nil")
	}
}

func testReservationMenusBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ReservationMenus().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testReservationMenusOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ReservationMenus().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReservationMenusAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reservationMenuOne := &ReservationMenu{}
	reservationMenuTwo := &ReservationMenu{}
	if err = randomize.Struct(seed, reservationMenuOne, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}
	if err = randomize.Struct(seed, reservationMenuTwo, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reservationMenuOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reservationMenuTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReservationMenus().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReservationMenusCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	reservationMenuOne := &ReservationMenu{}
	reservationMenuTwo := &ReservationMenu{}
	if err = randomize.Struct(seed, reservationMenuOne, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}
	if err = randomize.Struct(seed, reservationMenuTwo, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reservationMenuOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reservationMenuTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func reservationMenuBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func reservationMenuAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ReservationMenu) error {
	*o = ReservationMenu{}
	return nil
}

func testReservationMenusHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ReservationMenu{}
	o := &ReservationMenu{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ReservationMenu object: %s", err)
	}

	AddReservationMenuHook(boil.BeforeInsertHook, reservationMenuBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	reservationMenuBeforeInsertHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.AfterInsertHook, reservationMenuAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	reservationMenuAfterInsertHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.AfterSelectHook, reservationMenuAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	reservationMenuAfterSelectHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.BeforeUpdateHook, reservationMenuBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	reservationMenuBeforeUpdateHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.AfterUpdateHook, reservationMenuAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	reservationMenuAfterUpdateHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.BeforeDeleteHook, reservationMenuBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	reservationMenuBeforeDeleteHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.AfterDeleteHook, reservationMenuAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	reservationMenuAfterDeleteHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.BeforeUpsertHook, reservationMenuBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	reservationMenuBeforeUpsertHooks = []ReservationMenuHook{}

	AddReservationMenuHook(boil.AfterUpsertHook, reservationMenuAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	reservationMenuAfterUpsertHooks = []ReservationMenuHook{}
}

func testReservationMenusInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReservationMenusInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(reservationMenuColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReservationMenuToOneBeauticianMenuUsingBeauticianMenu(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ReservationMenu
	var foreign BeauticianMenu

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BeauticianMenuID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BeauticianMenu().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ReservationMenuSlice{&local}
	if err = local.L.LoadBeauticianMenu(ctx, tx, false, (*[]*ReservationMenu)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BeauticianMenu == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BeauticianMenu = nil
	if err = local.L.LoadBeauticianMenu(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BeauticianMenu == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testReservationMenuToOneReservationUsingReservation(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ReservationMenu
	var foreign Reservation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, reservationDBTypes, false, reservationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reservation struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ReservationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Reservation().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ReservationMenuSlice{&local}
	if err = local.L.LoadReservation(ctx, tx, false, (*[]*ReservationMenu)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Reservation == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Reservation = nil
	if err = local.L.LoadReservation(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Reservation == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testReservationMenuToOneSetOpBeauticianMenuUsingBeauticianMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ReservationMenu
	var b, c BeauticianMenu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, reservationMenuDBTypes, false, strmangle.SetComplement(reservationMenuPrimaryKeyColumns, reservationMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, beauticianMenuDBTypes, false, strmangle.SetComplement(beauticianMenuPrimaryKeyColumns, beauticianMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, beauticianMenuDBTypes, false, strmangle.SetComplement(beauticianMenuPrimaryKeyColumns, beauticianMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BeauticianMenu{&b, &c} {
		err = a.SetBeauticianMenu(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BeauticianMenu != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ReservationMenus[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BeauticianMenuID != x.ID {
			t.Error("foreign key was wrong value", a.BeauticianMenuID)
		}

		if exists, err := ReservationMenuExists(ctx, tx, a.ReservationID, a.BeauticianMenuID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testReservationMenuToOneSetOpReservationUsingReservation(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ReservationMenu
	var b, c Reservation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, reservationMenuDBTypes, false, strmangle.SetComplement(reservationMenuPrimaryKeyColumns, reservationMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, reservationDBTypes, false, strmangle.SetComplement(reservationPrimaryKeyColumns, reservationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, reservationDBTypes, false, strmangle.SetComplement(reservationPrimaryKeyColumns, reservationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Reservation{&b, &c} {
		err = a.SetReservation(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Reservation != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ReservationMenus[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ReservationID != x.ID {
			t.Error("foreign key was wrong value", a.ReservationID)
		}

		if exists, err := ReservationMenuExists(ctx, tx, a.ReservationID, a.BeauticianMenuID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testReservationMenusReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReservationMenusReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReservationMenuSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReservationMenusSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReservationMenus().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	reservationMenuDBTypes = map[string]string{`ReservationID`: `bigint`, `BeauticianMenuID`: `bigint`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `DeletedAt`: `datetime`}
	_                      = bytes.MinRead
)

func testReservationMenusUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(reservationMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(reservationMenuAllColumns) == len(reservationMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testReservationMenusSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(reservationMenuAllColumns) == len(reservationMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReservationMenu{}
	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reservationMenuDBTypes, true, reservationMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(reservationMenuAllColumns, reservationMenuPrimaryKeyColumns) {
		fields = reservationMenuAllColumns
	} else {
		fields = strmangle.SetComplement(
			reservationMenuAllColumns,
			reservationMenuPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ReservationMenuSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testReservationMenusUpsert(t *testing.T) {
	t.Parallel()

	if len(reservationMenuAllColumns) == len(reservationMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLReservationMenuUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ReservationMenu{}
	if err = randomize.Struct(seed, &o, reservationMenuDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReservationMenu: %s", err)
	}

	count, err := ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, reservationMenuDBTypes, false, reservationMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReservationMenu struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReservationMenu: %s", err)
	}

	count, err = ReservationMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}