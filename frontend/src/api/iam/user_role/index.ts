import type { ApiCollection } from "@/api/core/endpoint"
import create_user_role from "./create"
import list_user_role from "./list"
import delete_user_role from "./delete"
import update_user_role from "./update"
import get_user_role from "./read"
import get_user_role_profile from "./profile"

export const user_role = {
   create: create_user_role,
   list: list_user_role,
   read: get_user_role,
   read_profile: get_user_role_profile,
   update: update_user_role,
   delete: delete_user_role,
} satisfies ApiCollection
