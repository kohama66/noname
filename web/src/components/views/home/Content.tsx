import React, { FC } from 'react';

interface ContentProps {
  title: string;
}

const Content: FC<ContentProps> = ({title}) => (
    <div>
      <img src="/img/cut.png" alt="" />
      <h2>{title}</h2>
    </div>
  );

export default Content;
