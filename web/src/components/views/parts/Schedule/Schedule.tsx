import React, { FC, useState, useEffect } from 'react';
import DaySquare from './DaySquare/DaySquare';
import Square from './Square';
import "./Schedule.scss"
import { Reservation } from '../../../../package/interface/Reservation';

interface props {
  reservation?: Reservation[]
  weeks: Date[]
}

const Schedule: FC<props> = (props) => {
  const times = [9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
  return (
    <section id="schedule">
      <table>
        <tbody>
          <tr>
            <th></th>
            {/* {props.weeks?.map((day, i) => {
              return <DaySquare day={day} key={i} />
            })} */}
            {props.weeks.map((day, i) => {
              return <DaySquare day={day.getDate()} key={i} />
            })}
          </tr>
          {times.map((time, i) => {
            return <tr key={i}>
              <td>{time}時</td>
              {props.weeks.map((day, i) => {
                let d
                // if (props.reservation) {
                //   {
                //     props.reservation.map((res) => {
                //       const resday = new Date(res.date)
                //       if (day == resday.getDate() && time + ":00:00" == res.date) {
                //         d = resday.getDate()
                //       }
                //     })
                //   }
                // }
                return <Square id={d} key={i} day={day.getDate()} time={time}/>
              })}
            </tr>
          })}
        </tbody>
      </table>
    </section>
  )
};

export default Schedule;