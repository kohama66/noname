import React, { FC } from 'react';
import PopupComponent from "./Popup"

interface props {
  allSelectCheck: boolean
}

const Popup: FC<props> = (props) => {
  if (props.allSelectCheck) {
    return <PopupComponent display={props.allSelectCheck} />
  } else {
    return <></>
  }
}

export default Popup;