import React, { FC } from 'react';
import { Link, useRouteMatch } from 'react-router-dom';
import ChoosePlateComponent, { ChoosePlateProps } from '../../../components/views/guest/parts/ChoosePlate/ChoosePlate'

const ChoosePlate: FC<ChoosePlateProps> = (props) => {
  const match = useRouteMatch()
  return (
    <></>
  )
}

export default ChoosePlate;