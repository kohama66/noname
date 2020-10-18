import React, { FC } from 'react';

interface StoreCardDetailProps {
  name: string
  phoneNumber: string
  postalCode: string
  businessHours: string
  address: string
}

export const StoreCardDetail: FC<StoreCardDetailProps> = (props) => {
  return (
    <div>
      <h3>{props.name}</h3>
      <p>{props.phoneNumber}</p>
      <p>{props.businessHours}</p>
      <p>{props.postalCode}</p>
      <p>{props.address}</p>
    </div>
  )
}

interface BeauticianCardDetailProps {
  name: string
  phoneNumber: string
  lineId: string
  comemnt: string
  instaLink: string
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
