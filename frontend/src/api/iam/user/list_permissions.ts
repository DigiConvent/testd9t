import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import type { PermissionFacade } from "../permission/types"

export default async function list_user_permissions(): Promise<Either<string, PermissionFacade[]>> {
   return api_get<PermissionFacade[]>("/api/iam/user/permission/", (data: any) => {
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
