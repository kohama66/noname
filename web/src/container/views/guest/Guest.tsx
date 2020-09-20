import React, { FC, useState, createContext } from 'react';
import GuestComponent from '../../../components/views/guest/Guest'
import ChooseBeautician from '../../../container/views/guest/beautician/ChooseBeautician'
import ChooseStore from '../../../components/views/guest/store/ChooseStore'
import ChooseDate from '../../../components/views/guest/date/ChooseDate'
import ChooseMenu from '../../../components/views/guest/menu/ChooseMenu'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
  Link,
} from "react-router-dom";

export const setCheckedContext = createContext((checkName: string) => { });

const Guest: FC = () => {
  const [checkedBeautician, setCheckedBeautician] = useState(false)
  const [checkedStore, setCheckedStore] = useState(false)
  const [checkedMenu, setCheckedMenu] = useState(false)
  const [checkedDate, setCheckedDate] = useState(false)
  const handleChecked = (checkName: string) => {
    switch (checkName) {
      case "store":
        setCheckedStore(true)
      case "beautician":
        setCheckedBeautician(true)
      case "menu":
        setCheckedMenu(true)
      case "date":
        setCheckedDate(true)
    }
  }
  const match = useRouteMatch();
  return (
    <setCheckedContext.Provider value={handleChecked} >
      <Router>
        <Switch>
          <Route exact path={match.path} render={() => <GuestComponent checkedBeautician={checkedBeautician} checkedStore={checkedStore} checkedMenu={checkedMenu} checkedDate={checkedDate} />} />
          <Route exact path={match.path + "/beautician"} render={() => <ChooseBeautician />} />
          <Route exact path={match.path + "/store"} render={() => <ChooseStore />} />
          <Route exact path={match.path + "/date"} render={() => <ChooseDate />} />
          <Route exact path={match.path + "/menu"} render={() => <ChooseMenu />} />
        </Switch>
      </Router>
    </setCheckedContext.Provider>
  )
}

export default Guest;