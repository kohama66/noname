import React, { FC, useState } from 'react';
import "./MenuBar.scss"

interface ChooseMenuProps {
  name: string
  image: string
}

const MenuBar: FC<ChooseMenuProps> = (props) => {
  const [isCheck, toggleCheck] = useState(false)
  const [checkStyle, setCheckStyle] = useState({backgroundColor: "#fff", boxShadow: "", top: "5px", bottom: "", marginTop: "0"})
  const handleCheck = () => {
    toggleCheck(!isCheck)
    if(isCheck){
      setCheckStyle({backgroundColor: "#fff", boxShadow: "", top: "5px", bottom: "0", marginTop: "0"})
    } else {
      setCheckStyle({backgroundColor: "#ffe260", boxShadow: "0 0 20px #ffe260", top: "", bottom: "5px", marginTop: "40px"})
    }
  }
  // const [check, setCheck] = useState("0px")
  // const handleCheck = () => {
  //   toggleCheck(!isCheck)
  //   if (isCheck) {
  //     setCheck("0px")
  //   } else {
  //     setCheck("-50px")
  //   }
  // }

  return (
    <button className="menu-bar">
      <figure>
        {/* <img src={props.image} alt="メニューイメージ" /> */}
      </figure>
      <div>
        <h2>{props.name}</h2>
      </div>
      <button className="toggleBtn">
        <div style={checkStyle} onClick={handleCheck}></div>
      </button>
    </button>
  )
}

export default MenuBar;