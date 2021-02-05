import Cookie from "js-cookie";

export const getAuthToken = (): string | undefined => {
  return Cookie.get("authHeader");
};

// cookie削除
export const deleteAuthToken = () => {
  Cookie.set("authHeader", "", { expires: 5 })
}