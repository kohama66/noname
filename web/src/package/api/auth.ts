import { myFirebase } from "../../lib/firebase/firebase"
import { FirebaseUser } from "../interface/Firebase"
import Cookie from "js-cookie"
import { createGuest } from "."

const auth = myFirebase.auth()

export const signin = async (firebaseUser: FirebaseUser) => {
  try {
    const authResponse: firebase.default.auth.UserCredential = await auth.signInWithEmailAndPassword(firebaseUser.email, firebaseUser.password)
    await authCookieSet(authResponse.user)
  } catch (error) {
    let errorTetx: string
    switch (error.code) {
      case "auth/user-not-found":
        errorTetx = "ユーザーが見つかりませんでした。"
        break
      case "auth/wrong-password":
        errorTetx = "パスワードが一致しません。"
        break
      default:
        errorTetx = "エラーが発生しました"
        break
    }
    throw new Error(errorTetx)
  }
}

export const signup = async (firebaseUser: FirebaseUser) => {
  try {
    const authResponse: firebase.default.auth.UserCredential = await auth.createUserWithEmailAndPassword(firebaseUser.email, firebaseUser.password)
    await authCookieSet(authResponse.user)
    
  } catch (error) {
    let errorTetx: string
    switch (error.code) {
      case "auth/network-request-failed":
        errorTetx = "ネットワークエラーです、接続を確認して下さい。"
        break
      case "auth/email-already-in-use":
        errorTetx = "このメールアドレスは既に登録されています。"
        break
      default:
        errorTetx = "エラーが発生しました"
        break
    }
    throw new Error(errorTetx)
  }
}

const authCookieSet = async (user: firebase.default.User | null): Promise<void> => {
  if (user) {
    const token = await user.getIdToken()
    Cookie.set("authHeader", token, { expires: 1 })
  }
}

// firebase ユーザー消去
export const deleteUser = async () => {
  try {
    const user = auth.currentUser
    if (user) {
      await user.delete()
    }
  } catch (err) {
    console.log(err)
  }
}
