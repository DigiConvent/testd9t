import { api_get } from "@/api/core/fetch"
import type { UserRoleRead } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"
import { to_user_role_read } from "./utils"

const get_user_role: ApiGetById<UserRoleRead> = (pid: string) => {
   return api_get<UserRoleRead>("/api/iam/user-role/" + pid, to_user_role_read)
}

export default get_user_role
