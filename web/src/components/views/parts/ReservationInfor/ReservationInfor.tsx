import React, { FC } from 'react';
import { GuestByMyPage } from '../../../../package/interface/Guest';

interface props {
  titleText: string
  storeName: string
  beauticianLastName: string
  beauticianFirstName: string
}

const ReservationInfor: FC<props> = (props) => {
  return (
    <div id="reservationInfor">
      <h2>{props.titleText}</h2>
      <div>
        <dl>
          <span>
            <dt>店舗</dt>
            <dd>{props.storeName}</dd>
          </span>
          <span>
            <dt>美容師</dt>
            <dd>{props.beauticianLastName + " " + props.beauticianFirstName}</dd>
          </span>
          <span>
            <dt>日付</dt>
            <dd>11月05日</dd>
          </span>
          <span>
            <dt>時間</dt>
            <dd>15:00から</dd>
          </span>
          <span className="menus">
            <dt>メニュー</dt>
            <span>
              <dd>カット</dd>
              <dd>カラー</dd>
              <dd>パーマ</dd>
            </span>
          </span>
          <span>
            <dt>合計</dt>
            <dd>5500</dd>
          </span>
        </dl>
      </div>
    </div>
  )
}

export default ReservationInfor;