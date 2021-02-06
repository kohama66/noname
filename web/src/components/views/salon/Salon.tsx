import React, { FC } from 'react'
import { Route, Switch, useRouteMatch } from 'react-router-dom'
import SignUp from './signup/SignUp'

const Salon: FC = () => {
  const match = useRouteMatch().path
  return (
    <div id="salon">
      <Switch>
        <Route path={match + "/signup"} component={SignUp} />
      </Switch>
    </div>
  )
}

export default Salon
