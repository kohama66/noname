import React, { FC } from 'react';

interface props {
  clickFunction: () => void
}

const Square: FC<props> = (props) => {
  return (
    <td onClick={() => props.clickFunction()}></td>
  )
}

export default Square;
