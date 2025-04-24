import type { UserRoleProfile, UserRoleRead } from "./types"

export const to_user_role_read = (data: any): UserRoleRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      archived: data.archived,
      description: data.description,
      generated: data.generated,
      meta: data.meta,
      parent: data.parent,
      permissions: data.permissions,
   }
}

export const to_user_role_profile = (data: any): UserRoleProfile => {
   return {
      role: to_user_role_read(data.user_role),
      history: data.history.map((e: any) => ({
         user: e.user,
         start: new Date(e.start),
         end: new Date(e.end),
      })),
   }
}
