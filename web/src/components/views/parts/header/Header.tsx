import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import "./Header.scss"

const Header: FC = () => {
  return (
    <header>
      {/* <Link to="/" >TOP</ Link> */}
      <div className="header-login">
        <Link to="/guest/mypage" className="mypage-btn" >マイページ</Link>
      </div>
    </header>
  )
}

export default Header;