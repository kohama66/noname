import React, { FC } from 'react';
import { useHistory } from 'react-router-dom';
import Content from './Content';

const Home: FC = () => {
  const history = useHistory()
  return (
    <section id="home">
      <article className="home-image">
        <figure>
          <img src="img/Barber-bro.png" alt="" />
        </figure>
        <div>
          <h1>もっと自由な美容師へ</h1>
          <p>Cut Matchは髪を切りたい人と美容師さんを<br />マッチングさせるサービスです。<br />
          今までの美容院予約サービスより<br />お手軽に好みに合わせて予約して頂けます。</p>
        </div>
        <button onClick={() => history.push("/guest")}>
          予約する
      </button>
      </article>
      {/* <div>
      <h1>Cut Match</h1>
      <p>Cut Matchは美容室と美容師と髪を切りたい人をマッチングさせるサービスです。<br />
          美容師は特定の店舗に拘らずにお仕事をする事が可能です。
        </p>
    </div>
    <div className="home-contents">
      <Content title='髪を切りたい方はコチラ' path="guest"/>
      <Content title='美容師の方はコチラ' path="beautician"/>
      <Content title='美容院の方はコチラ' path=""/>
    </div> */}
    </section>
  )
};

export default Home;
