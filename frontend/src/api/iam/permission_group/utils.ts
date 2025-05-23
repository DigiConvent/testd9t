import { to_user_facade } from "../user/utils"
import type {
   PermissionGroupFacade,
   PermissionGroupProfile,
   PermissionGroupRead,
   PermissionGroupWrite,
} from "./types"

export const to_permission_group_read = (data: any): PermissionGroupRead => {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      description: data.description,
      parent: data.parent,
      meta: data.meta,
      generated: data.generated,
      permissions: data.permissions,
   }
}
export function to_permission_group_profile(data: any): PermissionGroupProfile {
   return {
      permission_group: to_permission_group_read(data.permission_group),
      permissions: data.permissions.map((entry: any) => {
         return {
            name: entry.name,
            implied: entry.implied,
            description: entry.description,
         }
      }),
      users: data.users.map((entry: any) => {
         return to_user_facade(entry)
      }),
      ancestors: data.ancestors.map((entry: any) => to_permission_group_facade(entry)),
      descendants: data.descendants.map((entry: any) => to_permission_group_facade(entry)),
   }
}

export function to_permission_group_facade(data: any): PermissionGroupFacade {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      meta: data.meta,
      implied: data.implied,
      parent: data.parent,
      generated: data.generated,
   }
}

export function to_permission_group_write(pg: PermissionGroupRead): PermissionGroupWrite {
   return {
      name: pg.name,
      abbr: pg.abbr,
      description: pg.description,
      parent: pg.parent,
      meta: pg.meta,
      permissions: pg.permissions.map((e) => e.name),
   }
}

export function get_icon(type: string | null): string {
   if (type == null || type == "") return "folders"
   if (type == "role") return "user-shield"
   if (type == "status") return "user-tag"
   return "folders"
}
