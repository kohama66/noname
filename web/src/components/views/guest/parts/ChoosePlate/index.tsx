import React, { FC, useContext, useEffect, useState } from 'react';
import { GeterSelectIDContext } from '../..';
import ChoosePlateComponent from './ChoosePlate'

interface props {
  type: "beautician" | "store" | "date" | "menu"
}

const ChoosePlate: FC<props> = (props) => {
  const [image, setImage] = useState("")
  const [title, setTitle] = useState("")
  const [linkPath, setLinkPath] = useState("")
  const [checked, setChecked] = useState(false)
  const geterSelect = useContext(GeterSelectIDContext)

  const setPlateData = (image: string, title: string, path: string, value: string | string[] | undefined) => {
    setImage(image)
    setTitle(title)
    setLinkPath(path)
    if (value !== undefined && value !== "" && value !== null && value.length !== 0) {
      setChecked(true)
    }
  }

  useEffect(() => {
    const selectValue = geterSelect(props.type)
    switch (props.type) {
      case "store":
        setPlateData("img/shop.svg", "お店を選ぶ", "/store", selectValue)
        break
      case "beautician":
        setPlateData("img/beautician.svg", "美容師を選ぶ", "/beautician", selectValue)
        break
      case "menu":
        setPlateData("img/menu.svg", "メニューを選ぶ", "/menu", selectValue)
        break
      case "date":
        setPlateData("img/karender.svg", "日付を選ぶ", "/date", selectValue)
        break
    }
  }, [])

  return (
    <ChoosePlateComponent image={image} title={title} path={linkPath} checked={checked} />
  )
}

export default ChoosePlate;