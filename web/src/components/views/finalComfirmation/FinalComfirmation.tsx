import React, { FC } from 'react';
import { Beautician } from '../../../package/interface/Beautician';
import { MenuDetail } from '../../../package/interface/Menu';
import { Salon } from '../../../package/interface/Salon';
import ChooseCard from '../guest/parts/ChooseCard';
import Title from '../guest/parts/Title/Title';
import "./FinalComfirmation.scss"

interface props {
  beautician: Beautician
  store: Salon
  date: string
  menus: MenuDetail[]
}

const FinalComfirmation: FC<props> = (props) => {
  return (
    <article id="final-comfirmation">
      <Title text="最終確認" title="FINAL COMFIRMATION" />
      <div className="final-answer">
        <h2>以下で予約を完了します</h2>
        <button>
          予約確定
        </button>
      </div>
      <div className="final-comfirmation-content-wrapper">
        <ChooseCard type="beautician" content={props.beautician} />
        <ChooseCard type="store" content={props.store} />
        <div className="final-comfirmation-content">
          <h2>日付</h2>
          <div>
            <p>{props.date}</p>
            {/* <p>2020年10-25 12:00:00</p> */}
          </div>
          <h2>メニュー</h2>
          <div>
            <dl>
              {props.menus.map((menu, id) => {
                return <>
                  <dt>{menu.name}</dt>
                  <dd>{menu.price}</dd>
                </>
              })}
              {/* <dt>カット</dt>
              <dd>1200</dd>
              <dt>パーマ</dt>
              <dd>3500</dd>
              <dt>カラー</dt>
              <dd>1800</dd> */}
            </dl>
            <div className="total-price">
              <h2>合計</h2>
              {/* <p>5500</p> */}
            </div>
          </div>
        </div>
      </div>
    </article>
  )
}

export default FinalComfirmation;