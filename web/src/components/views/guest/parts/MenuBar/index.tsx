import React, { FC, useState } from 'react';
import MenuBarComponent from './MenuBar'

interface props {
  name: string
  image: string
  randId: string
  handleSetIDs: (key: string, id: string, check: boolean) => void
}

const MenuBar: FC<props> = (props) => {
  const [isCheck, toggleCheck] = useState(false)
  const [checkStyle, setCheckStyle] = useState({ backgroundColor: "#fff", boxShadow: "", top: "5px", bottom: "", marginTop: "0" })

  const handleCheck = () => {
    toggleCheck(!isCheck)
    props.handleSetIDs(props.randId, props.randId, !isCheck)
    if (isCheck) {
      setCheckStyle({ backgroundColor: "#fff", boxShadow: "", top: "5px", bottom: "0", marginTop: "0" })
    } else {
      setCheckStyle({ backgroundColor: "#ffe260", boxShadow: "0 0 20px #ffff60", top: "", bottom: "5px", marginTop: "40px" })
    }
  }
  
  return (
    <MenuBarComponent name={props.name} style={checkStyle} handleCheck={handleCheck} />
  )
}

export default MenuBar;