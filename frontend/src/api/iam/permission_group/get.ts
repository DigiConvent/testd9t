import { api_get } from "@/api/core/fetch"
import type { PermissionGroupRead } from "./types"
import { to_permission_group_read } from "./utils"
import type { ApiGetById } from "@/api/core/endpoint"

const get_permission_group: ApiGetById<PermissionGroupRead> = (pid: string) => {
   return api_get<PermissionGroupRead>("/api/iam/permission-group/" + pid, to_permission_group_read)
}

export default get_permission_group
