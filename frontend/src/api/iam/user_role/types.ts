import type { PermissionGroupRead, PermissionGroupWrite } from "../permission_group/types"
import type { UserFacade } from "../user/types"

export type UserRoleRead = { archived: boolean; parent: string } & PermissionGroupRead

export type UserRoleProfile = {
   role: UserRoleRead
   history: UserBecameRole[]
}

export type UserBecameRole = {
   role: string
   user: UserFacade
   start: Date
   end: Date
}

export type UserRoleWrite = {
   archived: boolean
} & PermissionGroupWrite

export type UserRoleCreate = {
   name: string
   abbr: string
   description: string
   parent: string
}
