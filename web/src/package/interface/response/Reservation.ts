export interface Reservation {
  date: string;
  time: string;
}

export interface ReservationFindByBeautician {
  Reservations: Reservation[]
}

export interface Beautician {
  randId: string;
  lastName: string;
}

export interface BeauticianGetAll {
  beauticians: Beautician[]
}
