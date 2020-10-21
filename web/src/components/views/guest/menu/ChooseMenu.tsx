import React, { FC, useContext } from 'react';
import Title from '../parts/Title/Title'
import MenuBar from '../parts/MenuBar'
import './ChooseMenu.scss'
import { Menu } from '../../../../package/interface/Menu'

interface props {
  menus: Menu[]
  handleSetSelect: () => void
  handleSetIDs: (key: string, id: string, check: boolean) => void
}

const ChooseMenu: FC<props> = (props) => {
  return (
    <section id="choose-menu">
      <Title title="MENU" text="メニューを選ぶ" />
      <div className="menu-bar-wrapper">
        {props.menus.map((menu, i) => {
          return <MenuBar name={menu.name} image="/img/cuticon.png" randId={menu.randId} key={i} handleSetIDs={props.handleSetIDs} />
        })}
      </div>
      <button className="done-btn" onClick={() => props.handleSetSelect()}>
        <i className="fas fa-check fa-3x"></i>
      </button>
    </section>
  )
}

export default ChooseMenu;