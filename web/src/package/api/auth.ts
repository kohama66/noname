import { myFirebase } from "../../lib/firebase/firebase"
import { FirebaseUser } from "../interface/Firebase"
import Cookie from "js-cookie"

const auth = myFirebase.auth()

export const signin = async (firebaseUser: FirebaseUser) => {
  try {
    const authResponse: firebase.default.auth.UserCredential = await auth.signInWithEmailAndPassword(firebaseUser.email, firebaseUser.password)
    if (authResponse.user) {
      const token = await authResponse.user.getIdToken()
      Cookie.set("authHeader", token, { expires: 1 })
      return
    }
  } catch (error) {
    throw new Error(error)
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