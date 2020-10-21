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

export const setCheckedContext = createContext((checkName: string) => { });
export const SetSelectContext = createContext((id: string | string[], type: "beautician" | "store" | "date" | "menu") => { });
export const GeterSelectIDContext = createContext((type: "beautician" | "store" | "date" | "menu"): string | string[] => "")

const Guest: FC = () => {
  const match = useRouteMatch();

  const [checkedBeautician, setCheckedBeautician] = useState(false)
  const [checkedStore, setCheckedStore] = useState(false)
  const [checkedMenu, setCheckedMenu] = useState(false)
  const [checkedDate, setCheckedDate] = useState(false)
  const [beauticianID, setBeauticainID] = useState<string>("")
  const [storeID, setStoreID] = useState<string>("")
  const [menuIDs, setMenuIDs] = useState<string[]>([])
  const [date, setDate] = useState<string>("")

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

  const handleChecked = (checkName: string) => {
    switch (checkName) {
      case "store":
        setCheckedStore(true)
        break
      case "beautician":
        setCheckedBeautician(true)
        break
      case "menu":
        setCheckedMenu(true)
        break
      case "date":
        setCheckedDate(true)
        break
    }
  }
  return (
    <setCheckedContext.Provider value={handleChecked} >
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
            </Switch>
          </Router>
        </GeterSelectIDContext.Provider>
      </SetSelectContext.Provider>
    </setCheckedContext.Provider>
  )
}

export default Guest;