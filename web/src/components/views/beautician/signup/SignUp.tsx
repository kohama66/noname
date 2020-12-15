import React, { FC, useState } from 'react';
import { deleteUser, signup } from '../../../../package/api/auth';
import { deleteAuthToken } from '../../../../utils/function/Cookie';
import { useError } from '../../../../utils/hooks/Error';
import Title from '../../guest/parts/Title/Title';
import Form from '../../parts/form/Form';
import './SignUp.scss';

const SignUp: FC = () => {
  const [firstName, setFirstName] = useState<string>("")
  const [lastName, setLastName] = useState<string>("")
  const [firstNameKana, setFirstNameKana] = useState<string>("")
  const [lastNameKana, setLastNameKana] = useState<string>("")
  const [email, setEmail] = useState<string>("")
  const [passwoed, setPassword] = useState<string>("")
  const [disabled, setDisabled] = useState<boolean>(false)
  const {error, customError} = useError()

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault()
    setDisabled(true)
    try {
      await signup({
        email: email,
        password: passwoed
      })
      try {
        // const response = await createGuest({
        //   lastName: lastName,
        //   firstName: firstName,
        //   lastNameKana: lastNameKana,
        //   firstNameKana: firstNameKana,
        //   email: email
        // })
        // setGuest(response.guest)
        // history.push("/guest")
      } catch (error) {
        deleteAuthToken()
        deleteUser()
        throw new Error(error.message)
      }
    } catch (error) {
      customError(error.message)
      setDisabled(false)
    }
  }

  return (
    <div id="bt-signup">
      <Title title="SIGNUP" text="美容師登録" />
      <Form handleSubmit={handleSubmit} lastName={lastName} lastNameKana={lastNameKana} firstName={firstName} firstNameKana={firstNameKana} email={email} password={passwoed} disabled={disabled} error={error}
        setLastName={setLastName} setLastNameKana={setLastNameKana} setFirstName={setFirstName} setFirstNameKana={setFirstNameKana} setEmail={setEmail} setPassword={setPassword} />
    </div>
  )
}

export default SignUp;