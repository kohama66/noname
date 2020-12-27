import React, { FC } from 'react';
import './Square.scss';

interface props {
  clickFunction: () => void
  reserved: boolean
  pathName?: string
}

const Square: FC<props> = (props) => {
  return (
    <td onClick={props.clickFunction} className={props.reserved  && props.pathName !== "/beautician/mypage" ? "reserved-square" : ""}>
      {props.reserved ? "×" : ""}
    </td>
  )
}

export default Square;
