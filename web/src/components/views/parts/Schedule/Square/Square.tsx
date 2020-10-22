import React, { FC } from 'react';

interface props {
  clickFunction: () => void
  reserved: boolean
}

const Square: FC<props> = (props) => {
  return (
    <td onClick={() => props.clickFunction()}>{(() => {
      if(props.reserved){
        return "×"
      }
    })()}</td>
  )
}

export default Square;
