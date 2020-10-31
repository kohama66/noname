import React, { FC } from 'react';
import './App.scss';
import { Home, Beautician, Guest } from './endpoint'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Header from './views/parts/header';

const App: FC = () => (
  <>
    <section id="top">
      <section className="inner">
        <Router>
          <Header />
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
