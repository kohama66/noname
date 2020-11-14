import { Guest, GuestByMyPage } from "../Guest";

export interface guestResponse {
  guest: Guest
}

export interface guestMypageResponse {
  randId: string
  firstName: string
  lastName: string
  reservations: {
    id: number
    date: Date
    salonName: string
    beauticianFirstName: string
    beauticianLastName: string
  }[]
}