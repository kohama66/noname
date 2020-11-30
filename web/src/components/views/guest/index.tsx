import React, { FC, useState, createContext, useEffect } from 'react';
import GuestComponent from './home'
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
import Mypage from './mypage/Mypage';
import SignUp from './signup/SignUp';
import Login from '../login/Login';

const Guest: FC = () => {
  const match = useRouteMatch()

  return (
    <Switch>
      <Route exact path={match.path} render={() => <GuestComponent />} />
      <Route exact path={match.path + "/beautician"} render={() => <ChooseBeautician />} />
      <Route exact path={match.path + "/store"} render={() => <ChooseStore />} />
      <Route exact path={match.path + "/date"} render={() => <ChooseDate />} />
      <Route exact path={match.path + "/menu"} render={() => <ChooseMenu />} />
      <Route exact path={match.path + "/final_comfirmation"} component={FinalComfirmation} />
      <Route exact path={match.path + "/mypage"} component={Mypage} />
      <Route exact path={match.path + "/signup"} component={SignUp} />
      <Route exact path={match.path + "/login"} component={Login} />
    </Switch>
  )
}

export default Guest;