import React, { FC } from 'react';
import Popup from '../parts/Popup';
import ChoosePlate from './parts/ChoosePlate';
import Title from './parts/Title/Title';

interface props {
  allSelectCheck: boolean
}

const Guest: FC<props> = (props) => {
  return (
    <article className="chooses">
      <Title text="何から選びますか？" title="SELECT" />
      <div className="chooses-contents-wrapper">
        <div className="chooses-contents">
          <ChoosePlate type="store" />
          <ChoosePlate type="beautician" />
          <ChoosePlate type="menu" />
          <ChoosePlate type="date" />
        </div>
      </div>
      <Popup allSelectCheck={props.allSelectCheck}/>
    </article>
  )
}

export default Guest;