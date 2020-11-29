// 2000-00-00T00:00:00Z形式の月を出力します
export const getMonth = (date: string): RegExpMatchArray | null => {
  const month = date.match(/(?<=-).*?(?=-)/)
  return month
}
// 2000-00-00T00:00:00Z形式の日を出力します
export const getDay = (date: string): RegExpMatchArray | null => {
  var month = date.match(/(?<=-\d\d-)[1-9]\d(?=T)/)
  if (month == null){
    var month = date.match(/(?<=-\d\d-0)\d(?=T)/)
    return month
  }
  return month
}
// 2000-00-00T00:00:00Z形式の時間を出力します
export const getHours = (date: string): RegExpMatchArray | null => {
  var hours = date.match(/(?<=T)[1-9]\d(?=:)/)
  if (hours == null){
    var hours = date.match(/(?<=T0)\d(?=:)/)
    return hours
  }
  return hours
}