import React, { FC } from 'react';
import Title from '../../guest/parts/Title/Title';
import "./ReservationVerify.scss";
import * as H from 'history'
import { useLocation } from 'react-router-dom';
import { Menu } from '../../../../package/interface/Menu';

const ReservationVerify: FC = () => {
  const location: H.Location<{
    userName?: string,
    userNameKana?: string
    userPhoneNmber?: string
    salonName?: string
    salonAddress?: string
    reservedDate?: string
    reservedTime?: string
    menus?: Menu[]
  }> = useLocation()

  return (
    <div id="reservation-verify">
      <Title title="RESERVED" text="予約確認" />
      <div className="reservation-verify-contents">
        <div className="reservation-verify-content">
          <h2>名前</h2>
          <p>{location.state.userName}<span>{location.state.userNameKana}</span></p>
          <h2>電話番号</h2>
          <p>{location.state.userPhoneNmber ? location.state.userPhoneNmber : "未設定"}</p>
        </div>
        <div className="reservation-verify-content">
          <h2>美容院</h2>
          <p>{location.state.salonName}</p>
          <h2>住所</h2>
          <p>{location.state.salonAddress}</p>
        </div>
        <div className="reservation-verify-content">
          <h2>日付</h2>
          <p>{location.state.reservedDate}</p>
          <h2>時間</h2>
          <p>{location.state.reservedTime}</p>
          <p></p>
          <h2>メニュー</h2>
          <ul>
            {location.state.menus?.map((menu) => {
              return <li>{menu.name}</li>
            })}
          </ul>
          <h2>合計</h2>
          <p>4500</p>
        </div>
      </div>
    </div>
  )
}

export default ReservationVerify;