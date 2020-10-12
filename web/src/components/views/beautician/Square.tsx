import React, { FC } from 'react';

interface SqusreProps {
  id?: number;
}

const Square: FC<SqusreProps> = ({id}) => {
  return (
  <td>{id}</td>
  )
}

export default Square;
