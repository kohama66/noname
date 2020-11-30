import React, { FC, useContext, useEffect, useState } from 'react';
import { ReservedContext } from '../../../../../utils/context/ReservadContext ';
import ChoosePlateComponent from './ChoosePlate'

interface props {
  type: "beautician" | "store" | "date" | "menu"
}

const ChoosePlate: FC<props> = (props) => {
  const [image, setImage] = useState("")
  const [title, setTitle] = useState("")
  const [linkPath, setLinkPath] = useState("")
  const [checked, setChecked] = useState(false)
  const { store, beautician, getMenuIDs, reservationDate } = useContext(ReservedContext)


  const setPlateData = (image: string, title: string, path: string, value: string | string[] | undefined) => {
    setImage(image)
    setTitle(title)
    setLinkPath(path)
    if (value != null && value !== "" && value.length !== 0) {
      setChecked(true)
    }
  }

  useEffect(() => {
    switch (props.type) {
      case "store":
        setPlateData("img/shop.svg", "お店を選ぶ", "/store", store.randId)
        break
      case "beautician":
        setPlateData("img/beautician.svg", "美容師を選ぶ", "/beautician", beautician.randId)
        break
      case "menu":
        setPlateData("img/menu.svg", "メニューを選ぶ", "/menu", getMenuIDs())
        break
      case "date":
        setPlateData("img/karender.svg", "日付を選ぶ", "/date", reservationDate)
        break
    }
  }, [])

  return (
    <ChoosePlateComponent image={image} title={title} path={linkPath} checked={checked} />
  )
}

export default ChoosePlate;