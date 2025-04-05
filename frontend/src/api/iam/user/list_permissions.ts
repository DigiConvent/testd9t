import { api_get } from "@/api/core/fetch"
import type { PermissionFacade } from "../permission/types"
import type { ApiCall } from "@/api/core/endpoint"

const list_user_permissions: ApiCall<PermissionFacade[]> = () => {
   return api_get<PermissionFacade[]>("/api/iam/user/me/permissions", (data: any) => {
      const result: PermissionFacade[] = []
      for (const entry of data) {
         result.push({
            name: entry.name,
            implied: entry.implied,
         })
      }
      return result
   })
}

export default list_user_permissions
