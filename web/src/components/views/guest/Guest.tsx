import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import FinalComfirmation from '../finalComfirmation';
import Mypage from './mypage/Mypage';
import SignUp from './signup/SignUp';
import GuestHome from './home/GuestHome';
import ChooseStore from './store/ChooseStore';
import Auth from '../../container/auth/Auth';
import ChooseBeautician from './beautician/ChooseBeautician';
import ChooseDate from './date/ChooseDate';
import ChooseMenu from './menu/ChooseMenu';

const Guest: FC = () => {
  const match = useRouteMatch().path

  return (
    <Switch>
      <Route exact path={match} component={GuestHome} />
      <Route exact path={match + "/beautician"} component={ChooseBeautician} />
      <Route exact path={match + "/store"} component={ChooseStore} />
      <Route exact path={match + "/date"} component={ChooseDate} />
      <Route exact path={match + "/menu"} component={ChooseMenu} />
      <Route exact path={match + "/final_comfirmation"} component={FinalComfirmation} />
      <Auth>
        <Route exact path={match + "/mypage"} component={Mypage} />
      </Auth>
      <Route exact path={match + "/signup"} component={SignUp} />
    </Switch>
  )
}

export default Guest;