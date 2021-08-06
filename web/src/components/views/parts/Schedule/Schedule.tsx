import React, { FC, useState, useEffect } from 'react';
import "./Schedule.scss";
import { useMediaQuery } from 'react-responsive';
import { Reservation } from '../../../../package/interface/Reservation';
import DaySquare from './DaySquare/DaySquare';
import Square from './Square/Square';

interface props {
  reservations: Reservation[]
  handleModalClick?: (date: Date) => void
}

const Schedule: FC<props> = (props) => {
  const [displayWeek, setDisplayWeek] = useState<Date[]>([])
  const isDesktop = useMediaQuery({ query: "(min-width: 766px)" })
  const times = [9, 10, 11, 12, 13, 14, 15, 16, 17, 18]

  useEffect(() => {
    const createDisplayWeek = () => {
      const weeks = []
      var weekCount: number
      if (isDesktop) {
        weekCount = 14
      } else {
        weekCount = 7
      }
      for (let i = 0; i < weekCount; i++) {
        weeks.push(new Date())
        weeks[i].setDate(weeks[i].getDate() + i)
        weeks[i].setMinutes(0)
        weeks[i].setSeconds(0)
      }
      setDisplayWeek(weeks)
    }
    createDisplayWeek()
  }, [])

  

  return (
    <section id="schedule">
      <table>
        <tbody>
          <tr>
            <th></th>
            {displayWeek.map((day, i) => {
              return <DaySquare day={day.getDate()} date={day} key={i} />
            })}
          </tr>
          {times.map((time, i) => {
            return <tr key={i}>
              <td>{time}時</td>
              {displayWeek.map((day, i) => {
                return <Square key={i} day={day} time={time} reservations={props.reservations} handleModalClick={props.handleModalClick} />
              })}
            </tr>
          })}
        </tbody>
      </table>
    </section>
  )
};

export default Schedule