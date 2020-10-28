import React, { FC, useContext, useEffect, useState } from 'react';
import { GeterSelectIDContext } from '..';
import { findReservation } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import ChooseDateComponent from './ChooseDate'

const ChooseDate: FC = () => {
  const [reservations, setReservations] = useState<Reservation[]>([])
  const geterSelect = useContext(GeterSelectIDContext)

  const handleFindReservation = async () => {
    const beauticianID = geterSelect("beautician")
    try {
      if (typeof beauticianID === "string") {
        const response = await findReservation(beauticianID)
        setReservations(response.reservations)
      }
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    handleFindReservation()
  }, [])

  return (
    <ChooseDateComponent reservations={reservations} />
  )
}

export default ChooseDate;