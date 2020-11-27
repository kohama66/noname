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
    throw new Error(error)
  }
}

export const signup = async (firebaseUser: FirebaseUser) => {
  try {
    const authResponse: firebase.default.auth.UserCredential = await auth.createUserWithEmailAndPassword(firebaseUser.email, firebaseUser.password)
    await authCookieSet(authResponse.user)
  } catch (error) {
    throw new Error(error)
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

// cookie削除
export const deleteToken = () => {
  Cookie.set("authHeader", "", { expires: 1 })
}