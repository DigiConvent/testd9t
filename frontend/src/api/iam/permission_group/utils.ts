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
      is_group: data.is_group,
      is_node: data.is_node,
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
      members: data.members.map((entry: any) => {
         return to_user_facade(entry)
      }),
      permission_groups: data.permission_groups.map((entry: any) => {
         return {
            id: entry.id,
            name: entry.name,
            abbr: entry.abbr,
            description: entry.description,
            parent: entry.parent,
            is_group: entry.is_group,
            is_node: entry.is_node,
            generated: entry.generated,
         }
      }),
   }
}

export function to_permission_group_facade(data: any): PermissionGroupFacade {
   return {
      id: data.id,
      name: data.name,
      abbr: data.abbr,
      is_group: data.is_group,
      is_node: data.is_node,
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
      is_group: pg.is_group,
      is_node: pg.is_node,
      permissions: pg.permissions,
   }
}
