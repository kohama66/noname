import React, { FC, useContext, useEffect, useState } from 'react';
import Title from '../parts/Title/Title'
import MenuBar from '../parts/MenuBar'
import './ChooseMenu.scss'
import { Menu } from '../../../../package/interface/Menu'
import { useHistory } from 'react-router-dom';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import { findMenus } from '../../../../package/api';

const ChooseMenu: FC = () => {
  const history = useHistory()
  const [menus, setMenus] = useState<Menu[]>([])
  const [menuValues] = useState<{ [key: string]: Menu }>({})
  const { beautician, setSelectValue } = useContext(ReservedContext)

  const handleSetMenuValues = (key: string, value: Menu, check: boolean): void => {
    if (check) {
      menuValues[key] = value
    } else {
      delete menuValues[key]
    }
  }
  const handleSetSelect = () => {
    const selectMenus: Menu[] = []
    Object.entries(menuValues).forEach(([key, value]) => {
      selectMenus.push(value)
    })
    setSelectValue(selectMenus)
    history.push("/guest")
  }

  useEffect(() => {
    const handleFindMenus = async () => {
      try {
        const response = await findMenus(beautician.randId)
        setMenus(response.menus)
      } catch (err) {
        console.log(err)
      }
    }
    handleFindMenus()
  }, [])

  return (
    <section id="choose-menu">
      <Title title="MENU" text="メニューを選ぶ" />
      <div className="menu-bar-wrapper">
        {menus.map((menu, i) => {
          return <MenuBar menu={menu} image="/img/cut.svg" key={i} handleSetMenuValues={handleSetMenuValues} />
        })}
      </div>
      <button className="done-btn" onClick={handleSetSelect}>
        <i className="fas fa-check"></i>
      </button>
    </section>
  )
}

export default ChooseMenu;