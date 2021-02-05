import { createContext, useState } from 'react';
import { initUser, User } from '../../package/interface/User';

export const UserContext = createContext({
  user: initUser,
  setUser: (user: User) => { },
})

export const useUserContext = () => {
  const [user, settingUser] = useState<User>(initUser)
  const setUser = (user: User): void => {
    settingUser(user)
  }
  return { user, setUser }
}
