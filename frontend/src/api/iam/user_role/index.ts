import type { ApiCollection } from "@/api/core/endpoint"
import create_user_role from "./create"
import list_user_role from "./list"
import delete_user_role from "./delete"
import update_user_role from "./update"
import get_user_role from "./read"

export const user_role: ApiCollection = {
   create: create_user_role,
   list: list_user_role,
   read: get_user_role,
   update: update_user_role,
   delete: delete_user_role,
}
