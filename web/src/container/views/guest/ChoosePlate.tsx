import React, { FC } from 'react';
import { Link, useRouteMatch } from 'react-router-dom';
import ChoosePlateComponent, { ChoosePlateProps } from '../../../components/views/guest/parts/ChoosePlate'

const ChoosePlate: FC<ChoosePlateProps> = (props) => {
  const match = useRouteMatch()
  return (
    <ChoosePlateComponent text={props.text} image={props.image} path={props.path} checked={props.checked}/>
  )
}

export default ChoosePlate;