import React, { FC } from 'react';

interface props {
  day: number;
  date?: Date
}

const DaySquare: FC<props> = (props) => {
  const handleClick = () => {
    if (props.date?.getDate() !== (new Date).getDate()){
      
    } else {
      console.log("今日です")
    }
  }
  return (
    <th onClick={handleClick} >{props.day}</th>
  )
}

export default DaySquare;
