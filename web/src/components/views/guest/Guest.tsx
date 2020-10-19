import React, { FC, useState, useEffect, useContext } from 'react';
import {
  useRouteMatch,
  Link,
} from "react-router-dom";
import { GeterSelectCpntext } from '.';
import ChoosePlate from './parts/ChoosePlate';
import Title from './parts/Title';

interface GuestProps {
  checkedBeautician: boolean
  checkedStore: boolean
  checkedMenu: boolean
  checkedDate: boolean
  
}

const Guest: FC<GuestProps> = (props) => {
  const match = useRouteMatch();
  // const [checkedColor, setCheckedColor] = useState("100%")
  return (
    <article className="chooses">
      <Title text="何から選びますか？" title="SELECT" />
      <div className="chooses-contents-wrapper">
        <div className="chooses-contents">
          {/* <ChoosePlate text={"お店で選ぶ"} image={"img/1.png"} path={"/store"} checked={props.checkedStore} />
          <ChoosePlate text={"美容師から選ぶ"} image={"img/1.png"} path={"/beautician"} checked={props.checkedBeautician} />
          <ChoosePlate text={"メニューから選ぶ"} image={"img/1.png"} path={"/menu"} checked={props.checkedMenu} />
          <ChoosePlate text={"日付から選ぶ"} image={"img/4.png"} path={"/date"} checked={props.checkedDate} /> */}
          <ChoosePlate type="store" />
        </div>
      </div>
    </article>
  )
}

export default Guest;