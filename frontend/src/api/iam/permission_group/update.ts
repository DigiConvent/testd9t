import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"
import type { PermissionGroupWrite } from "./types"

export default async function update_permission_group(
   pid: string,
   permission_group_write: PermissionGroupWrite,
): Promise<Either<string, boolean>> {
   return api_post<boolean>(
      "/api/iam/permission-group/" + pid,
      permission_group_write,
      undefined,
      204,
   )
}
