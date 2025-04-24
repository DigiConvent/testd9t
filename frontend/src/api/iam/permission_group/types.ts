import type { PermissionFacade } from "../permission/types"
import type { UserFacade } from "../user/types"

export type PermissionGroupWrite = {
   name: string
   abbr: string
   description: string
   parent?: string
   permissions: string[]
}

export type PermissionGroupProfile = {
   permission_group: PermissionGroupRead
   permissions: PermissionFacade[]
   users: UserFacade[]
   ancestors: PermissionGroupFacade[]
   descendants: PermissionGroupFacade[]
}

export type PermissionGroupFacade = {
   id: string
   name: string
   abbr: string
   meta: "" | "role" | "status"
   implied: boolean
   parent?: string
   generated: boolean
}

export type PermissionGroupRead = {
   id: string
   name: string
   abbr: string
   description: string
   parent?: string
   meta: "" | "role" | "status"
   generated: boolean
   permissions: PermissionFacade[]
}
