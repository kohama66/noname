import React, { FC, useContext } from 'react';
import { Link } from 'react-router-dom';
import { initGuest } from '../../../../package/interface/Guest';
import { GuestContext } from '../../../../utils/context/GuestContext';
import "./Header.scss"

const Header: FC = () => {
  const { guest } = useContext(GuestContext)
  return (
    <header>
      <div className="header-top-icon">
        <Link to="/guest" >Cut Match</ Link>
      </div>
      <div className="header-login">
        <Link to={(() => {
          if(guest !== initGuest){
            return "/guest/mypage"
          } else {
            return "/login"
          }
        })()} className="mypage-btn" >マイページ</Link>
      </div>
    </header>
  )
}

export default Header;