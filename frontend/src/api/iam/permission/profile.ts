import { api_get } from "@/api/core/fetch"
import type { PermissionProfile } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"
import to_permission_profile from "./utils"

const get_permission_profile: ApiGetById<PermissionProfile> = async (id: string) => {
   return api_get<PermissionProfile>(`/api/iam/permission/${id}/profile`, to_permission_profile)
}

export default get_permission_profile
