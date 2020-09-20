import React, { FC, useState, useEffect } from 'react';
import {
  useRouteMatch,
  Link,
} from "react-router-dom";
import ChoosePlate from '../../../container/views/guest/ChoosePlate';
import Title from './parts/Title';

interface GuestProps {
  checkedBeautician: boolean
  checkedStore: boolean
  checkedMenu: boolean
  checkedDate: boolean
}

const Guest: FC<GuestProps> = (props) => {
  const match = useRouteMatch();
  const [checkedColor, setCheckedColor] = useState("100%")
  useEffect(() => {
    if (props.checkedBeautician) {
      // setCheckedColor("50%")
    }
  }, [])
  return (
    <article className="chooses">
      <Title titleText={"何から決めますか？"} image={"img/Questions-pana.png"} />
      <div className="chooses-contents">
        <ChoosePlate text={"お店で選ぶ"} image={"img/1.png"} path={"/store"}/>
        <ChoosePlate text={"美容師から選ぶ"} image={"img/3.png"} path={"/beautician"}/>
        <ChoosePlate text={"メニューから選ぶ"} image={"img/2.png"} path={"/menu"}/>
        <ChoosePlate text={"日付から選ぶ"} image={"img/4.png"} path={"/date"}/>
      </div>
    </article>
  )
}

export default Guest;