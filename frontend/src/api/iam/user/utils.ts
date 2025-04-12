import type { UserFacade } from "./types"

export function to_user_facade(data: any): UserFacade {
   return {
      id: data.id,
      first_name: data.first_name,
      last_name: data.last_name,
      status_id: data.status_id,
      status_name: data.status_name,
      roles: data.roles,
      implied: data.implied,
   }
}
