import type { ApiCollection } from "@/api/core/endpoint"
import create_user_status from "./create"
import list_user_status from "./list"

export const user_status: ApiCollection = {
   create: create_user_status,
   list: list_user_status,
}
