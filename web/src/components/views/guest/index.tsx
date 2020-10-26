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
import { Beautician, isBeauticianInterface } from '../../../package/interface/Beautician';
import { Salon } from '../../../package/interface/Salon';

export const SetSelectContext = createContext((id: string | string[], type: "beautician" | "store" | "date" | "menu", content?: Beautician | Salon) => { });
export const GeterSelectIDContext = createContext((type: "beautician" | "store" | "date" | "menu"): string | string[] => "")

const Guest: FC = () => {
  const match = useRouteMatch()
  const [beauticianID, setBeauticainID] = useState<string>("")
  const [storeID, setStoreID] = useState<string>("")
  const [menuIDs, setMenuIDs] = useState<string[]>([])
  const [date, setDate] = useState<string>("")

  // const [beautician, setBeauticain] = useState<Beautician>()
  // const [store, setStore] = useState<Salon>()
  // const [menus, setMenus] = useState<Menu[]>([])

  // const handleSetSelect = (type: "beautician" | "store" | "date" | "menu", content: Beautician | Salon | Menu[] | string): void => {
  //   switch (type) {
  //     case "beautician":
  //       if (isBeauticianInterface(content)) {
  //         setBeauticain(content)
  //       }
  //       break
  //     case "store":
  //       if (isSalonInterface(content)) {
  //         setStore(content)
  //       }
  //       break
  //     case "date":
  //       if (typeof content === "string") {
  //         setDate(content)
  //       }
  //       break
  //     case "menu":
  //       if (isMenusInterface(content)) {
  //         setMenus(content)
  //       }
  //       break
  //   }
  // }

  const handleSetSelect = (id: string | string[], type: "beautician" | "store" | "date" | "menu"): void => {
    switch (type) {
      case "beautician":
        if (typeof id === "string") {
          setBeauticainID(id)
        }
        break
      case "store":
        if (typeof id === "string") {
          setStoreID(id)
        }
        break
      case "date":
        if (typeof id === "string") {
          console.log(id)
          setDate(id)
        }
        break
      case "menu":
        if (typeof id !== "string") {
          setMenuIDs(id)
        }
        break
    }
  }

  const geterSelectID = (type: "beautician" | "store" | "date" | "menu"): string | string[] => {
    switch (type) {
      case "beautician":
        return beauticianID;
        break
      case "store":
        return storeID;
        break
      case "date":
        return date;
        break
      case "menu":
        return menuIDs;
        break
    }
  }

  return (
    <SetSelectContext.Provider value={handleSetSelect}>
      <GeterSelectIDContext.Provider value={geterSelectID}>
        <Router>
          <Switch>
            {/* <Route exact path={match.path} render={() => <GuestComponent checkedBeautician={checkedBeautician} checkedStore={checkedStore} checkedMenu={checkedMenu} checkedDate={checkedDate} />} /> */}
            <Route exact path={match.path} component={GuestComponent} />
            <Route exact path={match.path + "/beautician"} render={() => <ChooseBeautician />} />
            <Route exact path={match.path + "/store"} render={() => <ChooseStore />} />
            <Route exact path={match.path + "/date"} render={() => <ChooseDate />} />
            <Route exact path={match.path + "/menu"} render={() => <ChooseMenu />} />
            <Route exact path={match.path + "/final_comfirmation"} component={FinalComfirmation} />
          </Switch>
        </Router>
      </GeterSelectIDContext.Provider>
    </SetSelectContext.Provider>
  )
}

export default Guest;