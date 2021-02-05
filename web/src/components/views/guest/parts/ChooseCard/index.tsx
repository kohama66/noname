import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import { Salon } from '../../../../../package/interface/Salon';
import { User } from '../../../../../package/interface/User';
import { ReservedContext } from '../../../../../utils/context/ReservadContext ';
import ChooseCardComponent from './ChooseCard';

interface props {
  image?: string
  type: "store" | "beautician"
  content: Salon | User
}

const ChooseCard: FC<props> = (props) => {
  const history = useHistory()

  const { setSelectValue } = useContext(ReservedContext)

  const handleSetSelect = (id: string, type: "store" | "beautician", content: User | Salon) => {
    setSelectValue(content)
    history.push("/guest")
  }

  return (
    <ChooseCardComponent image={props.image} type={props.type} content={props.content} handleSetSelect={handleSetSelect} />
  )
}

export default ChooseCard;