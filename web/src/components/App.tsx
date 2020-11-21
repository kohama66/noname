import React, { FC } from 'react';
import './App.scss';
import { Home, Beautician, Guest } from './endpoint'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Header from './views/parts/header';
import { getGuestContext, GuestContext } from '../utils/GuestContext';

const App: FC = () => {
  const rootGuestContext = getGuestContext()
  return (
    <>
      <section id="top">
        <section className="inner">
          <GuestContext.Provider value={rootGuestContext} >
            <Router>
              <Header />
              <Switch>
                <Route path="/" exact component={Home} />
                <Route path="/beautician" component={Beautician} />
                <Route path="/guest" component={Guest} />
              </Switch>
            </Router>
          </GuestContext.Provider>
        </section>
      </section>
    </>
  )
};

export default App;
