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

func testBeauticianMenus(t *testing.T) {
	t.Parallel()

	query := BeauticianMenus()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBeauticianMenusDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
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

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBeauticianMenusQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BeauticianMenus().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBeauticianMenusSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BeauticianMenuSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBeauticianMenusExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BeauticianMenuExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BeauticianMenu exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BeauticianMenuExists to return true, but got false.")
	}
}

func testBeauticianMenusFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	beauticianMenuFound, err := FindBeauticianMenu(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if beauticianMenuFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBeauticianMenusBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BeauticianMenus().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBeauticianMenusOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BeauticianMenus().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBeauticianMenusAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	beauticianMenuOne := &BeauticianMenu{}
	beauticianMenuTwo := &BeauticianMenu{}
	if err = randomize.Struct(seed, beauticianMenuOne, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}
	if err = randomize.Struct(seed, beauticianMenuTwo, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = beauticianMenuOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = beauticianMenuTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BeauticianMenus().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBeauticianMenusCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	beauticianMenuOne := &BeauticianMenu{}
	beauticianMenuTwo := &BeauticianMenu{}
	if err = randomize.Struct(seed, beauticianMenuOne, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}
	if err = randomize.Struct(seed, beauticianMenuTwo, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = beauticianMenuOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = beauticianMenuTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func beauticianMenuBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func beauticianMenuAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BeauticianMenu) error {
	*o = BeauticianMenu{}
	return nil
}

func testBeauticianMenusHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BeauticianMenu{}
	o := &BeauticianMenu{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu object: %s", err)
	}

	AddBeauticianMenuHook(boil.BeforeInsertHook, beauticianMenuBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	beauticianMenuBeforeInsertHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.AfterInsertHook, beauticianMenuAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	beauticianMenuAfterInsertHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.AfterSelectHook, beauticianMenuAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	beauticianMenuAfterSelectHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.BeforeUpdateHook, beauticianMenuBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	beauticianMenuBeforeUpdateHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.AfterUpdateHook, beauticianMenuAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	beauticianMenuAfterUpdateHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.BeforeDeleteHook, beauticianMenuBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	beauticianMenuBeforeDeleteHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.AfterDeleteHook, beauticianMenuAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	beauticianMenuAfterDeleteHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.BeforeUpsertHook, beauticianMenuBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	beauticianMenuBeforeUpsertHooks = []BeauticianMenuHook{}

	AddBeauticianMenuHook(boil.AfterUpsertHook, beauticianMenuAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	beauticianMenuAfterUpsertHooks = []BeauticianMenuHook{}
}

func testBeauticianMenusInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBeauticianMenusInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(beauticianMenuColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBeauticianMenuToManyReservationMenus(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BeauticianMenu
	var b, c ReservationMenu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, reservationMenuDBTypes, false, reservationMenuColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BeauticianMenuID = a.ID
	c.BeauticianMenuID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ReservationMenus().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BeauticianMenuID == b.BeauticianMenuID {
			bFound = true
		}
		if v.BeauticianMenuID == c.BeauticianMenuID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BeauticianMenuSlice{&a}
	if err = a.L.LoadReservationMenus(ctx, tx, false, (*[]*BeauticianMenu)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ReservationMenus); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ReservationMenus = nil
	if err = a.L.LoadReservationMenus(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ReservationMenus); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBeauticianMenuToManyAddOpReservationMenus(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BeauticianMenu
	var b, c, d, e ReservationMenu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, beauticianMenuDBTypes, false, strmangle.SetComplement(beauticianMenuPrimaryKeyColumns, beauticianMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ReservationMenu{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, reservationMenuDBTypes, false, strmangle.SetComplement(reservationMenuPrimaryKeyColumns, reservationMenuColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ReservationMenu{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddReservationMenus(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BeauticianMenuID {
			t.Error("foreign key was wrong value", a.ID, first.BeauticianMenuID)
		}
		if a.ID != second.BeauticianMenuID {
			t.Error("foreign key was wrong value", a.ID, second.BeauticianMenuID)
		}

		if first.R.BeauticianMenu != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BeauticianMenu != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ReservationMenus[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ReservationMenus[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ReservationMenus().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBeauticianMenuToOneBeauticianUsingBeautician(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BeauticianMenu
	var foreign Beautician

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, beauticianDBTypes, false, beauticianColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Beautician struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BeauticianID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Beautician().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BeauticianMenuSlice{&local}
	if err = local.L.LoadBeautician(ctx, tx, false, (*[]*BeauticianMenu)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Beautician == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Beautician = nil
	if err = local.L.LoadBeautician(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Beautician == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBeauticianMenuToOneMenuUsingMenu(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BeauticianMenu
	var foreign Menu

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, beauticianMenuDBTypes, false, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, menuDBTypes, false, menuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Menu struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MenuID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Menu().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BeauticianMenuSlice{&local}
	if err = local.L.LoadMenu(ctx, tx, false, (*[]*BeauticianMenu)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Menu = nil
	if err = local.L.LoadMenu(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBeauticianMenuToOneSetOpBeauticianUsingBeautician(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BeauticianMenu
	var b, c Beautician

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, beauticianMenuDBTypes, false, strmangle.SetComplement(beauticianMenuPrimaryKeyColumns, beauticianMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, beauticianDBTypes, false, strmangle.SetComplement(beauticianPrimaryKeyColumns, beauticianColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, beauticianDBTypes, false, strmangle.SetComplement(beauticianPrimaryKeyColumns, beauticianColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Beautician{&b, &c} {
		err = a.SetBeautician(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Beautician != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BeauticianMenus[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BeauticianID != x.ID {
			t.Error("foreign key was wrong value", a.BeauticianID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BeauticianID))
		reflect.Indirect(reflect.ValueOf(&a.BeauticianID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BeauticianID != x.ID {
			t.Error("foreign key was wrong value", a.BeauticianID, x.ID)
		}
	}
}
func testBeauticianMenuToOneSetOpMenuUsingMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BeauticianMenu
	var b, c Menu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, beauticianMenuDBTypes, false, strmangle.SetComplement(beauticianMenuPrimaryKeyColumns, beauticianMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Menu{&b, &c} {
		err = a.SetMenu(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Menu != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BeauticianMenus[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MenuID != x.ID {
			t.Error("foreign key was wrong value", a.MenuID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MenuID))
		reflect.Indirect(reflect.ValueOf(&a.MenuID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MenuID != x.ID {
			t.Error("foreign key was wrong value", a.MenuID, x.ID)
		}
	}
}

func testBeauticianMenusReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
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

func testBeauticianMenusReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BeauticianMenuSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBeauticianMenusSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BeauticianMenus().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	beauticianMenuDBTypes = map[string]string{`ID`: `bigint`, `Name`: `varchar`, `Price`: `bigint`, `BeauticianID`: `bigint`, `MenuID`: `bigint`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `DeletedAt`: `datetime`}
	_                     = bytes.MinRead
)

func testBeauticianMenusUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(beauticianMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(beauticianMenuAllColumns) == len(beauticianMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBeauticianMenusSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(beauticianMenuAllColumns) == len(beauticianMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BeauticianMenu{}
	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, beauticianMenuDBTypes, true, beauticianMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(beauticianMenuAllColumns, beauticianMenuPrimaryKeyColumns) {
		fields = beauticianMenuAllColumns
	} else {
		fields = strmangle.SetComplement(
			beauticianMenuAllColumns,
			beauticianMenuPrimaryKeyColumns,
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

	slice := BeauticianMenuSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBeauticianMenusUpsert(t *testing.T) {
	t.Parallel()

	if len(beauticianMenuAllColumns) == len(beauticianMenuPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBeauticianMenuUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BeauticianMenu{}
	if err = randomize.Struct(seed, &o, beauticianMenuDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BeauticianMenu: %s", err)
	}

	count, err := BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, beauticianMenuDBTypes, false, beauticianMenuPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BeauticianMenu struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BeauticianMenu: %s", err)
	}

	count, err = BeauticianMenus().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
