import React, { FC } from 'react';
import { useHistory } from 'react-router-dom';

const Home: FC = () => {
  const history = useHistory()
  return (
    <section id="home">
      <article className="home-image">
        <div>
          <h1>もっと自由な美容師へ</h1>
          <p>Cut Matchは髪を切りたい人と美容師さんを<br />マッチングさせるサービスです。<br />
          今までの美容院予約サービスより<br />お手軽に好みに合わせて予約して頂けます。</p>
        </div>
        <button onClick={() => history.push("/guest")}>予約する</button>
      </article>
    </section>
  )
};

export default Home;
