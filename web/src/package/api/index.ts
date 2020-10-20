import Axios, { AxiosPromise } from 'axios';
import { BeauticianGetAll } from '../interface/Beautician';
import { beauticiansResponse } from '../interface/response/Beautician';
import { ReservationFindByBeautician } from '../interface/response/Reservation'
import { salonResponse, salonsResponse } from '../interface/response/Salon';

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
    return response.data
  } catch (err) {
    switch (err) {
      case err.response:
        console.log(err.response)
      case err.request:
        console.log(err.request)
      default:
        console.log('Error', err.message)
    }
  }
}

const requestAwait = async <T>(request: Promise<AxiosPromise<T>>): Promise<T> => {
  try {
    const response = await request
    return response.data
  } catch (err) {
    throw new Error(err)
  }
}

const get = <T>(url: string, data?: object): Promise<T> => {
  return requestAwait(axios.get(url, data))
}

export const findSalons = async (id: string): Promise<salonsResponse> => {
  return get<salonsResponse>("/api/v1/salon/find", {
    params: {
      beauticianRandId: id
    }
  })
}

export const findBeauticians = async (date?: string, menuRandID?: string, salonRandID?: string): Promise<beauticiansResponse> => {
  return get<beauticiansResponse>("/api/v1/beautician/find", {
    params: {
      date: date,
      menuRandID: menuRandID,
      salonRandID: salonRandID
    }
  })
}