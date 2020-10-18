import React, { FC } from 'react';
import { Salon } from '../../../../../package/interface/Salon';
import { BeauticianCardDetail, StoreCardDetail } from "./CardDetail"

interface props {
  type: "store" | "beautician"
  content: Salon
}

const CardDetail: FC<props> = (props) => {
  if (props.type == "store") {
    return <StoreCardDetail name={props.content.name} phoneNumber="075-353-5390" businessHours="営業時間: 09:00 ~ 21:00"
      postalCode="〒600-8216" address="京都府京都市下京区東塩小路町５５７−１ Station７階" />
  } else if (props.type == "beautician") {
    return <BeauticianCardDetail name="山田 太郎" phoneNumber="09012345678" lineId="TEST" comemnt="お客様に似合う最適なスタイルを提供致します！" instaLink="test" />
  }
  return (
    <></>
  )
}

export default CardDetail;