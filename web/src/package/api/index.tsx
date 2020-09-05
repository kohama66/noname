import Axios from 'axios';

const axios = Axios.create({
  baseURL: "http://localhost:8080",
  headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
  }
})

export const getMe = () => (
  axios.get("/api/v1/beautician", {})
  .then((res) => {
    console.log("ok")
    console.log(res)
  })
  .catch((error) => {
    console.log("NG")
    console.log(error)
  })
);
