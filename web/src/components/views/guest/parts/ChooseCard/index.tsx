import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import { SetSelectValueContext } from '../..';
import { Beautician } from '../../../../../package/interface/Beautician';
import { Salon } from '../../../../../package/interface/Salon';
import ChooseCardComponent from './ChooseCard';

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon | Beautician
}


const ChooseCard: FC<props> = (props) => {
  const history = useHistory()

  const setSelectValue = useContext(SetSelectValueContext)

  const handleSetSelect = (id: string, type: "store" | "beautician", content: Beautician | Salon) => {
    setSelectValue(content)
    history.push("/guest")
  }

  return (
    <ChooseCardComponent image={props.image} type={props.type} content={props.content} handleSetSelect={handleSetSelect} />
  )
}

export default ChooseCard;