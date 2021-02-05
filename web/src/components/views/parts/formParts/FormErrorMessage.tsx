import React, { FC } from 'react';
import './FormErrorMessage.scss';

interface props {
  error: string | undefined
}

const FormErrorMessage: FC<props> = (props) => {
  return (
    <span className="error-message">
      <p>{props.error}</p>
    </span>
  )
}

export default FormErrorMessage;