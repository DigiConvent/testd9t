import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import { type PermissionGroupFacade } from "./types"
import { to_permission_group_facade } from "./utils"

export default async function list_permission_groups(): Promise<
   Either<string, PermissionGroupFacade[]>
> {
   return api_get<PermissionGroupFacade[]>("/api/iam/permission-group", (data: any) => {
      const result: PermissionGroupFacade[] = []
      for (const entry of data) result.push(to_permission_group_facade(entry))
      return result
   })
}
