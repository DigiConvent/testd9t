import { api_get } from "@/api/core/fetch"
import type { UserRead } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"

const get_user: ApiGetById<UserRead> = (id?: string) => {
   return api_get<UserRead>(`/api/iam/user${id ? "/" + id : "/me"}`, (data: any) => {
      return {
         id: data.id,
         emailaddress: data.emailaddress,
         first_name: data.first_name,
         last_name: data.last_name,
         enabled: data.enabled,
      }
   })
}

export default get_user
