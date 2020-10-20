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
  name: string
  phoneNumber: string
  lineId: string
  comemnt: string
  instaLink: string
}

export const isBeauticianCardDetail = (arg: any): arg is BeauticianCardDetailProps => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.name === "string" && arg.phoneNumber === "string" &&
    arg.lineId === "string" && arg.comment === "string" && arg.instalink === "strng"
}

export const BeauticianCardDetail: FC<BeauticianCardDetailProps> = (props) => {
  return (
    <div>
      <h3>{props.name}</h3>
      <p>LINE ID: {props.lineId}</p>
      <p>{props.phoneNumber}</p>
      <p>コメント</p>
      <p>{props.comemnt}</p>
      <a href={props.instaLink} className="fab fa-instagram fa-2x"></a>
    </div>
  )
}
