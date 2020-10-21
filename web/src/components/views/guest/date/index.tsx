import React, { FC, useContext, useState } from 'react';
import { GeterSelectIDContext } from '..';
import { findReservation } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import ChooseDateComponent from './ChooseDate'

const ChooseDate: FC = () => {
  const [reservation, setReservation] = useState<Reservation[]>([])
  const geterSelect = useContext(GeterSelectIDContext)

  const handleFindReservation = async () => {
    const beauticianID = geterSelect("beautician")
    try {
      if (typeof beauticianID === "string") {
        const response = await findReservation()
        setReservation(response.reservations)
      }
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <ChooseDateComponent reservation={reservation}/>
  )
}

export default ChooseDate;