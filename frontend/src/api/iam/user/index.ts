import { connect_telegram } from "./connect_telegram"
import create_user from "./create"
import get_user from "./get"
import list_users from "./list"
import list_permissions from "./list_permissions"
import get_user_profile from "./profile"

export const user = {
   connect_telegram: connect_telegram,
   create: create_user,
   get: get_user,
   profile: get_user_profile,
   list_permissions: list_permissions,
   list: list_users,
}
