import React, { FC } from 'react';
import { isBeautician } from '../../../../../package/interface/Beautician';
import { Salon } from '../../../../../package/interface/Salon';
import { User } from '../../../../../package/interface/User';
import { BeauticianCardDetail, isStoreCardDetailProps, StoreCardDetail } from "./CardDetail"

interface props {
  type: "store" | "beautician"
  content: Salon | User
}

const CardDetail: FC<props> = (props) => {
  if (isStoreCardDetailProps(props.content)) {
    return <StoreCardDetail name={props.content.name} phoneNumber={props.content.phoneNumber} openingHours={props.content.openingHours} closingHours={props.content.closingHours}
      postalCode={props.content.postalCode} prefectures={props.content.prefectures} city={props.content.city}
      town={props.content.town} addressCode={props.content.addressCode} addressOther={props.content.addressOther} />
  } else if (isBeautician(props.content)) {
    return <BeauticianCardDetail firstName={props.content.firstName} lasttName={props.content.lastName}
      phoneNumber={props.content.phoneNumber} lineId={props.content.beauticianInfo.lineId}
      instagramId={props.content.beauticianInfo.instagramId} menus={props.content.beauticianMenus ? props.content.beauticianMenus : []} />
  }
  return (
    <div>???</div>
  )
}

export default CardDetail;
