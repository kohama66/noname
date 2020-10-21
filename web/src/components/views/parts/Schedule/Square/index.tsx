import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { SetSelectContext } from '../../../guest';
import SquareComponent from './Square'

interface props {
  // id?: number
  day: Date
  time: number
}

const Square: FC<props> = (props) => {
  const history = useHistory()
  const [time, setTime] = useState<string>("")
  const setSelectDate = useContext(SetSelectContext)

  const parseTime = (time: number) => {
    setTime(("0" + time).slice(-2) + ":00:00")
  }
  const handleSetSelectDate = () => {
    setSelectDate(props.day.getFullYear() + "-" + ("0" + (props.day.getMonth() + 1)).slice(-2) + "-" + ("0" + props.day.getDate()).slice(-2) + " " + time, "date")
    history.push("/guest")
  }

  useEffect(() => {
    parseTime(props.time)
  }, [])

  return (
    <SquareComponent clickFunction={handleSetSelectDate} />
  )
}

export default Square;