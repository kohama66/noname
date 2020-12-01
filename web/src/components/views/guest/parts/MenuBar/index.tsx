import React, { FC, useState } from 'react';
import { Menu } from '../../../../../package/interface/Menu';
import MenuBarComponent from './MenuBar'

interface props {
  image: string
  menu: Menu
  handleSetMenuValues: (key: string, value: Menu, check: boolean) => void
}

const MenuBar: FC<props> = (props) => {
  const [isCheck, toggleCheck] = useState(false)

  const handleCheck = () => {
    toggleCheck(!isCheck)
    props.handleSetMenuValues(props.menu.randId, props.menu, !isCheck)
  }

  return (
    <MenuBarComponent name={props.menu.name} handleCheck={handleCheck} check={isCheck} image={props.image} />
  )
}

export default MenuBar;