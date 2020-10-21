import React, { FC } from 'react';
import SquareComponent from './Square'

interface props {
  id?: number
  day: number
  time: number
}

const Square: FC<props> = (props) => {
  
  return (
    <SquareComponent id={props.id} day={props.day} time={props.time}/>
  )
}

export default Square;