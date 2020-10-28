import React, { FC, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import GuestComponent from './Guest'

interface props {
  allSelectCheck: () => boolean
}

const GuestHome: FC<props> = (props) => {
  const history = useHistory()
  const [allCheck, setAllCheck] = useState<boolean>(false)
  const check = props.allSelectCheck()

  useEffect(() => {
    setAllCheck(check)
    if (allCheck) {
      history.push("guest/final_comfirmation")
    }
  }, [allCheck])
  return (
    <GuestComponent />
  )
}

export default GuestHome;