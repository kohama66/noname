import React, { FC, useState, useEffect } from 'react';
import { Reservation } from '../../../../package/interface/Reservation';
import ScheduleComponent from "./Schedule"

interface props {
  reservations: Reservation[]
}

const Schedule: FC<props> = (props) => {
  const [twoWeeks, setTwoWeeks] = useState<Date[]>([])

  const getWeek = () => {
    const weeks = []
    for (let i = 0; i < 14; i++) {
      weeks.push(new Date())
      weeks[i].setDate(weeks[i].getDate() + i)
    }
    setTwoWeeks(weeks)
  }

  useEffect(() => {
    getWeek()
  }, [])

  return (
    <ScheduleComponent weeks={twoWeeks} reservations={props.reservations} />
  )
};

export default Schedule