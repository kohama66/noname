import React, { FC, useContext } from 'react';
import { Redirect } from 'react-router-dom';
import { UserContext } from '../../../utils/context/UserContext';

const Auth: FC = ({children}) => {
  const { user } = useContext(UserContext)
  if (user) return <>{children}</>
  return <Redirect to="/login" />
}

export default Auth;