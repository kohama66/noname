import React, { FC, useContext, useEffect, useState } from 'react';
import { GeterSelectCpntext } from '../..';
import ChoosePlateComponent from './ChoosePlate'

interface props {
  type: "beautician" | "store" | "date" | "menu"
}

const ChoosePlate: FC<props> = (props) => {
  const [image, setImage] = useState("")
  const [title, setTitle] = useState("")
  const [linkPath, setLinkPath] = useState("")
  const [checked, setChecked] = useState(false)
  const geterSelect = useContext(GeterSelectCpntext)

  useEffect(() => {
    const selectValue = geterSelect(props.type)
    switch (props.type) {
      case "store":
        setImage("img/1.png")
        setTitle("お店で選ぶ")
        setLinkPath("/store")
        if (selectValue != "") {
          setChecked(true)
        }
        break
    }
  }, [])
  
  return (
    <ChoosePlateComponent image={image} title={title} path={linkPath} checked={checked} />
  )
}

export default ChoosePlate;