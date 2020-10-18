import React, { FC } from 'react';
import { Salon } from '../../../../../package/interface/Salon';
import { BeauticianCardDetail, StoreCardDetail } from "./CardDetail"

interface props {
  type: "store" | "beautician"
  content: Salon
}

const CardDetail: FC<props> = (props) => {
  if (props.type == "store") {
    return <StoreCardDetail name={props.content.name} phoneNumber={props.content.phoneNumber} businessHours={`営業時間: ${props.content.openingHours} ~ ${props.content.closingHours}`}
      postalCode={`〒 ${props.content.postalCode}`}
      address={`${props.content.prefectures}${props.content.city}${props.content.town}${props.content.addressCode}${props.content.addressOther}`}
    />
  } else if (props.type == "beautician") {
    return <BeauticianCardDetail name="山田 太郎" phoneNumber="09012345678" lineId="TEST" comemnt="お客様に似合う最適なスタイルを提供致します！" instaLink="test" />
  }
  return (
    <></>
  )
}

export default CardDetail;