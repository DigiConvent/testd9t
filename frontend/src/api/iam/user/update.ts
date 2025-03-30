import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"
import type { UserUpdate } from "./types"

export default async function update_user(
   user_update_data: UserUpdate,
   id?: string,
): Promise<Either<string, string>> {
   return api_post<string>(`/api/iam/user${id ? "/" + id : "/me"}`, user_update_data, (data) => {
      return data.id
   })
}
