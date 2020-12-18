import React, { FC, useState } from 'react';
import './Input.scss'

interface props {
  type: "text" | "email" | "password" | "submit" | "hidden"
  value: string
  setState?: React.Dispatch<React.SetStateAction<string>>
  required?: true
  maxLength?: number
  minLength?: number
  placeHolder?: string
  disabled?: boolean
}

const Input: FC<props> = (props) => {
  const handleChange = (e: React.FormEvent<HTMLInputElement>) => {
    if (props.setState) {
      props.setState(e.currentTarget.value)
    }
  }

  return (
    <input type={props.type} value={props.value} onChange={handleChange} required={props.required}
      maxLength={props.maxLength} minLength={props.minLength} placeholder={props.placeHolder}
      disabled={props.disabled}
    />
  )
}

export default Input;