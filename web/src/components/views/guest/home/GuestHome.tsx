import React, { FC, useContext, useEffect } from 'react';
import "./GuestHome.scss";
import { useHistory } from 'react-router-dom';
import { ReservedContext } from '../../../../utils/context/ReservadContext ';
import ChoosePlate from '../parts/ChoosePlate';
import Title from '../parts/Title/Title';

const GuestHome: FC = () => {
  const { isAllChecked } = useContext(ReservedContext)
  const history = useHistory()
  useEffect(() => {
    if (isAllChecked()) {
      history.push("guest/final_comfirmation")
    }
  }, [])
  return (
    <article id="chooses">
      <Title text="何から選びますか？" title="SELECT" />
      <div className="chooses-contents-wrapper">
        <div className="chooses-contents">
          <ChoosePlate type="store" />
          <ChoosePlate type="beautician" />
          <ChoosePlate type="menu" />
          <ChoosePlate type="date" />
        </div>
      </div>
    </article>
  )
}

export default GuestHome;