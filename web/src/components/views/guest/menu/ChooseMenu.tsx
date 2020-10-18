import React, { FC, useContext } from 'react';
import Title from '../parts/Title'
import MenuBar from '../parts/MenuBar/MenuBar'
import './ChooseMenu.scss'
import { setCheckedContext } from '../../guest'
import { useHistory } from 'react-router-dom';

const ChooseMenu: FC = () => {
  const history = useHistory()
  const setChecked = useContext(setCheckedContext)

  const handleMenu = () => {
    setChecked("menu")
    history.push("/guest")
  }

  return (
    <section id="choose-menu">
      <Title title="MENU" text="メニューを選ぶ" />
      <div className="menu-bar-wrapper">
      <MenuBar name="カット" image="/img/cuticon.png"/>
      </div>
      <button className="done-btn">
      <i className="fas fa-check fa-3x"></i>
      </button>
      {/* <Title titleText={"メニューを選ぶ"} image={"/img/choose1.png"} /> */}
      {/* <div className=".choose-menu-contents-wrapper"> */}
        {/* <div className="choose-menu-contents">
          <MenuBar menuName={"カット"} image={"/img/cuticon.png"} />
          <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
          <MenuBar menuName={"カラー"} image={"/img/color.png"} />
          <MenuBar menuName={"カット"} image={"/img/cuticon.png"} />
          <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
          <MenuBar menuName={"カラー"} image={"/img/color.png"} />
          <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
          <MenuBar menuName={"カラー"} image={"/img/color.png"} />
          <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
          <MenuBar menuName={"カラー"} image={"/img/color.png"} />
          <button className="fin-btn" onClick={handleMenu}>決定</button>
        </div> */}
      {/* </div> */}
    </section>
  )
}

export default ChooseMenu;