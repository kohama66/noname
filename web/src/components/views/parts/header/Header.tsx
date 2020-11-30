import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import "./Header.scss"

const Header: FC = () => {
  return (
    <header>
      <div className="header-top-icon">
        <Link to="/guest" >Cut Match</ Link>
      </div>
      <div className="header-login">
        <Link to="/guest/mypage" className="mypage-btn" >マイページ</Link>
      </div>
    </header>
  )
}

export default Header;