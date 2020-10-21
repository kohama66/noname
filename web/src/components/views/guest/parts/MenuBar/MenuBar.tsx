import React, { FC, useState } from 'react';
import "./MenuBar.scss"

interface props {
  name: string
  handleCheck: () => void
  style: {
    backgroundColor: string;
    boxShadow: string;
    top: string;
    bottom: string;
    marginTop: string;
  }

}

const MenuBar: FC<props> = (props) => {
  return (
    <div className="menu-bar">
      <figure>
        {/* <img src={props.image} alt="メニューイメージ" /> */}
      </figure>
      <div>
        <h2>{props.name}</h2>
      </div>
      <button className="toggleBtn">
        <div style={props.style} onClick={props.handleCheck}></div>
      </button>
    </div>
  )
}

export default MenuBar;