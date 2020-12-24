import React, { FC } from 'react';
import { MenuDetail } from '../../../package/interface/Menu';
import { Salon } from '../../../package/interface/Salon';
import { User } from '../../../package/interface/User';
import ChooseCard from '../guest/parts/ChooseCard';
import Title from '../guest/parts/Title/Title';
import "./FinalComfirmation.scss"

interface props {
  beautician: User
  store: Salon
  date: string
  menus: MenuDetail[]
  totalPrice: number
  handleReserve: () => void
  error?: string
}

const FinalComfirmation: FC<props> = (props) => {
  return (
    <article id="final-comfirmation">
      <Title text="最終確認" title="FINAL COMFIRMATION" />
      <div className="final-answer">
        <h2>以下で予約を完了します</h2>
        <button onClick={() => props.handleReserve()}>
          予約確定
        </button>
        <p> {props.error}</p>
      </div>
      <div className="final-comfirmation-content-wrapper">
        <ChooseCard type="beautician" content={props.beautician} image="/img/beautician_1.jpg" />
        <ChooseCard type="store" content={props.store} image="/img/salan.jpg" />
        <div className="final-comfirmation-content">
          <h2>日付</h2>
          <div>
            <p>{props.date}</p>
          </div>
          <h2>メニュー</h2>
          <div>
            <dl>
              {props.menus.map((menu, i) => {
                return <span>
                  <dt>{menu.name}</dt>
                  <dd>{menu.price}</dd>
                </span>
              })}
            </dl>
          </div>
          <div className="total-price">
            <h2>合計</h2>
            <p>{props.totalPrice}</p>
          </div>
        </div>
      </div>
    </article>
  )
}

export default FinalComfirmation;