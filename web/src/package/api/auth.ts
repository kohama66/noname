import { myFirebase } from "../../lib/firebase/firebase"

const auth = myFirebase.auth()

export const signin = async () => {
  try {
    const authResponse: firebase.default.auth.UserCredential = await auth.signInWithEmailAndPassword("test@test.com", "00000000")
    if (authResponse.user) {
      const uid = await authResponse.user.getIdToken()
      console.log(authResponse.user.uid)
    }
  } catch (error) {
    console.log(error)
  }
}