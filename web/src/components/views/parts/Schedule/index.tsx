import React, { FC, useState, useEffect } from 'react';
import ScheduleComponent from "./Schedule"

const Schedule: FC = () => {
  // const [reservation, setReservation] = useState<ReservationFindByBeautician>()
  const [twoWeeks, setTwoWeeks] = useState<Date[]>([])
  const handleReservation = async () => {
    // const response = await getReservationBeautician()
    // setReservation(response)
  }
  // const getWeek = () => {
  //   let weeks = []
  //   let today = new Date();
  //   for (let i = 0; i < 14; i++) {
  //     weeks.push(today.getDate());
  //     today.setDate(today.getDate() + 1)
  //   }
  //   setTowWeeks(weeks)
  // }
  const getWeek = () => {
    const weeks = []
    for (let i = 0; i < 14; i++) {
      weeks.push(new Date())
      weeks[i].setDate(weeks[i].getDate() + i)
    }
    setTwoWeeks(weeks)
  }

  useEffect(() => {
    handleReservation()
    getWeek()
  }, [])

  return (
    <ScheduleComponent weeks={twoWeeks} />
  )
};

export default Schedule