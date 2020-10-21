import React, { FC, useState, useEffect } from 'react';
import DaySquare from './DaySquare';
import Square from './Square';
// import { ReservationFindByBeautician } from '../../../package/interface/response/Reservation'

interface props {
  // reservation?: ReservationFindByBeautician | undefined
  weeks?: number[]
}

const Schedule: FC<props> = (props) => {
  const times = [9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
  return (
    <section id="schedule">
      {/* <h2>SEPTEMBER</h2>
      <table>
        <tbody>
          <tr>
            <th></th>
            {weeks?.map((day, i) => {
              return <DaySquare day={day} key={i} />
            })}
          </tr>
          {times.map((time, i) => {
            return <tr key={i}>
              <td>{time}時</td>
              {weeks?.map((day, i) => {
                let d
                if (reservation) {
                  {
                    reservation.Reservations.map((res) => {
                      const resday = new Date(res.date)
                      if (day == resday.getDate() && time + ":00:00" == res.time) {
                        d = resday.getDate()
                      }
                    })
                  }
                }
                return <Square id={d} key={i} />
              })}
            </tr>
          })}
        </tbody>
      </table> */}
    </section>
  )
};

export default Schedule;
