import React, { FC, useState, useEffect } from 'react';
import DaySquare from './DaySquare/DaySquare';
import Square from './Square';
import "./Schedule.scss"
import { Reservation } from '../../../../package/interface/Reservation';

interface props {
  reservations: Reservation[]
  weeks: Date[]
  pathName?: string
}

const Schedule: FC<props> = (props) => {
  const times = [9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
  return (
    <section id="schedule">
      <table>
        <tbody>
          <tr>
            <th></th>
            {props.weeks.map((day, i) => {
              return <DaySquare day={day.getDate()} key={i} />
            })}
          </tr>
          {times.map((time, i) => {
            return <tr key={i}>
              <td>{time}時</td>
              {props.weeks.map((day, i) => {
                return <Square key={i} day={day} time={time} reservations={props.reservations} pathName={props.pathName} />
              })}
            </tr>
          })}
        </tbody>
      </table>
    </section>
  )
};

export default Schedule;