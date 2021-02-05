import { BeauticianInfo } from "../Beautician";
import { BeauticianMenu } from "../Menu";
import { Salon } from "../Salon";
import { User } from "../User";

export interface BeauticianMyPage {
  user: User
  beautician: BeauticianInfo
  salons: Salon[]
  beauticianMenus: BeauticianMenu[]
}