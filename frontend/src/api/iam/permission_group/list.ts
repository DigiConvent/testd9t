import { api_get } from "@/api/core/fetch"
import { type PermissionGroupFacade } from "./types"
import { to_permission_group_facade } from "./utils"
import type { ApiCall } from "@/api/core/endpoint"

const list_permission_groups: ApiCall<PermissionGroupFacade[]> = () => {
   return api_get<PermissionGroupFacade[]>("/api/iam/permission-group", (data: any) => {
      const result: PermissionGroupFacade[] = []
      for (const entry of data) result.push(to_permission_group_facade(entry))
      return result
   })
}

export default list_permission_groups
