import React, { FC, useEffect, useState } from 'react';
import { getGuest } from '../../../../package/api';
import { GuestByMyPage, initGuestByMyPage } from '../../../../package/interface/Guest';
import ReservationInfor from '../../parts/ReservationInfor/ReservationInfor';
import Title from '../parts/Title/Title';
import './Mypage.scss'

const Mypage: FC = () => {
  const [me, setMe] = useState<GuestByMyPage>(initGuestByMyPage)
  const handleGetMe = async () => {
    try {
      const response = await getGuest()
      setMe(response)
    } catch (error) {
      console.log(error)
    }
  }
  useEffect(() => {
    handleGetMe()
  }, [])

  return (
    <div id="guest-mypage">
      <Title title="MY PAGE" text="マイページ" />
      <section>
        <div className="profile-card">
          <figure></figure>
          <div>
            <dl>
              <span>
                <dt>名前</dt>
                <dd>{me.lastName + " " + me.firstName}</dd>
              </span>
              <span>
                <dt>歳</dt>
                <dd>29</dd>
              </span>
              <span>
                <dt>性別</dt>
                <dd>man</dd>
              </span>
              <span>
                <dt>電話</dt>
                <dd>09012345678</dd>
              </span>
            </dl>
          </div>
        </div>
        <div className="reserved-history-card">
          <div>
            <h2>現在の予約</h2>
            <div>
              <dl>
                <span>
                  <dt>店舗</dt>
                  <dd>{(() => {
                    if (me != initGuestByMyPage) {
                      return me.reservations[0].salonName
                    }
                  })()}</dd>
                </span>
                <span>
                  <dt>美容師</dt>
                  <dd>山田 隆</dd>
                </span>
                <span>
                  <dt>日付</dt>
                  <dd>11月05日</dd>
                </span>
                <span>
                  <dt>時間</dt>
                  <dd>15:00から</dd>
                </span>
                <span className="menus">
                  <dt>メニュー</dt>
                  <span>
                    <dd>カット</dd>
                    <dd>カラー</dd>
                    <dd>パーマ</dd>
                  </span>
                </span>
                <span>
                  <dt>合計</dt>
                  <dd>5500</dd>
                </span>
              </dl>
            </div>
          </div>
          {(() => {
            if (me != initGuestByMyPage) {
              return <ReservationInfor titleText="前回の予約" storeName={me.reservations[1].salonName}
                beauticianLastName={me.reservations[1].beauticianLastName} beauticianFirstName={me.reservations[1].beauticianFirstName} />
            }
          })()}
        </div>
      </section>
    </div>
  )
}

export default Mypage;