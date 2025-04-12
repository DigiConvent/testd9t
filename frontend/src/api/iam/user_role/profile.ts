import { api_get } from "@/api/core/fetch"
import type { UserRoleProfile } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"
import { to_user_role_profile } from "./utils"

const get_user_role_profile: ApiGetById<UserRoleProfile> = (pid: string) => {
   return api_get<UserRoleProfile>("/api/iam/user-role/" + pid + "/profile", to_user_role_profile)
}

export default get_user_role_profile
