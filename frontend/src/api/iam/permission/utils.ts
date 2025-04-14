import type { PermissionProfile } from "./types"

export default function to_permission_profile(data: any): PermissionProfile {
   return {
      permission: data.permission,
      permission_groups: data.permission_groups,
      users: data.users,
      descendants: data.descendants,
   }
}
