import { type } from 'os';
import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { GeterSelectIDContext, SetSelectContext } from '..';
import { findMenus } from '../../../../package/api';
import { Menu } from '../../../../package/interface/Menu';
import { Beautician } from '../../../endpoint';
import ChooseMenuComponent from './ChooseMenu'

const ChooseMenu: FC = () => {
  const history = useHistory()
  const [menus, setMenus] = useState<Menu[]>([])
  const [MenuIDs, setMenuIDs] = useState<{ [key: string]: string }>({})
  const geterSelectID = useContext(GeterSelectIDContext)
  const setSelectMenu = useContext(SetSelectContext)

  const handleFindMenus = async () => {
    const beauticianID = geterSelectID("beautician")
    try {
      if (typeof beauticianID === "string") {
        const response = await findMenus(beauticianID)
        setMenus(response.menus)
      }
    } catch (err) {
      console.log(err)
    }
  }
  const handleSetMenuIDs = (key: string, id: string, check: boolean): void => {
    if (check) {
      MenuIDs[key] = id
    } else {
      delete MenuIDs[key]
    }
  }
  const handleSetSelect = () => {
    const IDs: string[] = []
    Object.entries(MenuIDs).forEach(([key, value]) => {
      IDs.push(value)
    });
    setSelectMenu(IDs, "menu")
    history.push("/guest")
  }

  useEffect(() => {
    handleFindMenus()
  }, [])

  return (
    <ChooseMenuComponent menus={menus} handleSetSelect={handleSetSelect} handleSetIDs={handleSetMenuIDs} />
  )
}

export default ChooseMenu;