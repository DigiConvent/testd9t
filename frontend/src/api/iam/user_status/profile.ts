import { api_get } from "@/api/core/fetch"
import type { UserStatusProfile } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"
import { to_user_status_profile } from "./utils"

const get_user_status_profile: ApiGetById<UserStatusProfile> = (pid: string) => {
   return api_get<UserStatusProfile>(
      "/api/iam/user-status/" + pid + "/profile",
      to_user_status_profile,
   )
}

export default get_user_status_profile
