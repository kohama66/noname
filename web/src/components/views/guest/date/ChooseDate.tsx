import React, { FC, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import Schedule from '../../../../container/views/guest/date/Schedule';
import { setCheckedContext } from '../../../../container/views/guest/Guest';
import Title from '../parts/Title'
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
      <Title titleText={"日付を選ぶ"} image={"/img/choose1.png"}/>
      <div className="choose-date-contents">
        <Schedule />
      </div>
    </section>
  )
}

export default ChooseDate;