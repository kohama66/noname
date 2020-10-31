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
  const [checked, setChecked] = useState<boolean>(false)

  useEffect(() => {
    if (props.checked) {
      setChecked(true)
    }
  })

  return (
    <Link to={match.path + props.path} className={"choose-plate" + (checked ? " choose-plate-checked": "")}>
      <div>
        <img src={props.image} alt="" />
        <h2>{props.title}</h2>
        <p className={props.checked ? "checked-choose-plate" : ""}>選択済み</p>
      </div>
    </Link>
  )
}

export default ChoosePlate;