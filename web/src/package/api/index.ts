import Axios, { AxiosPromise } from 'axios';
import { beauticianResponse, beauticiansResponse } from '../interface/response/Beautician';
import { menuDetailsResponse, menusResponse } from '../interface/response/Menu';
import { reservationsResponse, reservationResponse, guestMypageReservationsResponse } from '../interface/response/Reservation'
import { salonsResponse } from '../interface/response/Salon';
import qs from "qs"
import { guestResponse } from '../interface/response/Guest';
import { guestCreateRequest } from '../interface/request/Guest';
import {getAuthToken} from '../../utils/function/Cookie'
import { apiurl } from '../../config/config';

const axios = Axios.create({
  baseURL: apiurl,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json"
  },
  paramsSerializer: (params: any) => {
    return qs.stringify(params, { arrayFormat: 'repeat' })
  }
})
axios.interceptors.request.use(
  config => {
    const token = getAuthToken()
    config.headers.Authorization = `Bearer ${token}`
    return config
  }
)

const requestAwait = async <T>(request: Promise<AxiosPromise<T>>): Promise<T> => {
  try {
    const response = await request
    return response.data
  } catch (error) {
    // console.log(error.response)
    let errorTetx: string
    switch (error.message) {
      case "Request failed with status code 500":
        errorTetx = "エラーが発生しました、管理者に問い合わせて下さい。"
        break
      default:
        errorTetx = "エラーが発生しました"
        break
    }
    throw new Error(errorTetx)
  }
}

const get = <T>(url: string, data?: object): Promise<T> => {
  return requestAwait(axios.get(url, data))
}

const post = <T>(url: string, data?: object): Promise<T> => {
  return requestAwait(axios.post(url, data))
}

export const findSalons = async (id?: string): Promise<salonsResponse> => {
  return get<salonsResponse>("/api/v1/salon/find", {
    params: {
      beauticianRandId: id
    }
  })
}

export const findBeauticians = async (date?: string, menuIDs?: string[], salonRandID?: string): Promise<beauticiansResponse> => {
  return get<beauticiansResponse>("/api/v1/beautician/find", {
    params: {
      date: date,
      menuRandIds: menuIDs,
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

export const findMenuDetails = async (beauticianID: string, menuIDs?: string[]): Promise<menuDetailsResponse> => {
  return get<menuDetailsResponse>(`api/v1/menu/find/${beauticianID}`, {
    params: {
      menuRandIds: menuIDs,
    }, paramsSerializer: (params: any) => {
      return qs.stringify(params, { arrayFormat: 'repeat' })
    }
  })
}

export const createReservation = async (beauticianRandID: string, salonRandID: string, menuIDs: string[], date: string): Promise<reservationResponse> => {
  return post<reservationResponse>(`api/v1/reservation`, {
    beauticianRandId: beauticianRandID,
    salonRandId: salonRandID,
    menuIds: menuIDs,
    date: date,
  })
}

export const getGuest = async (): Promise<guestResponse> => {
  return get<guestResponse>(`api/v1/guest`)
}

export const createGuest = async (guest: guestCreateRequest): Promise<guestResponse> => {
  return post<guestResponse>(`api/v1/guest`, guest)
}

export const getGuestMypage = async (): Promise<guestMypageReservationsResponse> => {
  return get<guestMypageReservationsResponse>(`api/v1/reservation/guest`)
}
