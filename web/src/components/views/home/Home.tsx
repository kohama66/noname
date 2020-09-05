import React, { FC } from 'react';
import Content from './Content';

const Home: FC = () => (
  <>
    <div>
      <h1>Cut Match</h1>
      <p>Cut Matchは美容室と美容師と髪を切りたい人をマッチングさせるサービスです。<br />
          美容師は特定の店舗に拘らずにお仕事をする事が可能です。
        </p>
    </div>
    <div className="home-contents">
      <Content title='髪を切りたい方はコチラ'/>
      <Content title='美容師の方はコチラ'/>
      <Content title='美容院の方はコチラ'/>
    </div>
  </>
);

export default Home;
