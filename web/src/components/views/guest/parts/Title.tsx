import React, { FC } from 'react';
import { Path } from 'typescript';
import './Title.scss'

interface TitleProps {
  titleText: string
  image: string
}

const Title: FC<TitleProps> = (props) => {
  return (
    <span className="choose-title">
      <img src={props.image} alt="" />
      <h1>{props.titleText}</h1>
    </span>
  )
}

export default Title;