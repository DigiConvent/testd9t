import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

export type UserStatusCreate = {
   name: string
   abbr: string
   description: string
   archived: boolean
   parent: string
}

const create_user_status: ApiCall<string> = (user_status_create: UserStatusCreate) => {
   return api_post<string>("/api/iam/user-status", user_status_create, (data: any) => {
      return data.id
   })
}

export default create_user_status
