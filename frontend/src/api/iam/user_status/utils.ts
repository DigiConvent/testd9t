import type { UserStatusRead } from "./types"

export const to_user_status_read = (data: any): UserStatusRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      description: data.description,
      archived: data.archived,
   }
}
