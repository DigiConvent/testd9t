import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"

export default async function set_permission_group_parent(
   pid: string,
   parent: string,
): Promise<Either<string, boolean>> {
   return api_post<boolean>(
      "/api/iam/permission-group/" + pid + "/",
      { parent: parent },
      undefined,
      204,
   )
}
