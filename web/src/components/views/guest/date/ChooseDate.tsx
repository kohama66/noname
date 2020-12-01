import React, { FC, useContext } from 'react';
import MediaQuery from 'react-responsive';
import { Reservation } from '../../../../package/interface/Reservation';
import Schedule, { Schedule_smartphoneResponsive } from '../../parts/Schedule';
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
        <MediaQuery query="(min-width: 767px)">
          <Schedule reservations={props.reservations} />
        </MediaQuery>
        <MediaQuery query="(max-width: 767px)">
          <Schedule_smartphoneResponsive reservations={props.reservations} />
        </MediaQuery>
      </div>
    </section>
  )
}

export default ChooseDate;