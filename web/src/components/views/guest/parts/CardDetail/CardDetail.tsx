import React, { FC } from 'react';

export interface StoreCardDetailProps {
  addressCode: string
  addressOther: string
  city: string
  closingHours: string
  name: string
  openingHours: string
  phoneNumber: string
  postalCode: string
  prefectures: string
  town: string
}

export const isStoreCardDetailProps = (arg: any): arg is StoreCardDetailProps => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.name === "string" && typeof arg.phoneNumber === "string" &&
    typeof arg.postalCode === "string" && typeof arg.openingHours === "string" && typeof arg.closingHours === "string" &&
    typeof arg.prefectures === "string" && typeof arg.city === "string" && typeof arg.town === "string" && typeof arg.addressCode && typeof arg.addressOther === "string"
}

export const StoreCardDetail: FC<StoreCardDetailProps> = (props) => {
  return (
    <div>
      <h3>{props.name}</h3>
      <p>{props.phoneNumber}</p>
      <p>{`営業時間: ${props.openingHours} ~ ${props.closingHours}`}</p>
      <p>{`〒 ${props.postalCode}`}</p>
      <p>{`${props.prefectures}${props.city}${props.town}${props.addressCode}${props.addressOther}`}</p>
    </div>
  )
}

export interface BeauticianCardDetailProps {
  firstName: string
  lasttName: string
  phoneNumber: string
  lineId: string
  comment: string
  instagramId: string
}

export const isBeauticianCardDetail = (arg: any): arg is BeauticianCardDetailProps => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.firstName === "string" && typeof arg.lastName === "string" && typeof arg.phoneNumber === "string" &&
    typeof arg.lineId === "string" && typeof arg.comment === "string" && typeof arg.instagramId === "string"
}

export const BeauticianCardDetail: FC<BeauticianCardDetailProps> = (props) => {
  return (
    <div>
      <h3>{props.lasttName + props.firstName}</h3>
      <p>LINE ID: {props.lineId}</p>
      <p>{props.phoneNumber}</p>
      <p>コメント</p>
      <p>{props.comment}</p>
      <a href={props.instagramId} className="fab fa-instagram fa-2x"></a>
    </div>
  )
}
