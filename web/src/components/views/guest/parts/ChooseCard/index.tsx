import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import { SelectContext } from '../..';
import { Salon } from '../../../../../package/interface/Salon';
import ChooseCardComponent from './ChooseCard';

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon
}

const ChooseCard: FC<props> = (props) => {
  const history = useHistory()

  const selectCtx = useContext(SelectContext)

  const handleSelect = (id: string, type: "store" | "beautician") => {
    selectCtx(id, type)
    history.push("/guest")
  }

  return (
    <ChooseCardComponent image={props.image} type={props.type} content={props.content} handleSelect={handleSelect} />
  )
}

export default ChooseCard;