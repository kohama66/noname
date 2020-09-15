import React, { Component, FC } from 'react';
import ChooseBeautician from '../../../container/views/guest/ChooseBeautician'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
  Link,
} from "react-router-dom";

const Guest: FC = () => {
  const match = useRouteMatch();
  return (
    <section id="guest">
      <Router>
        <Switch>
          <Route exact path={match.path} component={GuestHome} />
          <Route exact path={match.path + "/beautician"} component={ChooseBeautician} />
        </Switch>
      </Router>
    </section>
  )
}

const GuestHome: FC = () => {
  const match = useRouteMatch();
  return (
    <article>
      <Link to={match.path + "/beautician"}><div>美容師を選ぶ</div></Link>
      <div>日付を選ぶ</div>
      <div>お店で選ぶ</div>
      <div>メニューで選ぶ</div>
    </article>
  )
}

export default Guest;
