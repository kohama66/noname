import React, { FC, useState } from 'react';
import "./MenuBar.scss"

interface props {
  name: string
  check: boolean
  image: string
  handleCheck: () => void
}

const MenuBar: FC<props> = (props) => {
  return (
    <div className="menu-bar">
      <figure>
        <img src={props.image} alt="メニューイメージ" />
      </figure>
      <div>
        <h2>{props.name}</h2>
      </div>
      <button className={"toggleBtn" + (props.check ? " toggleBtn_check": "")}>
        <div onClick={props.handleCheck}></div>
      </button>
    </div>
  )
}

export default MenuBar;