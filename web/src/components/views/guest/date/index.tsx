import React, { FC, useContext, useEffect, useState } from 'react';
import { findReservation } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseDateComponent from './ChooseDate'

const ChooseDate: FC = () => {
  const [reservations, setReservations] = useState<Reservation[]>([])
  const { getSelectID } = useContext(ReservedContext)
  const handleFindReservation = async () => {
    const beauticianID = getSelectID("beautician")
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