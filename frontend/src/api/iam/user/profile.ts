import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import type { UserProfile } from "./types"

export default async function get_user_profile(id?: string): Promise<Either<string, UserProfile>> {
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
