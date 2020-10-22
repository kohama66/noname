import React, { FC, useContext } from 'react';
import { Reservation } from '../../../../package/interface/Reservation';
import Schedule from '../../parts/Schedule';
import Title from '../parts/Title/Title'
import './ChooseDate.scss'

interface props {
  reservations: Reservation[]
}

const ChooseDate: FC<props> = (props) => {
  return (
    <section id="choose-date">
      <Title title="SCHEDULE" text="日付から選ぶ" />
      <div className="shedule-wrapper">
        <Schedule reservations={props.reservations}/>
      </div>
    </section>
  )
}

export default ChooseDate;