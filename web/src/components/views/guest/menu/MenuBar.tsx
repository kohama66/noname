import React, { FC, useState } from 'react';

interface ChooseMenuProps {
  menuName: string
  image: string
}

const MenuBar: FC<ChooseMenuProps> = (props) => {
  const [isCheck, toggleCheck] = useState(false)
  const [check, setCheck] = useState("0px")
  const handleCheck = () => {
    toggleCheck(!isCheck)
    if(isCheck){
      setCheck("0px")
    } else {
      setCheck("-50px")
    }
  }

  return (
    <button className="choose-menu-content" onClick={handleCheck} style={{transform: `translateX(${check})`}}>
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