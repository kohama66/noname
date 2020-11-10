import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import "./Header.scss"

const Header: FC = () => {
  return(
    <header>
      <Link to="/" >TOP</ Link>
    </header>
  )
}

export default Header;