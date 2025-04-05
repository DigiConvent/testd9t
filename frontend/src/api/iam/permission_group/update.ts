import { api_post } from "@/api/core/fetch"
import type { PermissionGroupWrite } from "./types"
import type { ApiCall } from "@/api/core/endpoint"

const update_permission_group: ApiCall<boolean> = (
   pid: string,
   permission_group_write: PermissionGroupWrite,
) => {
   return api_post<boolean>(
      "/api/iam/permission-group/" + pid,
      permission_group_write,
      undefined,
      204,
   )
}

export default update_permission_group
