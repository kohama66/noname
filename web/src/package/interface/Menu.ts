export interface Menu {
  randId: string
  name: string
}

export interface MenuDetail {
  menuId: number
  beauticianId: number
  price: number
  name: string
}

export const initMenu: Menu = <Menu>{}

export const initMenuDetail: MenuDetail = <MenuDetail>{}

export const isMenuInterface = (arg: any): arg is Menu => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.randID === "string" && typeof arg.name === "string"
}
