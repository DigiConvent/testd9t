import { api_post } from "@/api/core/fetch"
import type { ApiCall } from "@/api/core/endpoint"
import type { UserStatusWrite } from "./types"

const update_user_status: ApiCall<boolean> = (pid: string, user_status_write: UserStatusWrite) => {
   return api_post<boolean>("/api/iam/user-status/" + pid, user_status_write, undefined, 204)
}

export default update_user_status
