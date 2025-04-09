import { api_get } from "@/api/core/fetch"
import type { UserStatusRead } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"
import { to_user_status_read } from "./utils"

const get_user_status: ApiGetById<UserStatusRead> = (pid: string) => {
   return api_get<UserStatusRead>("/api/iam/user-status/" + pid, to_user_status_read)
}

export default get_user_status
