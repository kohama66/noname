import React, { FC } from 'react';
import logo from './logo.svg';
import './App.scss';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         {/* <img src={logo} className="App-logo" alt="logo" /> */}
//         <p>
//           Edit <code>src/App.tsx</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }
const App: FC = () => (
  <>
  <header></header>
  <section id="home">
    <section className="inner">
      <div>
        <h1>Cut Match</h1>
        <p>Cut Matchは美容室と美容師と髪を切りたい人をマッチングさせるサービスです。<br/>
          美容師は特定の店舗に拘らずにお仕事をする事が可能です。
        </p>
      </div>
      <div className="home-contents">
        <div>
          <img src="/img/cut.png" alt=""/>
          <h2>髪を切りたい人はコチラ</h2>
        </div>
        <div>
          <img src="/img/cut.png" alt=""/>
          <h2>美容師の方はコチラ</h2>
        </div>
        <div>
          <img src="/img/cut.png" alt=""/>
          <h2>美容室の方はコチラ</h2>
        </div>
      </div>
    </section>
  </section>
  </>
);

export default App;
