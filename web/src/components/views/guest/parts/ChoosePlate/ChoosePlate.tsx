import React, { FC, useEffect, useState } from 'react';
import { Link, useRouteMatch } from 'react-router-dom';
import "./ChoosePlate.scss"

export interface ChoosePlateProps {
  title: string
  image: string
  path: string
  checked: boolean
}

const ChoosePlate: FC<ChoosePlateProps> = (props) => {
  const match = useRouteMatch()
  const [checkedColor, setCheckedColor] = useState("100")

  useEffect(() => {
    if (props.checked) {
      setCheckedColor("50")
    }
  }, [props.checked])

  return (
    <Link to={match.path + props.path} style={{ filter: `brightness(${checkedColor}%)` }} className="choose-plate">
      <div>
        <img src={props.image} alt="" />
        <h2>{props.title}</h2>
      </div>
    </Link>
  )
}

export default ChoosePlate;