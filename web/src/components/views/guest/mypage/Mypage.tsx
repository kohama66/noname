import React, { FC } from 'react';
import './Mypage.scss'

const Mypage: FC = () => {
  return (
    <div id="guest-mypage">
      <div className="profile-card">
        <figure></figure>
        <div>
          <dl>
            <span>
              <dt>Name</dt>
              <dd>小濱 陵</dd>
            </span>
            <span>
              <dt>Age</dt>
              <dd>29</dd>
            </span>
            <span>
              <dt>Gender</dt>
              <dd>man</dd>
            </span>
          </dl>
        </div>
      </div>
      <div className="reserved-history-card">

      </div>
    </div>
  )
}

export default Mypage;