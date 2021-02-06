import React, { FC } from 'react';
import './App.scss';
import { Home, Beautician, Guest } from './endpoint'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Header from './views/parts/header';
import Login from './views/login/Login';
import Reserved from './views/reserved/Reserved';
import Salon from './views/salon/Salon';

const App: FC = () => {
  return (
    <>
      <section id="top">
        <section className="inner">
          <Router>
            <Header />
            <Switch>
              <Route path="/" exact component={Home} />
              <Route path="/beautician" component={Beautician} />
              <Route path="/guest" component={Guest} />
              <Route path="/login" component={Login} />
              <Route path="/reserved" component={Reserved} />
              <Route path="/salon" component={Salon} />
            </Switch>
          </Router>
        </section>
      </section>
    </>
  )
};

export default App;
