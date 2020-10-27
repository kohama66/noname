import React, { FC, useState, createContext } from 'react';
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
} from "react-router-dom";
import FinalComfirmation from '../finalComfirmation';
import { Beautician, initBeautician, isBeauticianInterface } from '../../../package/interface/Beautician';
import { initSalon, isSalonInterface, Salon } from '../../../package/interface/Salon';
import { Menu } from '../../../package/interface/Menu';

export const SetSelectValueContext = createContext((value: Beautician | Salon | Menu[] | string) => { });
export const GeterSelectIDContext = createContext((type: "beautician" | "store" | "date" | "menu"): string | string[] => "")

const Guest: FC = () => {
  const match = useRouteMatch()

  const [beautician, setBeautician] = useState<Beautician>(initBeautician)
  const [store, setStore] = useState<Salon>(initSalon)
  const [menus, setMenus] = useState<Menu[]>([])
  const [reservationDate, setReservationDate] = useState<string>("")

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
    // if (beautician !== undefined && store !== undefined && menus.length !== 0 && reservationDate !== "") {
    //   console.log("kkk")
    // }
  }

  // const hd = () => {
  //   if (beautician !== undefined && store !== undefined && menus.length !== 0 && reservationDate !== "") {
  //     console.log("kkk")
  //   }
  // }

  const geterSelectID = (type: "beautician" | "store" | "date" | "menu"): string | string[] => {
    switch (type) {
      case "beautician":
        if (beautician !== initBeautician) {
          return beautician.randId;
        } else {
          return ""
        }
        break
      case "store":
        if (store !== initSalon) {
          return store.randId;
        } else {
          return ""
        }
        break
      case "date":
        return reservationDate;
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

  return (
    <SetSelectValueContext.Provider value={handleSetSelectValue}>
      <GeterSelectIDContext.Provider value={geterSelectID}>
        <Router>
          <Switch>
            <Route exact path={match.path} component={GuestComponent} />
            <Route exact path={match.path + "/beautician"} render={() => <ChooseBeautician />} />
            <Route exact path={match.path + "/store"} render={() => <ChooseStore />} />
            <Route exact path={match.path + "/date"} render={() => <ChooseDate />} />
            <Route exact path={match.path + "/menu"} render={() => <ChooseMenu />} />
            <Route exact path={match.path + "/final_comfirmation"} component={FinalComfirmation} />
          </Switch>
        </Router>
      </GeterSelectIDContext.Provider>
    </SetSelectValueContext.Provider>
  )
}

export default Guest;