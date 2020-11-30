import React, { FC, useContext, useEffect, useState } from 'react';
import { findReservation } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChooseDateComponent from './ChooseDate'

const ChooseDate: FC = () => {
  const [reservations, setReservations] = useState<Reservation[]>([])
  const { beautician } = useContext(ReservedContext)

  const handleFindReservation = async () => {
    try {
      const response = await findReservation(beautician.randId)
      setReservations(response.reservations)
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