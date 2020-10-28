import React, { FC, useState, createContext, useEffect } from 'react';
import GuestComponent from '../../../components/views/guest/Guest'
import ChooseBeautician from "./beautician"
import ChooseStore from './store'
import ChooseDate from './date'
import ChooseMenu from './menu'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
  useHistory,
} from "react-router-dom";
import FinalComfirmation from '../finalComfirmation';
import { Beautician, initBeautician, isBeauticianInterface } from '../../../package/interface/Beautician';
import { initSalon, isSalonInterface, Salon } from '../../../package/interface/Salon';
import { Menu } from '../../../package/interface/Menu';

// コンテキスト
export const SetSelectValueContext = createContext((value: Beautician | Salon | Menu[] | string) => { });
export const GeterSelectIDContext = createContext((type: "beautician" | "store" | "date" | "menu"): string | string[] | undefined => "")
export const GeterSelectValueContext = createContext((): [Beautician, Salon, Menu[], string] => [initBeautician, initSalon, [], ""])

const Guest: FC = () => {
  const match = useRouteMatch()
  const history = useHistory()

  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [menus, setMenus] = useState<Menu[]>([])
  const [reservationDate, setReservationDate] = useState<string>("")
  const [allSelected, setAllSelected] = useState<boolean>(false)

  const handleSetSelectValue = (value: Beautician | Salon | Menu[] | string) => {
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
  const geterSelectID = (type: "beautician" | "store" | "date" | "menu"): string | string[] | undefined => {
    switch (type) {
      case "beautician":
        if (beautician !== initBeautician) {
          return beautician.randId;
        } else {
          return undefined
        }
        break
      case "store":
        if (store !== initSalon) {
          return store.randId;
        } else {
          return undefined
        }
        break
      case "date":
        if (reservationDate !== "") {
          return reservationDate;
        } else {
          return undefined
        }
        break
      case "menu":
        const menuIDs: string[] = []
        if (Object.prototype.toString.call(menus) === "[object Array]" && Array.isArray(menus)) {
          menus.map((menu) => {
            menuIDs.push(menu.randId)
          })
        }
        return menuIDs
        break
    }
  }
  const geterSelectValue = (): [Beautician, Salon, Menu[], string] => {
    return [beautician, store, menus, reservationDate]
  }

  useEffect(() => {
    // let allSelected: boolean = false
    // const allSelectedCheck = () => {
    //   if (beautician !== undefined && store !== undefined && menus.length !== 0 && reservationDate !== "") {
    // setAllSelected(true)
    // allSelected = true
    //     history.push(match.path + "/final_comfirmation")
    //   }
    // }
    // const jumpFinalPage = (allSelected: boolean) => {
    //   if (allSelected){
    //     history.push(match.path + "/final_comfirmation")
    //   } else {
    //     console.log(allSelected)
    //   }
    // }
    // allSelectedCheck()
    // jumpFinalPage(allSelected)
    // history.push("guest/final_comfirmation")
  }, [beautician, store, menus, reservationDate])

  return (
    <SetSelectValueContext.Provider value={handleSetSelectValue}>
      <GeterSelectIDContext.Provider value={geterSelectID}>
        <GeterSelectValueContext.Provider value={geterSelectValue}>
          <Router>
            <Switch>
              <Route exact path={match.path} render={() => <GuestComponent allSelectCheck={allSelected} />} />
              <Route exact path={match.path + "/beautician"} render={() => <ChooseBeautician />} />
              <Route exact path={match.path + "/store"} render={() => <ChooseStore />} />
              <Route exact path={match.path + "/date"} render={() => <ChooseDate />} />
              <Route exact path={match.path + "/menu"} render={() => <ChooseMenu />} />
              <Route exact path={match.path + "/final_comfirmation"} component={FinalComfirmation} />
            </Switch>
          </Router>
        </GeterSelectValueContext.Provider>
      </GeterSelectIDContext.Provider>
    </SetSelectValueContext.Provider>
  )
}

export default Guest;