import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";
import SignUp from './signup/SignUp';

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Router>
        <Switch>
          <Route path={match.path + "/signup"} component={SignUp} />
        </ Switch>
      </Router>
    </div>
  )
};

export default Beautician;