import React, { FC } from 'react';

interface props {
  day: number;
}

const DaySquare: FC<props> = ({day}) => {
  return (
  <th>{day}</th>
  )
}

export default DaySquare;
