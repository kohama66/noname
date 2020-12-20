import React, { FC, useContext, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { initUser } from '../../../../package/interface/User';
import { UserContext } from '../../../../utils/context/UserContext';
import MypageComponent from './Mypage'

const Mypage: FC = () => {
  const { user } = useContext(UserContext)
  const history = useHistory()

  useEffect(() => {
    if(user === initUser){
      history.push("/")
    }
  },[])
  
  if (user === initUser) return null
  return <MypageComponent />
}

export default Mypage;