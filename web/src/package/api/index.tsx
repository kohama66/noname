import Axios from 'axios';
import { ReservationFindByBeautician, BeauticianGetAll } from '../interface/response/Reservation'

const axios = Axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
    "X-DEBUG-ID": "test1"
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

export const getReservationBeautician = async () => {
  try {
    const response = await axios.get<ReservationFindByBeautician>("/api/v1/reservation/beautician")
    const { data, status } = response
    switch (status) {
      case 200:
        if (data) return data
        return
      default:
        console.log("200以外")
        if (data) return data
        return
    }
  } catch (err) {
    console.log(err)
  }
};

export const getAllBeauticians = async () => {
  try {
    const response = await axios.get<BeauticianGetAll>("api/v1/beautician/all")
    const { data, status } = response
    switch (status) {
      case 200:
        if (data) return data
        return
      default:
        console.log("200以外")
        if (data) return data
        return
    }
  } catch (err) {
    console.log(err)
  }
}