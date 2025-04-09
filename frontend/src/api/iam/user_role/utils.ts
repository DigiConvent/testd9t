import type { UserRoleRead } from "./types"

export const to_user_role_read = (data: any): UserRoleRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      description: data.description,
      archived: data.archived,
   }
}
