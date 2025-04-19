import type { ApiCollection } from "@/api/core/endpoint"
import create_user_status from "./create"
import list_user_status from "./list"
import get_user_status from "./read"
import update_user_status from "./update"
import delete_user_status from "./delete"
import get_user_status_profile from "./profile"

export const user_status = {
   create: create_user_status,
   list: list_user_status,
   get: get_user_status,
   profile: get_user_status_profile,
   update: update_user_status,
   delete: delete_user_status,
} satisfies ApiCollection
