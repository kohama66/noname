import React, { FC, useContext } from 'react';
import { Redirect } from 'react-router-dom';
import { UserContext } from '../../../utils/context/UserContext';

const BtAuth: FC = ({children}) => {
  const { user } = useContext(UserContext)
  if (user.isBeautician) return <>children</>
  return <Redirect to="/login" />
}

export default BtAuth;