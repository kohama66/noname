import React, { FC } from 'react';
import { Salon } from '../../../../../package/interface/Salon';
import CardDetail from '../CardDetail';
import "./ChooseCard.scss"

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon
  handleSelect: (id: string, type: "beautician" | "store") => void
}

const ChooseCard: FC<props> = (props) => {

  return (
    <div className="choose-card" onClick={() => props.handleSelect(props.content.randId, props.type)} >
      <figure>
        <img src={props.image} alt="" />
      </figure>
      <CardDetail type={props.type} content={props.content}/>
    </div>
  )
}

export default ChooseCard;