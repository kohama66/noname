// 2000-00-00T00:00:00Z形式の月を出力します
export const getMonth = (date: string): RegExpMatchArray | null => {
  const month = date.match(/(?<=-).*?(?=-)/)
  return month
}
// 2000-00-00T00:00:00Z形式の日を出力します
export const getDay = (date: string): RegExpMatchArray | null => {
  var month = date.match(/(?<=-\d\d-)[1-9]\d(?=T)/)
  if (month == null) {
    var month = date.match(/(?<=-\d\d-0)\d(?=T)/)
    return month
  }
  return month
}
// 2000-00-00T00:00:00Z形式の時間を出力します
export const getHours = (date: string): RegExpMatchArray | null => {
  var hours = date.match(/(?<=T)[1-9]\d(?=:)/)
  if (hours == null) {
    var hours = date.match(/(?<=T0)\d(?=:)/)
    return hours
  }
  return hours
}
// 2000-00-00T00:00:00Zへ変換します
export const dateToString = (date: Date): string => {
  var y = date.getFullYear()
  var m = date.getMonth() + 1
  var d = date.getDate()
  var h = date.getHours()
  var mi = date.getMinutes()
  var s = date.getSeconds()

  var format_str = 'YYYY-MM-DDThh:mm:ssZ';
  format_str = format_str.replace(/YYYY/g, y.toString());
  format_str = format_str.replace(/MM/g, (("0" + m).slice(-2)).toString());
  format_str = format_str.replace(/DD/g, (("0" + d).slice(-2)).toString());
  format_str = format_str.replace(/hh/g, (("0" + h).slice(-2)).toString());
  format_str = format_str.replace(/mm/g, (("0" + mi).slice(-2)).toString());
  format_str = format_str.replace(/ss/g, (("0" + s).slice(-2)).toString());
  return format_str
}