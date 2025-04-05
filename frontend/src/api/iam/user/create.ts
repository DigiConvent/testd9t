import { api_post } from "@/api/core/fetch"
import type { UserCreate } from "./types"
import type { ApiCall } from "@/api/core/endpoint"

const create_user: ApiCall<string> = (user_create: UserCreate) => {
   return api_post<string>("/api/iam/user", user_create, (data) => {
      return data.id
   })
}

export default create_user
