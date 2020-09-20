import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import Title from '../parts/Title'
import { setCheckedContext } from '../../../../container/views/guest/Guest'

const ChooseStore: FC = () => {
  const history = useHistory()
  const setChecked = useContext(setCheckedContext)

  const handleStore = () => {
    setChecked("store")
    history.push("/guest")
  }
  return (
    <section id="choose-store">
      <Title titleText={"店舗を選ぶ"} image={"/img/choose1.png"} />
      <div className="choose-store-contents">
        <article className="choose-store-content" onClick={handleStore}>
          <figure>
            <img src="/img/salan.jpg" alt="" />
          </figure>
          <div>
            <h2>KYOTO-CUT</h2>
            <div>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </div>
        </article>
        <article className="choose-store-content">
          <figure>
            <img src="/img/salan.jpg" alt="" />
          </figure>
          <div>
            <h2>KYOTO-CUT</h2>
            <div>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </div>
        </article>
        <article className="choose-store-content">
          <figure>
            <img src="/img/salan.jpg" alt="" />
          </figure>
          <div>
            <h2>KYOTO-CUT</h2>
            <div>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </div>
        </article>
        <article className="choose-store-content">
          <figure>
            <img src="/img/salan.jpg" alt="" />
          </figure>
          <div>
            <h2>KYOTO-CUT</h2>
            <div>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </div>
        </article>
      </div>
    </section>
  )
}

export default ChooseStore;