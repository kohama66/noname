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
    return <BeauticianCardDetail firstName={props.content.firstName} lasttName={props.content.lastName}
    phoneNumber={props.content.phoneNumber} lineId={props.content.lineId} comment={props.content.comment}
    instagramId={props.content.instagramId} menus={props.content.menus} />
  }
  return (
    <div>???</div>
  )
}

export default CardDetail;
