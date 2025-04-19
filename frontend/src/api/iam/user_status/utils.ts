import type { UserStatusProfile, UserStatusRead } from "./types"

export const to_user_status_read = (data: any): UserStatusRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      description: data.description,
      archived: data.archived,
      parent: data.parent,
   }
}

export const to_user_status_profile = (data: any): UserStatusProfile => {
   return {
      user_status: data.user_status,
      users_became_status: data.users_became_status.map((e: any) => ({
         user: e.user,
         status: e.status,
         start: new Date(e.start),
         end: new Date(e.end),
      })),
   }
}
