import React, { FC } from 'react';
import { Beautician } from '../../../../../package/interface/Beautician';
import { Salon } from '../../../../../package/interface/Salon';
import { BeauticianCardDetail, isBeauticianCardDetail, isStoreCardDetailProps, StoreCardDetail, StoreCardDetailProps } from "./CardDetail"

interface props {
  type: "store" | "beautician"
  content: Salon | Beautician
}

const CardDetail: FC<props> = (props) => {
  if (isStoreCardDetailProps(props.content)) {
    return <StoreCardDetail name={props.content.name} phoneNumber={props.content.phoneNumber} openingHours={props.content.openingHours} closingHours={props.content.closingHours}
      postalCode={props.content.postalCode} prefectures={props.content.prefectures} city={props.content.city}
      town={props.content.town} addressCode={props.content.addressCode} addressOther={props.content.addressOther} />
  } else if (isBeauticianCardDetail(props.content)) {
    return <BeauticianCardDetail name="山田 太郎" phoneNumber="09012345678" lineId="TEST" comemnt="お客様に似合う最適なスタイルを提供致します！" instaLink="test" />
  }
  return (
    <div>???</div>
  )
}

export default CardDetail;
