import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"

export default async function add_user(pid: string, uid: string): Promise<Either<string, boolean>> {
   return api_post<boolean>("/api/iam/permission-group/" + pid, { uid: uid }, undefined, 204)
}
