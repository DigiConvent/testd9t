import type { PermissionFacade } from "../permission/types"
import type { UserFacade } from "../user/types"

export type PermissionGroupWrite = {
   name: string
   abbr: string
   description: string
   parent?: string
   is_group: boolean
   is_node: boolean
   permissions: string[]
}

export type PermissionGroupProfile = {
   permission_group: PermissionGroupRead
   permissions: PermissionFacade[]
   members: UserFacade[]
   permission_groups: PermissionGroupFacade[]
}

export type PermissionGroupFacade = {
   id: string
   name: string
   abbr: string
   is_group: boolean
   is_node: boolean
   implied: boolean
   parent: string
   generated: boolean
}

export type PermissionGroupRead = {
   id: string
   name: string
   abbr: string
   description: string
   parent: string
   is_group: boolean
   is_node: boolean
   generated: boolean
   permissions: string[]
}
