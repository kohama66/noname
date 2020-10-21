import React, { FC, useState, useEffect, useContext } from 'react';
import {
  useRouteMatch,
  Link,
} from "react-router-dom";
import ChoosePlate from './parts/ChoosePlate';
import Title from './parts/Title/Title';

const Guest: FC = () => {
  const match = useRouteMatch();
  // const [checkedColor, setCheckedColor] = useState("100%")
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
    </article>
  )
}

export default Guest;