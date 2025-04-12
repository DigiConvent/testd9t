import type { UserRoleProfile, UserRoleRead } from "./types"

export const to_user_role_read = (data: any): UserRoleRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      description: data.description,
      archived: data.archived,
   }
}

export const to_user_role_profile = (data: any): UserRoleProfile => {
   return {
      role: to_user_role_read(data.role),
      users_became_role: [],
   }
}
