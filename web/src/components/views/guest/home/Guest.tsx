import React, { FC } from 'react';
import ChoosePlate from '../parts/ChoosePlate';
import Title from '../parts/Title/Title';

interface props {
}

const Guest: FC<props> = (props) => {
  return (
    <article id="chooses">
      <Title text="何から選びますか？" title="SELECT" />
      <div className="chooses-contents-wrapper">
        <div className="chooses-contents">
          <ChoosePlate type="store" />
          <ChoosePlate type="beautician" />
          <ChoosePlate type="menu" />
          <ChoosePlate type="date" />
        </div>
      </div>
    </article>
  )
}

export default Guest;