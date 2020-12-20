import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import ChangeProfile from './changeProfile/ChangeProfile';
import Mypage from './mypage';
import SignUp from './signup/SignUp';

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Switch>
        <Route path={match.path + "/signup"} component={SignUp} />
        <Route path={match.path + "/mypage"} component={Mypage} />
        <Route path={match.path + "/changeprofile"} component={ChangeProfile} />
      </ Switch>
    </div>
  )
};

export default Beautician;