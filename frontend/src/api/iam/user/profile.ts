import { api_get } from "@/api/core/fetch"
import type { UserProfile } from "./types"
import type { ApiGetById } from "@/api/core/endpoint"

const get_user_profile: ApiGetById<UserProfile> = (id?: string) => {
   return api_get<UserProfile>(`/api/iam/user${id ? "/" + id : "/me"}/profile`, (data: any) => {
      return {
         user: {
            id: data.user.id,
            emailaddress: data.user.emailaddress,
            first_name: data.user.first_name,
            last_name: data.user.last_name,
            enabled: data.user.enabled,
         },
         status: data.status,
         groups: data.groups,
         permissions: data.permissions,
      }
   })
}

export default get_user_profile
