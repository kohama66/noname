import React, { FC, useContext, useEffect, useState } from 'react';
import { getMe } from '../../../package/api';
import { UserContext } from '../../../utils/context/UserContext';
import App from '../../App';

const Loading: FC = () => {
  const { setUser } = useContext(UserContext)
  const [isCheckLogin, setCheck] = useState<boolean>(false)

  useEffect(() => {
    const checkUser = async () => {
      try {
        const response = await getMe()
        setUser(response.user)
      } catch (error) {
        console.log(error)
      }
      setCheck(true)
    }
    checkUser()
  }, [])

  if (!isCheckLogin) return <div></div> //ログイン確認してから<App>へ移行します
  return <App />
}

export default Loading;