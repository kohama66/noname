export interface Guest {
  randId: string
  firstName: string
  lastName: string
}

export interface GuestByMyPage {
  randId: string
  firstName: string
  lastName: string
  reservations: {
    id: number
    date: string
    salonName: string
    beauticianFirstName: string
    beauticianLastName: string
  }[]
}

export const initGuest: Guest = <Guest>{}
export const initGuestByMyPage: GuestByMyPage = <GuestByMyPage>{}
