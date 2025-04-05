import { api_get } from "@/api/core/fetch"
import type { PermissionGroupProfile } from "./types"
import { to_permission_group_profile } from "./utils"
import type { ApiGetById } from "@/api/core/endpoint"

const get_permission_group_profile: ApiGetById<PermissionGroupProfile> = (pid: string) => {
   return api_get<PermissionGroupProfile>(
      "/api/iam/permission-group/profile/" + pid,
      to_permission_group_profile,
   )
}

export default get_permission_group_profile
