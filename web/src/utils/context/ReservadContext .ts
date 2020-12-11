import { createContext, useState } from "react"
import { Beautician, initBeautician, isBeauticianInterface } from "../../package/interface/Beautician"
import { Menu } from "../../package/interface/Menu"
import { initSalon, isSalonInterface, Salon } from "../../package/interface/Salon"

export const ReservedContext = createContext({
  setSelectValue: (value: Beautician | Salon | Menu[] | string) => { },
  isAllChecked: (): boolean => false,
  beautician: initBeautician,
  store: initSalon,
  getMenuIDs: (): string[] => [""],
  reservationDate: <string | undefined>undefined,
  resetAllReservedState: () => {},
})

export const useReservedContext = () => {
  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [menus, setMenus] = useState<Menu[]>([])
  const [reservationDate, setReservationDate] = useState<string | undefined>(undefined)

  const setSelectValue = (value: Beautician | Salon | Menu[] | string) => {
    if (isBeauticianInterface(value)) {
      setBeautician(value)
    } else if (isSalonInterface(value)) {
      setStore(value)
    } else if (typeof value === "string") {
      setReservationDate(value)
    } else if (Object.prototype.toString.call(value) === "[object Array]" && Array.isArray(value) && value.length !== 0) {
      setMenus(value)
    }
  }
  const getSelectID = (type: "beautician" | "store" | "date" | "menu"): string | string[] | undefined => {
    switch (type) {
      case "beautician":
        if (beautician !== initBeautician) {
          return beautician.randId;
        } else {
          return undefined
        }
      case "store":
        if (store !== initSalon) {
          return store.randId;
        } else {
          return undefined
        }
      case "date":
        if (reservationDate !== "") {
          return reservationDate;
        } else {
          return undefined
        }
      case "menu":
        const menuIDs: string[] = []
        if (Object.prototype.toString.call(menus) === "[object Array]" && Array.isArray(menus)) {
          menus.map((menu) => {
            menuIDs.push(menu.randId)
          })
        }
        return menuIDs
    }
  }
  const getMenuIDs = (): string[] => {
    const menuIDs: string[] = []
    if (Object.prototype.toString.call(menus) === "[object Array]" && Array.isArray(menus)) {
      menus.map((menu) => {
        menuIDs.push(menu.randId)
      })
    }
    return menuIDs
  }
  const isAllChecked = (): boolean => {
    if (beautician !== initBeautician && store !== initSalon && menus.length !== 0 && reservationDate !== undefined) {
      return true
    } else {
      return false
    }
  }
  const resetAllReservedState = () => {
    setBeautician(initBeautician)
    setStore(initSalon)
    setMenus([])
    setReservationDate(undefined)
  }
  return { setSelectValue, isAllChecked, beautician, store, getMenuIDs, reservationDate, resetAllReservedState }
}
