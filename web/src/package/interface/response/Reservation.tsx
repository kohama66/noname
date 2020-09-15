export interface Reservation {
  date: string;
  time: string;
}

export interface ReservationFindByBeautician {
  Reservations: Reservation[]
}

export interface Beautician {
  lastName: string;
}

export interface BeauticianGetAll {
  Beauticians: Beautician[]
}