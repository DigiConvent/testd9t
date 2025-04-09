import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"
import type { UserRoleCreate } from "./types"

const create_user_role: ApiCall<string> = (user_role_create: UserRoleCreate) => {
   return api_post<string>("/api/iam/user-role", user_role_create, (data: any) => {
      return data.id
   })
}

export default create_user_role
