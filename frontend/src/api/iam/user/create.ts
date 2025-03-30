import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"
import type { UserCreate } from "./types"

export default async function create_user(
   user_create: UserCreate,
): Promise<Either<string, string>> {
   return api_post<string>("/api/iam/user", user_create, (data) => {
      return data.id
   })
}
