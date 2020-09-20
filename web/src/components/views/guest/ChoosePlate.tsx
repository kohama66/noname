import React, { FC } from 'react';
import { Link, useRouteMatch } from 'react-router-dom';

export interface ChoosePlateProps {
  text: string
  image: string
  path: string
}

const ChoosePlate: FC<ChoosePlateProps> = (props) => {
  const match = useRouteMatch()
  return (
    <Link to={match.path + props.path}>
      <div>
        <img src={props.image} alt="" />
        <h2>{props.text}</h2>
      </div>
    </Link>
  )
}

export default ChoosePlate;