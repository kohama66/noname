import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import Title from '../parts/Title'
import { setCheckedContext } from '../../guest'
import './ChooseStore.scss'

const ChooseStore: FC = () => {
  const history = useHistory()
  const setChecked = useContext(setCheckedContext)

  const handleStore = () => {
    setChecked("store")
    history.push("/guest")
  }
  return (
    <section id="choose-store">
      <Title titleText={"店舗を選ぶ"} image={"/img/thinkmen.png"} />
      <div className="choose-store-contents">
        <h2 className="sub-title">BEAUTICIAN</h2>
        <div className="choose-store-content-wrapper">
          <article className="choose-store-content" onClick={handleStore}>
            <figure>
              <img src="/img/salan.jpg" alt="" />
            </figure>
            <div>
              <h2>KYOTO-CUT</h2>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </article>
          <article className="choose-store-content" onClick={handleStore}>
            <figure>
              <img src="/img/salan.jpg" alt="" />
            </figure>
            <div>
              <h2>KYOTO-CUT</h2>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </article>
          <article className="choose-store-content" onClick={handleStore}>
            <figure>
              <img src="/img/salan.jpg" alt="" />
            </figure>
            <div>
              <h2>KYOTO-CUT</h2>
              <h3>京都駅から徒歩5分</h3>
              <p>京都市下京区東塩小路10丁目...</p>
            </div>
          </article>
        </div>
      </div>
    </section>
  )
}

export default ChooseStore;