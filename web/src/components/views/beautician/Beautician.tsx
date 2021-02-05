import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import Auth from '../../container/auth/Auth';
import BtAuth from '../../container/btAuth/BtAuth';
import ChangeProfile from './changeProfile/ChangeProfile';
import Mypage from './mypage/Mypage';
import ReservationVerify from './reservedInfo/ReservationVerify';
import SignUp from './signup/SignUp';

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Switch>
        <Route path={match.path + "/signup"} component={SignUp} />
        <BtAuth >
          <Auth>
            <Route path={match.path + "/mypage"} component={Mypage} />
            <Route path={match.path + "/changeprofile"} component={ChangeProfile} />
            <Route path={match.path + "/reservationverify"} component={ReservationVerify} />
          </Auth>
        </BtAuth>
      </ Switch>
    </div>
  )
};

export default Beautician;