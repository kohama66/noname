import React, { FC, useState, createContext } from 'react';
import GuestComponent from '../../../components/views/guest/Guest'
import ChooseBeautician from "./beautician"
import ChooseStore from '../guest/store'
import ChooseDate from '../../../components/views/guest/date/ChooseDate'
import ChooseMenu from '../../../components/views/guest/menu/ChooseMenu'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
  Link,
  useHistory,
} from "react-router-dom";

export const setCheckedContext = createContext((checkName: string) => { });
export const SelectContext = createContext((id: string, type: "beautician" | "store") => { });
export const BeauticainIDContext = createContext<string>("");
export const GeterSelectCpntext = createContext((type: "beautician" | "store" | "date" | "menu"): string => "")

const Guest: FC = () => {
  const match = useRouteMatch();

  const [checkedBeautician, setCheckedBeautician] = useState(false)
  const [checkedStore, setCheckedStore] = useState(false)
  const [checkedMenu, setCheckedMenu] = useState(false)
  const [checkedDate, setCheckedDate] = useState(false)
  const [beauticianID, setBeauticainID] = useState<string>("")
  const [storeID, setStoreID] = useState<string>("")

  const handleSelect = (id: string, type: "beautician" | "store") => {
    switch (type) {
      case "beautician":
        setBeauticainID(id)
        break
      case "store":
        setStoreID(id)
        break
    }
  }
  const geterSelect = (type: "beautician" | "store" | "date" | "menu"): string => {
    switch (type) {
      case "beautician":
        return beauticianID;
        break
      case "store":
        return storeID;
        break
      case "date":
        return ""
        break
      case "menu":
        return ""
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
      <SelectContext.Provider value={handleSelect}>
        <BeauticainIDContext.Provider value={beauticianID}>
          <GeterSelectCpntext.Provider value={geterSelect}>
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
          </GeterSelectCpntext.Provider>
        </BeauticainIDContext.Provider>
      </SelectContext.Provider>
    </setCheckedContext.Provider>
  )
}

export default Guest;