// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Beauticians", testBeauticians)
	t.Run("Guests", testGuests)
	t.Run("Menus", testMenus)
	t.Run("Stores", testStores)
}

func TestDelete(t *testing.T) {
	t.Run("Beauticians", testBeauticiansDelete)
	t.Run("Guests", testGuestsDelete)
	t.Run("Menus", testMenusDelete)
	t.Run("Stores", testStoresDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Beauticians", testBeauticiansQueryDeleteAll)
	t.Run("Guests", testGuestsQueryDeleteAll)
	t.Run("Menus", testMenusQueryDeleteAll)
	t.Run("Stores", testStoresQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Beauticians", testBeauticiansSliceDeleteAll)
	t.Run("Guests", testGuestsSliceDeleteAll)
	t.Run("Menus", testMenusSliceDeleteAll)
	t.Run("Stores", testStoresSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Beauticians", testBeauticiansExists)
	t.Run("Guests", testGuestsExists)
	t.Run("Menus", testMenusExists)
	t.Run("Stores", testStoresExists)
}

func TestFind(t *testing.T) {
	t.Run("Beauticians", testBeauticiansFind)
	t.Run("Guests", testGuestsFind)
	t.Run("Menus", testMenusFind)
	t.Run("Stores", testStoresFind)
}

func TestBind(t *testing.T) {
	t.Run("Beauticians", testBeauticiansBind)
	t.Run("Guests", testGuestsBind)
	t.Run("Menus", testMenusBind)
	t.Run("Stores", testStoresBind)
}

func TestOne(t *testing.T) {
	t.Run("Beauticians", testBeauticiansOne)
	t.Run("Guests", testGuestsOne)
	t.Run("Menus", testMenusOne)
	t.Run("Stores", testStoresOne)
}

func TestAll(t *testing.T) {
	t.Run("Beauticians", testBeauticiansAll)
	t.Run("Guests", testGuestsAll)
	t.Run("Menus", testMenusAll)
	t.Run("Stores", testStoresAll)
}

func TestCount(t *testing.T) {
	t.Run("Beauticians", testBeauticiansCount)
	t.Run("Guests", testGuestsCount)
	t.Run("Menus", testMenusCount)
	t.Run("Stores", testStoresCount)
}

func TestHooks(t *testing.T) {
	t.Run("Beauticians", testBeauticiansHooks)
	t.Run("Guests", testGuestsHooks)
	t.Run("Menus", testMenusHooks)
	t.Run("Stores", testStoresHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Beauticians", testBeauticiansInsert)
	t.Run("Beauticians", testBeauticiansInsertWhitelist)
	t.Run("Guests", testGuestsInsert)
	t.Run("Guests", testGuestsInsertWhitelist)
	t.Run("Menus", testMenusInsert)
	t.Run("Menus", testMenusInsertWhitelist)
	t.Run("Stores", testStoresInsert)
	t.Run("Stores", testStoresInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("MenuToBeauticianUsingBeautician", testMenuToOneBeauticianUsingBeautician)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BeauticianToMenus", testBeauticianToManyMenus)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("MenuToBeauticianUsingMenus", testMenuToOneSetOpBeauticianUsingBeautician)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BeauticianToMenus", testBeauticianToManyAddOpMenus)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Beauticians", testBeauticiansReload)
	t.Run("Guests", testGuestsReload)
	t.Run("Menus", testMenusReload)
	t.Run("Stores", testStoresReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Beauticians", testBeauticiansReloadAll)
	t.Run("Guests", testGuestsReloadAll)
	t.Run("Menus", testMenusReloadAll)
	t.Run("Stores", testStoresReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Beauticians", testBeauticiansSelect)
	t.Run("Guests", testGuestsSelect)
	t.Run("Menus", testMenusSelect)
	t.Run("Stores", testStoresSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Beauticians", testBeauticiansUpdate)
	t.Run("Guests", testGuestsUpdate)
	t.Run("Menus", testMenusUpdate)
	t.Run("Stores", testStoresUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Beauticians", testBeauticiansSliceUpdateAll)
	t.Run("Guests", testGuestsSliceUpdateAll)
	t.Run("Menus", testMenusSliceUpdateAll)
	t.Run("Stores", testStoresSliceUpdateAll)
}
