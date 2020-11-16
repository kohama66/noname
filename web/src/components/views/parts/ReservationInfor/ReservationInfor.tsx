import React, { FC, useEffect, useState } from 'react';
import { GuestByMyPage } from '../../../../package/interface/Guest';
import { MenuDetail } from '../../../../package/interface/Menu';

interface props {
  titleText: string
  storeName: string
  beauticianLastName: string
  beauticianFirstName: string
  month: RegExpMatchArray | null
  day: RegExpMatchArray | null
  hours: RegExpMatchArray | null
  menus: MenuDetail[]
}

const ReservationInfor: FC<props> = (props) => {
  const [totalPrice, setTotalPrice] = useState<number>(0)
  useEffect(() => {
    var price = 0
    props.menus.map((menu) => {
      price += menu.price
    })
    setTotalPrice(price)
  },[])
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
            <dd>{`${props.month}月${props.day}日`}</dd>
          </span>
          <span>
            <dt>時間</dt>
            <dd>{`${props.hours}時から`}</dd>
          </span>
          <span className="menus">
            <dt>メニュー</dt>
            <span>
              {props.menus.map((menu, i) => {
                return <dd key={i}>{menu.name}</dd>
              })}
            </span>
          </span>
          <span>
            <dt>合計</dt>
            <dd>{totalPrice}</dd>
          </span>
        </dl>
      </div>
    </div>
  )
}

export default ReservationInfor;