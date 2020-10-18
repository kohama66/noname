import React, { FC, useState, useEffect } from 'react';
import { ReservationFindByBeautician } from '../../../../package/interface/response/Reservation';
import ScheduleComponent from "./Schedule"

const Schedule: FC = () => {
  // const [reservation, setReservation] = useState<ReservationFindByBeautician>()
  const [twoWeeks, setTowWeeks] = useState<number[]>([])
  const handleReservation = async () => {
    // const response = await getReservationBeautician()
    // setReservation(response)
  }
  const getToday = () => {
    let weeks = []
    let today = new Date();
    for (let i = 0; i < 14; i++) {
      weeks.push(today.getDate());
      today.setDate(today.getDate() + 1)
    }
    setTowWeeks(weeks)
  }
  useEffect(() => {
    handleReservation()
    getToday()
  }, [])
  return (
    <ScheduleComponent weeks={twoWeeks} />
  )
};

export default Schedule