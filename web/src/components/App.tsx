import React, { FC } from 'react';
import './App.scss';
import Home from '../components/views/home/Home'

const App: FC = () => (
  <>
  <header></header>
  <section id="home">
    <section className="inner">
      <Home />
    </section>
  </section>
  </>
);

export default App;
