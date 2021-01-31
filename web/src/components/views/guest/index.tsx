import React, { FC, useState, createContext, useEffect } from 'react';
import ChooseBeautician from "./beautician"
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
import GuestHome from './home/Guest';
import ChooseStore from './store/ChooseStore';
import Auth from '../../container/auth/Auth';

const Guest: FC = () => {
  const match = useRouteMatch()

  return (
    <Switch>
      <Route exact path={match.path} component={GuestHome} />
      <Route exact path={match.path + "/beautician"} component={ChooseBeautician} />
      <Route exact path={match.path + "/store"} component={ChooseStore} />
      <Route exact path={match.path + "/date"} component={ChooseDate} />
      <Route exact path={match.path + "/menu"} component={ChooseMenu} />
      <Route exact path={match.path + "/final_comfirmation"} component={FinalComfirmation} />
      <Auth>
        <Route exact path={match.path + "/mypage"} component={Mypage} />
      </Auth>
      <Route exact path={match.path + "/signup"} component={SignUp} />
    </Switch>
  )
}

export default Guest;