import React, { FC } from 'react';

interface DaySqusreProps {
  day: number;
}

const DaySquare: FC<DaySqusreProps> = ({day}) => {
  return (
  <th>{day}</th>
  )
}

export default DaySquare;
