import type Either from "@/api/core/either"
import { api_delete } from "@/api/core/fetch"

export default async function delete_permission_group(
   pid: string,
): Promise<Either<string, boolean>> {
   return api_delete("/api/iam/permission-group/" + pid, 204)
}
