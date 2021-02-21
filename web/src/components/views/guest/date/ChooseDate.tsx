import React, { FC, useContext, useEffect, useState } from 'react';
import { findReservation } from '../../../../package/api';
import { Reservation } from '../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import Schedule from '../../parts/Schedule/Schedule';
import Title from '../parts/Title/Title';
import './ChooseDate.scss'

const ChooseDate: FC = () => {
  const [reservations, setReservations] = useState<Reservation[]>([])
  const { beautician } = useContext(ReservedContext)

  useEffect(() => {
    const handleFindReservation = async () => {
      try {
        const response = await findReservation(beautician.randId)
        setReservations(response.reservations)
      } catch (error) {
        console.log(error)
      }
    }
    handleFindReservation()
  }, [])

  return (
    <section id="choose-date">
      <Title title="SCHEDULE" text="日付から選ぶ" />
      <div className="shedule-wrapper">
        <Schedule reservations={reservations} />
      </div>
    </section>
  )
}

export default ChooseDate;