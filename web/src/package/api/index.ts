import Axios, { AxiosPromise } from 'axios';
import { beauticianMenuResponse, menuDetailsResponse, menusResponse } from '../interface/response/Menu';
import { reservationsResponse, reservationResponse, guestMypageReservationsResponse, reservationInfoResponse } from '../interface/response/Reservation'
import { salonsResponse } from '../interface/response/Salon';
import qs from "qs"
import { userResponse, usersResponse } from '../interface/response/User';
import { userCreateRequest } from '../interface/request/User';
import { getAuthToken } from '../../utils/function/Cookie'
import { apiurl } from '../../config/config';
import { beauticianCreateRequest, beauticianUpdateRequest } from '../interface/request/Beautician';
import { setHolidayRequest } from '../interface/request/Reservation';
import { beauticianMenuRequest } from '../interface/request/Menu';

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

const put = <T>(url: string, data?: object): Promise<T> => {
  return requestAwait(axios.put(url, data))
}

export const findSalons = async (id?: string): Promise<salonsResponse> => {
  return get<salonsResponse>("/api/v1/salon/find", {
    params: {
      beauticianRandId: id
    }
  })
}

export const findBeauticians = async (date?: string, menuIDs?: string[], salonRandID?: string): Promise<usersResponse> => {
  return get<usersResponse>("/api/v1/beautician/find", {
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

export const getMe = async (): Promise<userResponse> => {
  return get<userResponse>(`api/v1/user`)
}

export const createUser = async (guest: userCreateRequest): Promise<userResponse> => {
  return post<userResponse>(`api/v1/user`, guest)
}

export const getGuestMypage = async (): Promise<guestMypageReservationsResponse> => {
  return get<guestMypageReservationsResponse>(`api/v1/reservation/user`)
}

export const createBeautician = async (request: beauticianCreateRequest): Promise<userResponse> => {
  return post<userResponse>(`api/v1/beautician`, request)
}

export const updateBeautician = async (request: beauticianUpdateRequest) => {
  return put(`api/v1/beautician`, request)
}

export const getReservationBeautician = async (): Promise<reservationsResponse> => {
  return get<reservationsResponse>(`api/v1/reservation/beautician`)
}

export const getReservationInfo = async (randID: string): Promise<reservationInfoResponse> => {
  return get<reservationInfoResponse>(`api/v1/reservation/${randID}`)
}

export const setHoliday = async (req: setHolidayRequest): Promise<reservationResponse> => {
  return post<reservationResponse>(`api/v1/reservation/beautician`, req)
}

export const postBeauticianMenu = async (req: beauticianMenuRequest): Promise<beauticianMenuResponse> => {
  return post<beauticianMenuResponse>(`api/v1/menu/beautician`, req)
}