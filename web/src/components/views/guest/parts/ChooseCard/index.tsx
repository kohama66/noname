import React, { FC, useContext } from 'react';
import { SelectContext } from '../..';
import { Salon } from '../../../../../package/interface/Salon';
import ChooseCardComponent from './ChooseCard';

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon
}

const ChooseCard: FC<props> = (props) => {
  const handleSelect = useContext(SelectContext)

  return (
    <ChooseCardComponent image={props.image} type={props.type} content={props.content} handleSelect={handleSelect}/>
  )
}

export default ChooseCard;