import React, { FC } from 'react'
import { Reservation } from '../../../../package/interface/Reservation'
import Accordion from '../../parts/accordion/Accordion'
import Schedule from '../../parts/Schedule/Schedule'
import "./MyPage.scss"

const MyPage: FC = () => {
  const ttt: Reservation[] = []
  return (
    <div id="mypage">
      <div className="contents">
        <figure>
          <img src="/img/salan.jpg" alt="" />
        </figure>
        <div className="profile">
          <div>
            <h1>TOKYO SALON</h1>
            <span>
              <h2>PHONE</h2>
              <p>09012341234</p>
            </span>
            <span>
              <h2>ADDRESS</h2>
              <p>東京都渋谷区新町15</p>
            </span>
            <span>
              <h2>OPEN</h2>
              <p>09:00</p>
            </span>
            <span>
              <h2>CLOSE</h2>
              <p>21:00</p>
            </span>
          </div>
          <button>変更</button>
        </div>
        <div className="beauticians">
          <Accordion buttonText="STYLIST ▼">
            <ul>
              <li>山田 太郎</li>
              <li>田中 真也</li>
            </ul>
          </Accordion>
        </div>
        <div className="reservation">
          <Schedule reservations={ttt} />
        </div>
      </div>
    </div>
  )
}

export default MyPage
