import { createContext, useContext, useState } from "react"
import { Beautician, initBeautician } from "../../package/interface/Beautician"

export const BeauticainContext = createContext({
  beautician: initBeautician,
  setBeauticain: (beautician: Beautician) => { },
})

export const useBeauticianContext = () => {
  const [beautician, settingBeautican] = useState<Beautician>(initBeautician)
  const setBeauticain = (beautician: Beautician) => {
    settingBeautican(beautician)
  }
  return { beautician, setBeauticain }
}