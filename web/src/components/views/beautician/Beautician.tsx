import React, { FC } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch,
} from "react-router-dom";

const Beautician: FC = () => {
  const match = useRouteMatch();
  return (
    <div id="beautician">
      <Router>
        <Switch>
          <Route path={match.path + "/aaa"} component={signIn} />
          {/* <Route path={match.path} component={Schedule} /> */}
        </ Switch>
      </Router>
    </div>
  )
};

const signIn: FC = () => {
  return (
    <section className="inner">
      <figure>
        <img src="../img/beautician_1.jpg" alt="" />
        <div>
          {/* <p>Cut Matchは</p> */}
        </div>
      </figure>
      <form className="signIn">
        <label>メールアドレス</label>
        <input type="mail" />
        <label>パスワード</label>
        <input type="password" />
        <button>ログイン</button>
      </form>
    </section>
  )
};



export default Beautician;
