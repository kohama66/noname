import React, { FC } from 'react'
import { Route, Switch, useRouteMatch } from 'react-router-dom'
import MyPage from './mypage/MyPage'
import SignUp from './signup/SignUp'

const Salon: FC = () => {
  const match = useRouteMatch().path
  return (
    <Switch>
      <Route path={match + "/signup"} component={SignUp} />
      <Route path={match + "/mypage"} component={MyPage} />
    </Switch>
  )
}

export default Salon
