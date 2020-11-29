import { useState } from "react"

export const useError = () => {
  const [error, setError] = useState<string>()
  const customError = (e: string) => {
    setError(e)
  }
  return { error, customError }
}
