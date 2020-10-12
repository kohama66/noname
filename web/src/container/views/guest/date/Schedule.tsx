import React, { FC, useState, useEffect } from 'react';
import ScheduleComponent from "../../../../components/views/beautician/Schedule"
import { getReservationBeautician } from '../../../../package/api/index';
import { ReservationFindByBeautician } from '../../../../package/interface/response/Reservation'

const Schedule: FC = () => {
  const [twoWeeks, setTowWeeks ] = useState<number[]>([])
  const getToday =  () => {
    let weeks = []
    let today = new Date();
    for (let i = 0; i < 14; i++){
      weeks.push(today.getDate());
      today.setDate(today.getDate() + 1)
    }
    setTowWeeks(weeks)
  }
  useEffect(() => {
    getToday()
  }, [])
  return (
    <ScheduleComponent weeks={twoWeeks}/>
  )
};

export default Schedule