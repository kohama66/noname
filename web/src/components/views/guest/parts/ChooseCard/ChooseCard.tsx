import React, { FC } from 'react';
import CardDetail from '../CardDetail';
import "./ChooseCard.scss"

interface props {
  image?: string
  type: "store" | "beautician"
}

const ChooseCard: FC<props> = (props) => {

  return (
    <div className="choose-card">
      <figure>
        <img src={props.image} alt="" />
      </figure>
      <CardDetail type={props.type} />
    </div>
  )
}

export default ChooseCard;