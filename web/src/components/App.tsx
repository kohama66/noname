import React, { FC, useEffect } from 'react';
import './App.scss';
import {Home, Beautician} from '../components/index'
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
          </Switch>
        </Router>
      </section>
    </section>
  </>
);

export default App;
