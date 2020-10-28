import React, { FC } from 'react';
import { Link, useHistory, useRouteMatch } from 'react-router-dom';
import "./Popup.scss"

interface props {
  display: boolean
}

const Popup: FC<props> = (props) => {
  const match = useRouteMatch()
  const history = useHistory()

  return (
    <article id="popup" className={props.children ? "" : "none" }>
      <div>
        <Link to={match.path + "/final_comfirmation"}>
          最終確認画面
        </Link>
      </div>
    </article>
  )
}

export default Popup;