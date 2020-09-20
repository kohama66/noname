import React, { FC, useState } from 'react';

interface ChooseMenuProps {
  menuName: string
  image: string
}

const MenuBar: FC<ChooseMenuProps> = (props) => {
  const [isCheck, toggleCheck] = useState(false)
  const [color, setColor] = useState("#fff")
  const handleCheck = () => {
    toggleCheck(!isCheck)
    if(isCheck){
      setColor("#fff")
    } else {
      setColor("#ffc0a7")
    }
  }

  return (
    <button className="choose-menu-content" onClick={handleCheck} style={{backgroundColor: color}}>
      <figure>
        <img src={props.image} alt="" />
      </figure>
      <div>
        <h2>{props.menuName}</h2>
      </div>
    </button>
  )
}

export default MenuBar;