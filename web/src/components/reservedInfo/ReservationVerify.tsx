import React, { FC } from 'react';
import Title from '../views/guest/parts/Title/Title';
import "./ReservationVerify.scss";
import * as H from 'history'
import { useLocation } from 'react-router-dom';
import { Menu } from '../../package/interface/Menu';

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
          <h2>NAME</h2>
          <p>{location.state.userName}<span>{location.state.userNameKana}</span></p>
          <h2>PHONE-NUMBER</h2>
          <p>{location.state.userPhoneNmber ? location.state.userPhoneNmber : "未設定"}</p>
          <h2>SALON</h2>
          <p>{location.state.salonName}</p>
          <h2>SALON-ADDRESS</h2>
          <p>{location.state.salonAddress}</p>
          <h2>DATE</h2>
          <p>{location.state.reservedDate}</p>
          <h2>TIME</h2>
          <p>{location.state.reservedTime}</p>
          <p></p>
          <h2>MENUS</h2>
          <ul>
            {location.state.menus?.map((menu) => <li>{menu.name}</li>)}
          </ul>
          <h2>PRICE</h2>
          <p>4500</p>
        </div>
      </div>
    </div>
  )
}

export default ReservationVerify;