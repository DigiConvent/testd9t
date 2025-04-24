import type { ApiCollection } from "@/api/core/endpoint"
import connect_telegram from "./connect_telegram"
import create_user from "./create"
import get_user from "./get"
import list_users from "./list"
import list_permissions from "./list_permissions"
import get_user_profile from "./profile"
import set_enabled from "./set_enabled"
import set_password from "./set_password"
import update_user from "./update"

export const user = {
   connect_telegram: connect_telegram,
   create: create_user,
   list: list_users,
   read: get_user,
   read_profile: get_user_profile,
   list_permissions: list_permissions,
   update: update_user,
   set_enabled: set_enabled,
   set_password: set_password,
} satisfies ApiCollection
