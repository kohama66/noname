import React, { FC, useContext } from 'react';
import { Link } from 'react-router-dom';
import { initUser } from '../../../../package/interface/User';
import { UserContext } from '../../../../utils/context/UserContext';
import "./Header.scss"

const Header: FC = () => {
  const { user } = useContext(UserContext)
  return (
    <header>
      <div className="header-top-icon">
        <Link to="/guest" >Cut Match</ Link>
      </div>
      <div className="header-login">
        <Link to={(() => {
          if (user !== initUser) {
            return "/guest/mypage"
          } else {
            return "/login"
          }
        })()} className="mypage-btn" >My Page</Link>
      </div>
    </header>
  )
}

export default Header;