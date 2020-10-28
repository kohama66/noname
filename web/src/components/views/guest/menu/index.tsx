import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { GeterSelectIDContext, SetSelectValueContext } from '..';
import { findMenus } from '../../../../package/api';
import { Menu } from '../../../../package/interface/Menu';
import ChooseMenuComponent from './ChooseMenu'

const ChooseMenu: FC = () => {
  const history = useHistory()
  const [menus, setMenus] = useState<Menu[]>([])
  const [menuValues, setMenuValues] = useState<{ [key: string]: Menu }>({})
  const geterSelectID = useContext(GeterSelectIDContext)
  const setSelectValue = useContext(SetSelectValueContext)

  const handleFindMenus = async () => {
    const beauticianID = geterSelectID("beautician")
    try {
      if (typeof beauticianID === "string" || beauticianID == null) {
        const response = await findMenus(beauticianID)
        setMenus(response.menus)
      }
    } catch (err) {
      console.log(err)
    }
  }
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
    handleFindMenus()
  }, [])

  return (
    <ChooseMenuComponent menus={menus} handleSetSelect={handleSetSelect} handleSetMenuValues={handleSetMenuValues} />
  )
}

export default ChooseMenu;