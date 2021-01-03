import React, { FC } from 'react';
import Input from './Input';
import './FormParts.scss';
import Select from './Select';

interface props {
  type: "text" | "email" | "password" | "submit" | "select"
  value: string
  setState?: React.Dispatch<React.SetStateAction<string>>
  required?: true
  maxLength?: number
  minLength?: number
  placeHolder?: string
  disabled?: boolean
  label?: string
}

const FormParts: FC<props> = (props) => {
  return (
    <span className="form-parts">
      <label>{props.label}</label>
      {(() => {
        if (props.type === "select") {
          return <Select type="menu" setState={props.setState} value={props.value} />
        } else return <Input type={props.type} value={props.value} required={true} setState={props.setState} disabled={props.disabled}
          placeHolder={props.placeHolder} maxLength={props.maxLength} minLength={props.minLength} />
      })()}
    </span>
  )
}

export default FormParts;