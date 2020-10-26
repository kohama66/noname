import React, { FC } from 'react';
import { Beautician } from '../../../package/interface/Beautician';
import { Salon } from '../../../package/interface/Salon';
import ChooseCard from '../guest/parts/ChooseCard';
import Title from '../guest/parts/Title/Title';
import "./FinalComfirmation.scss"

interface props {
  beautician: Beautician
  store: Salon
  date: string
}

const FinalComfirmation: FC<props> = (props) => {
  return (
    <article className="final-comfirmation">
      <Title text="最終確認" title="FINAL COMFIRMATION" />
      <div className="final-comfirmation-content-wrapper">
        <ChooseCard type="beautician" content={props.beautician} />
        <ChooseCard type="store" content={props.store} />
        <div className="final-comfirmation-content">
          <h2>メニュー</h2>
          <div>
            <ul>
              <li>カット</li>
              <li>カラー</li>
              <li>パーマ</li>
            </ul>
          </div>
        </div>
        <div className="final-comfirmation-content">
          <h2>日付</h2>
          <div>
            <p>2020-10-26 09:00:00</p>
          </div>
        </div>
      </div>
    </article>
  )
}

export default FinalComfirmation;