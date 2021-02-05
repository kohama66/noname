import React, { FC, useState } from 'react';
import "./Accordion.scss";

interface props {
  buttonText: string
}

const Accordion: FC<props> = (props) => {
  const [isOpen, toggle] = useState<boolean>(false)
  const handleClick = () => {
    toggle(!isOpen)
  }
  return (
    <div className="accordion-wrapper">
      <button onClick={handleClick}>{props.buttonText}</button>
      <div className={"accordion" + (isOpen ? " open-accordion" : "")} >
        {props.children}
      </div>
    </div>
  )
}

export default Accordion;