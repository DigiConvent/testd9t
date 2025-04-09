import { api_get } from "@/api/core/fetch"
import type { ApiCall } from "@/api/core/endpoint"
import type { UserRoleRead } from "./types"

const list_user_role: ApiCall<UserRoleRead[]> = () => {
   return api_get<UserRoleRead[]>("/api/iam/user-role", (data: any) => {
      const result: UserRoleRead[] = []
      for (const entry of data.items || []) {
         result.push({
            id: entry.id,
            name: entry.name,
            abbr: entry.abbr,
            description: entry.description,
            archived: entry.archived,
         })
      }
      return result
   })
}

export default list_user_role
