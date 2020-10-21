import React, { FC } from 'react';
import { Path } from 'typescript';
import './Title.scss'

interface TitleProps {
  title: string
  text: string
}

const Title: FC<TitleProps> = (props) => {
  return (
    <div className="heading">
      <h2>{props.title}</h2>
      <p>{props.text}</p>
    </div>
  )
}

export default Title;