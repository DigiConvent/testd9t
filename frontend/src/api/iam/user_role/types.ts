import type { UserFacade } from "../user/types"

export type UserRoleRead = {
   id: string
   name: string
   abbr: string
   description: string
   archived: boolean
}

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
   id?: string
   name: string
   abbr: string
   description: string
   archived: boolean
}

export type UserRoleCreate = {
   name: string
   abbr: string
   description: string
   parent: string
}
