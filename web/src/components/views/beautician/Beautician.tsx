import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import BtAuth from '../../container/btAuth/BtAuth';
import ChangeProfile from './changeProfile/ChangeProfile';
import Mypage from './mypage/Mypage';
import SignUp from './signup/SignUp';

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Switch>
        <Route path={match.path + "/signup"} component={SignUp} />
        <BtAuth >
          <Route path={match.path + "/mypage"} component={Mypage} />
          <Route path={match.path + "/changeprofile"} component={ChangeProfile} />
        </BtAuth>
      </ Switch>
    </div>
  )
};

export default Beautician;