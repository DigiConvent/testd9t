import { api_post } from "@/api/core/fetch"
import type { ApiCall } from "@/api/core/endpoint"
import type { UserRoleWrite } from "./types"

const update_user_role: ApiCall<boolean> = (pid: string, user_role_write: UserRoleWrite) => {
   return api_post<boolean>("/api/iam/user-role/" + pid, user_role_write, undefined, 204)
}

export default update_user_role
