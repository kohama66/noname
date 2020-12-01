import React, { FC, useContext, useEffect, useState } from 'react';
import { getGuest } from '../../../package/api';
import { GuestContext } from '../../../utils/context/GuestContext';
import App from '../../App';

const Loading: FC = () => {
  const { setGuest } = useContext(GuestContext)
  const [isCheckLogin, setCheck] = useState<boolean>(false)

  useEffect(() => {
    const checkGuest = async () => {
      try {
        const response = await getGuest()
        setGuest(response.guest)
      } catch (error) {
        console.log(error)
      }
      setCheck(true)
    }
    checkGuest()
  }, [])

  if(!isCheckLogin) return <div></div> //ログイン確認してから<App>へ移行します
  return <App />
}

export default Loading;