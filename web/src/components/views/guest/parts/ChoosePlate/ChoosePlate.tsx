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

  const handleChangeChecedView = () => {

  }

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
      </div>
    </Link>
  )
}

export default ChoosePlate;