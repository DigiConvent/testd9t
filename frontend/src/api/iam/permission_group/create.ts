import { api_post } from "@/api/core/fetch"
import type { PermissionGroupWrite } from "./types"
import type { ApiCall } from "@/api/core/endpoint"

const create_permission_group: ApiCall<string> = (permission_group_write: PermissionGroupWrite) => {
   return api_post<string>(
      "/api/iam/permission-group",
      permission_group_write,
      (data: any) => {
         return data.id
      },
      201,
   )
}

export default create_permission_group
