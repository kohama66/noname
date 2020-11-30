import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { Reservation } from '../../../../../package/interface/Reservation';
import { ReservedContext } from '../../../../../utils/context/ReservadContext ';
import SquareComponent from './Square'

interface props {
  day: Date
  time: number
  reservations: Reservation[]
}

const Square: FC<props> = (props) => {
  const history = useHistory()
  const [time, setTime] = useState<string>("")
  const [reserved, setReserved] = useState<boolean>(false)
  // const setSelectValue = useContext(SetSelectValueContext)
  const { setSelectValue } = useContext(ReservedContext)

  const parseTime = (time: number) => {
    setTime(("0" + time).slice(-2) + ":00:00")
  }
  const handleSetSelectDate = () => {
    setTime(("0" + props.time).slice(-2) + ":00:00")
    setSelectValue(props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + " " + time)
    history.push("/guest")
  }
  const verifyReserved = () => {
    props.reservations.map((reservation) => {
      if (reservation.date === props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + "T" + (("0" + props.time).slice(-2) + ":00:00Z")) {
        setReserved(true)
      }
    })
  }

  useEffect(() => {
    parseTime(props.time)
    verifyReserved()
  })

  return (
    <SquareComponent clickFunction={handleSetSelectDate} reserved={reserved} />
  )
}

export default Square;