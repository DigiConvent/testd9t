import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import type { PermissionGroupRead } from "./types"
import { to_permission_group_read } from "./utils"

export default async function get_permission_group(
   pid: string,
): Promise<Either<string, PermissionGroupRead>> {
   return api_get<PermissionGroupRead>(
      "/api/iam/permission-group/" + pid + "/",
      to_permission_group_read,
   )
}
