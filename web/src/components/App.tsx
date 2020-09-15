import React, { FC } from 'react';
import './App.scss';
import {Home, Beautician, Guest} from './endpoint'
// import { getReservationBeautician } from '../package/api/index';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

const App: FC = () => (
  <>
    <header></header>
    <section id="top">
      <section className="inner">
        <Router>
          <Switch>
            <Route path="/" exact component={Home} />
            <Route path="/beautician" component={Beautician} />
            <Route path="/guest" component={Guest} />
          </Switch>
        </Router>
      </section>
    </section>
  </>
);

export default App;
