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
   console.log(data.users_became_role)
   return {
      role: to_user_role_read(data.user_role),
      users_became_role: data.users_became_role.map((e: any) => ({
         user: e.user,
         start: new Date(e.start),
         end: new Date(e.end),
      })),
   }
}
