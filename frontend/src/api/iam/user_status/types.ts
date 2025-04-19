import type { UserFacade } from "../user/types"

export type UserStatusRead = {
   id: string
   name: string
   abbr: string
   description: string
   archived: boolean
   parent: string
}

export type UserStatusWrite = {
   id: string
   name: string
   abbr: string
   description: string
   archived: boolean
   parent: string
}

export type UserStatusCreate = {
   name: string
   abbr: string
   description: string
   archived: boolean
   parent: string
}

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
