import { api_get } from "@/api/core/fetch"
import type { ApiCall } from "@/api/core/endpoint"
import type { UserStatusRead } from "./types"

const list_user_status: ApiCall<UserStatusRead[]> = () => {
   return api_get<UserStatusRead[]>("/api/iam/user-status", (data: any) => {
      const result: UserStatusRead[] = []
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

export default list_user_status
