import React, { FC, useEffect, useState } from 'react';
import { Link, useRouteMatch } from 'react-router-dom';

export interface ChoosePlateProps {
  text: string
  image: string
  path: string
  checked?: boolean
}

const ChoosePlate: FC<ChoosePlateProps> = ({
  text,image,path,
  checked = false
}) => {
  const match = useRouteMatch()
  const [checkedColor, setCheckedColor] = useState("100")
  console.log(checked)
  useEffect(() => {
    if(checked){
      setCheckedColor("50")
    }
  },[])
  return (
    <Link to={match.path + path} style={{filter: `brightness(${checkedColor}%)`}}>
      <div>
        <img src={image} alt="" />
        <h2>{text}</h2>
      </div>
    </Link>
  )
}

export default ChoosePlate;