import { api_get } from "@/api/core/fetch"
import type { PermissionFacade } from "./types"
import type { ApiCall } from "@/api/core/endpoint"

const list_permissions: ApiCall<PermissionFacade[]> = async () => {
   return api_get<PermissionFacade[]>("/api/iam/permission", (data: any) => {
      return data.map((entry: any) => {
         return {
            name: entry.name,
            implied: entry.implied,
         }
      })
   })
}

export default list_permissions
