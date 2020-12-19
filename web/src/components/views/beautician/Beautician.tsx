import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import Mypage from './mypage/Mypage';
import SignUp from './signup/SignUp';

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Switch>
        <Route path={match.path + "/signup"} component={SignUp} />
        <Route path={match.path + "/mypage"} component={Mypage} />
      </ Switch>
    </div>
  )
};

export default Beautician;