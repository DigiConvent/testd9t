import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"
import type { PermissionGroupWrite } from "./types"

export default async function create_permission_group(
   permission_group_write: PermissionGroupWrite,
): Promise<Either<string, string>> {
   return api_post<string>(
      "/api/iam/permission-group",
      permission_group_write,
      (data: any) => {
         return data.id
      },
      201,
   )
}
