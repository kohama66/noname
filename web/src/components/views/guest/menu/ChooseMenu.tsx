import React, { FC, useContext } from 'react';
import Title from '../parts/Title'
import MenuBar from './MenuBar'
import './ChooseMenu.scss'
import { setCheckedContext } from '../../../../container/views/guest/Guest'
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
      <Title titleText={"メニューを選ぶ"} image={"/img/choose1.png"} />
      <div className="choose-menu-contents">
        <MenuBar menuName={"カット"} image={"/img/cuticon.png"} />
        <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
        <MenuBar menuName={"カラー"} image={"/img/color.png"} />
        <MenuBar menuName={"カット"} image={"/img/cuticon.png"} />
        <MenuBar menuName={"パーマ"} image={"/img/curl.png"} />
        <MenuBar menuName={"カラー"} image={"/img/color.png"} />
        <button className="fin-btn" onClick={handleMenu}>決定</button>
      </div>
    </section>
  )
}

export default ChooseMenu;