export interface Menu {
  randId: string
  name: string
}

export interface MenuDetail {
  randId: string
  menuId: number
  beauticianId: number
  price: number
  name: string
}

export interface BeauticianMenu {
  randId: string
  beauticianId: number
  menuId: number
  name: string
  price: number
}

export const initMenu: Menu = <Menu>{}

export const initMenuDetail: MenuDetail = <MenuDetail>{}

export const isMenuInterface = (arg: any): arg is Menu => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.randID === "string" && typeof arg.name === "string"
}
