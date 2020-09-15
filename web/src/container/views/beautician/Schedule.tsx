import React, { FC, useState, useEffect } from 'react';
import ScheduleComponent from "../../../components/views/beautician/Schedule"
import { getReservationBeautician } from '../../../package/api/index';
import { ReservationFindByBeautician } from '../../../package/interface/response/Reservation'

const Schedule: FC = () => {
  const [reservation, setReservation] = useState<ReservationFindByBeautician>()
  const handleReservation = async() => {
    const response = await getReservationBeautician()
    setReservation(response)
  }
  useEffect(() => {
    handleReservation()
  },[])
  return (
    <ScheduleComponent reservation={reservation} />
  )
};

export default Schedule