import { connect_telegram } from "./connect_telegram"
import create_user from "./create"
import get_user from "./get"
import list_users from "./list"
import list_permissions from "./list_permissions"

export const user = {
   connect_telegram: connect_telegram,
   create: create_user,
   get: get_user,
   list_permissions: list_permissions,
   list: list_users,
}
