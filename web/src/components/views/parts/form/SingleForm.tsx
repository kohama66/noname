import React, { FC, useContext, useEffect, useState } from 'react';
import { findMenus, getMe, postBeauticianMenu } from '../../../../package/api';
import { Menu } from '../../../../package/interface/Menu';
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
  const [menus, setMenus] = useState<Menu[]>([])
  const { error, customError } = useError()
  const { setUser } = useContext(UserContext)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
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
          setStateName("")
          setStatePrice("")
          setStateType("")
          customError("")
        } catch (error) {
          customError("エラーです、やり直して下さい")
        }
        break
    }
  }

  useEffect(() => {
    const handleFindMenus = async () => {
      try {
        const response = await findMenus()
        setMenus(response.menus)
      } catch (error) {
        console.log(error)
      }
    }
    handleFindMenus()
  }, [])

  return (
    <form id="single-form" onSubmit={handleSubmit}>
      <Input type="text" value={stateName} setState={setStateName} disabled={props.disabled} required={true} placeHolder="新しいメニューを追加する" />
      <Input type="text" value={statePrice} setState={setStatePrice} disabled={props.disabled} required={true} placeHolder="値段設定" />
      <Select defaultValueLabel="カテゴリー" optionData={menus} value={stateType} setState={setStateType} required={true} />
      <Input type="submit" value="登録" />
      <FormErrorMessage error={error} />
    </form>
  )
}

export default SingleForm;