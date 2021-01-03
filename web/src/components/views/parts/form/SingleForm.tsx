import React, { FC, useContext, useEffect, useState } from 'react';
import { getMe, postBeauticianMenu } from '../../../../package/api';
import { UserContext } from '../../../../utils/context/UserContext';
import { useError } from '../../../../utils/hooks/Error';
import FormErrorMessage from '../formParts/FormErrorMessage';
import Input from '../formParts/Input';
import Select from '../formParts/Select';
import "./SingleForm.scss";

interface props {
  type: "menu" | "salon"
  disabled: boolean
}

const SingleForm: FC<props> = (props) => {
  const [stateName, setStateName] = useState<string>("")
  const [statePrice, setStatePrice] = useState<string>("")
  const [stateType, setStateType] = useState<string>("")
  const { error, customError } = useError()
  const { setUser } = useContext(UserContext)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    console.log(stateType)
    switch (props.type) {
      case "menu":
        try {
          await postBeauticianMenu({
            menuName: stateName,
            menuCategory: stateType,
            price: Number(statePrice)
          })
          const response = await getMe()
          setUser(response.user)
          customError("")
        } catch (error) {
          customError("エラーです、やり直して下さい")
        }
        break
    }
  }

  useEffect(() => {

  },[])

  return (
    <form id="single-form" onSubmit={handleSubmit}>
      <Input type="text" value={stateName} setState={setStateName} disabled={props.disabled} required={true} placeHolder="新しいメニューを追加する" />
      <Input type="text" value={statePrice} setState={setStatePrice} disabled={props.disabled} required={true} placeHolder="値段設定" />
      <Select type="menu" value={stateType} setState={setStateType} />
      <Input type="submit" value="登録" />
      <FormErrorMessage error={error} />
    </form>
  )
}

export default SingleForm;