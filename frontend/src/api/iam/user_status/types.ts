import type { PermissionGroupRead, PermissionGroupWrite } from "../permission_group/types"
import type { UserFacade } from "../user/types"

export type UserStatusRead = {
   archived: boolean
} & PermissionGroupRead

export type UserStatusWrite = {
   archived: boolean
} & PermissionGroupWrite

export type UserStatusCreate = {
   archived: boolean
} & PermissionGroupWrite

export type UserBecameStatus = {
   role: string
   user: UserFacade
   start: Date
   end: Date
}

export type UserStatusProfile = {
   user_status: UserStatusRead
   users_became_status: UserBecameStatus[]
}
