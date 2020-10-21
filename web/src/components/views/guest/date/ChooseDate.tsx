import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
// import Schedule from '../../../../container/views/guest/date/Schedule';
import { setCheckedContext } from '../../guest';
import Schedule from '../../parts/Schedule';
import Title from '../parts/Title/Title'
import './ChooseDate.scss'

const ChooseDate: FC = () => {
  const history = useHistory()
  const setChecked = useContext(setCheckedContext)

  const handleMenu = () => {
    setChecked("date")
    history.push("/guest")
  }

  return (
    <section id="choose-date">
      <Title title="SCHEDULE" text="日付から選ぶ" />
      <div className="shedule-wrapper">
        <Schedule />
      </div>
    </section>
  )
}

export default ChooseDate;