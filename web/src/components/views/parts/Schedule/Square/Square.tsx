import React, { FC } from 'react';

interface props {
  id?: number
  day: number
  time: number
}

const Square: FC<props> = (props) => {
  return (
  <td>{props.id}</td>
  )
}

export default Square;
