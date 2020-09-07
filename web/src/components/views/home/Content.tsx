import React, { FC } from 'react';
import { Link } from "react-router-dom";

interface ContentProps {
  title: string;
  path: string;
}

const Content: FC<ContentProps> = ({ title, path }) => (
  <div>
    <img src="/img/cut.png" alt="" />
    <Link to={path}>{title}</Link>
  </div>
);

export default Content;
