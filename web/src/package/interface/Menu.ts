export interface Menu {
  randId: string
  name: string
}

export const isMenuInterface = (arg: any): arg is Menu => {
  return arg !== null &&
    typeof arg === "object" &&
    typeof arg.randID === "string" && typeof arg.name === "string"
}

// export const isMenusInterface = (arg: any): arg is Menu[] => {
//   return arg !== null &&
//     typeof arg === "object" &&
//     typeof arg[0].randID === "string" && typeof arg[0].name === "string"
// }