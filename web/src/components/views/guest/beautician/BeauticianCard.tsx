import React, { FC, useContext, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
// import { setCheckedBeauticianContext } from '../../../../components/views/guest/Guest'
import {setCheckedContext} from '../../../../container/views/guest/Guest'

type props = {
  lastName: string
  randId: string
  // setState: (props: string) => void 
}

const BeauticianCard: FC<props> = (props) => {
  const [randID, setRandId] = useState(props.randId)
  const history = useHistory();
  const setChecked = useContext(setCheckedContext)

  const handleBeautician = () => {
    setChecked("beautician")
    history.push("/guest")
  }

  return (
      <div className="beautician-card" onClick={handleBeautician} >
        <figure>
          <img src="/img/bijin.jpg" alt="" />
        </figure>
        <div>
          <h2>{props.lastName}</h2>
          <p>美容師歴5年</p>
          <h3>得意なスタイル</h3>
          <p>〇〇カット、カラー</p>
          <h4>お客様に最適なスタイルを提案致します</h4>
        </div>
      </div>
  )
}

export default BeauticianCard;