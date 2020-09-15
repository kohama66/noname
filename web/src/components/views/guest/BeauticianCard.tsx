import React, { FC, useState } from 'react';

type props = {
  lastName: string
  key: number
  randId: string
}

const BeauticianCard: FC<props> = (props) => {
  const [randID, setRandId] = useState(props.randId)
  
  return (
    <div className="beautician-card" >
      <figure></figure>
      <div>
        <h2>{props.lastName}</h2>
      </div>
    </div>
  )
}

export default BeauticianCard;