import Axios, { AxiosPromise } from 'axios';
import { beauticianResponse, beauticiansResponse } from '../interface/response/Beautician';
import { menusResponse } from '../interface/response/Menu';
import { reservationsResponse } from '../interface/response/Reservation'
import { salonsResponse } from '../interface/response/Salon';
import qs from "qs"

const axios = Axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
    "X-DEBUG-ID": "test1"
  }
})

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

export const findBeauticians = async (date?: string, menuRandIDs?: string[], salonRandID?: string): Promise<beauticiansResponse> => {
  return get<beauticiansResponse>("/api/v1/beautician/find", {
    params: {
      date: date,
      menuRandIds: menuRandIDs,
      salonRandId: salonRandID
    }, paramsSerializer: (params: any) => {
      return qs.stringify(params, { arrayFormat: 'repeat' });
    }
  })
}

export const findMenus = async (beauticianRandID?: string): Promise<menusResponse> => {
  return get<menusResponse>("/api/v1/menu/find", {
    params: {
      beauticianRandId: beauticianRandID
    }
  })
}

export const findReservation = async (beauticianRandID?: string): Promise<reservationsResponse> => {
  return get<reservationsResponse>('/api/v1/reservation/find', {
    params: {
      beauticianRandId: beauticianRandID
    }
  })
}

export const getBeautician = async (randID: string): Promise<beauticianResponse> => {
  return get<beauticianResponse>(`api/v1/beautician/${randID}`)
}