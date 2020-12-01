import React, { FC, useContext } from 'react';
import Title from '../parts/Title/Title'
import MenuBar from '../parts/MenuBar'
import './ChooseMenu.scss'
import { Menu } from '../../../../package/interface/Menu'

interface props {
  menus: Menu[]
  handleSetSelect: () => void
  handleSetMenuValues: (key: string, value: Menu, check: boolean) => void
}

const ChooseMenu: FC<props> = (props) => {
  return (
    <section id="choose-menu">
      <Title title="MENU" text="メニューを選ぶ" />
      <div className="menu-bar-wrapper">
        {props.menus.map((menu, i) => {
          return <MenuBar menu={menu} image="/img/cut.svg" key={i} handleSetMenuValues={props.handleSetMenuValues} />
        })}
      </div>
      <button className="done-btn" onClick={() => props.handleSetSelect()}>
        <i className="fas fa-check"></i>
      </button>
    </section>
  )
}

export default ChooseMenu;