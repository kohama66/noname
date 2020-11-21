import React, { FC } from 'react';
import './Input.scss'

interface props {
  type: "text" | "email" | "password" | "submit"
  value: string
}

const Input: FC<props> = (props) => {
  return (
    <input type={props.type} value={props.value} />
  )
}

export default Input;