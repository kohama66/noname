import React, { FC, useState, useEffect } from 'react';
import DaySquare from './DaySquare';
import Square from './Square';
import { ReservationFindByBeautician } from '../../../package/interface/response/Reservation'

type props = {
  reservation: ReservationFindByBeautician | undefined
}

const Schedule: FC<props> = ({ reservation }) => {
  const days = [1, 2, 3, 4, 5, 6, 7, 20, 21, 22, 23, 24, 25]
  const times = [7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18]
  return (
    <section id="schedule">
      <table>
        <tbody>
          <tr>
            <th></th>
            {days.map((day, i) => {
              return <DaySquare day={day} key={i} />
            })}
          </tr>
          {times.map((time, i) => {
            return <tr key={i}>
              <td>{time}時</td>
              {days.map((day, i) => {
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
      </table>
    </section>
  )
};

export default Schedule;
