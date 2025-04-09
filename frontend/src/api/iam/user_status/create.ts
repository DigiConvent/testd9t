import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"
import type { UserStatusCreate } from "./types"

const create_user_status: ApiCall<string> = (user_status_create: UserStatusCreate) => {
   return api_post<string>("/api/iam/user-status", user_status_create, (data: any) => {
      return data.id
   })
}

export default create_user_status
