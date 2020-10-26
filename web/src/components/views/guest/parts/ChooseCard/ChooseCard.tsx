import React, { FC } from 'react';
import { Beautician } from '../../../../../package/interface/Beautician';
import { Salon } from '../../../../../package/interface/Salon';
import CardDetail from '../CardDetail';
import "./ChooseCard.scss"

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon | Beautician
  handleSetSelect: (id: string, type: "beautician" | "store", content: Beautician | Salon) => void
}

const ChooseCard: FC<props> = (props) => {

  return (
    <div className="choose-card" onClick={() => props.handleSetSelect(props.content.randId, props.type, props.content)} >
      <figure>
        <img src={props.image} alt="" />
      </figure>
      <CardDetail type={props.type} content={props.content}/>
    </div>
  )
}

export default ChooseCard;