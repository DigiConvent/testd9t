import type { PermissionGroupFacade } from "../permission_group/types"

export type UserCreate = {
   emailaddress: string
   first_name: string
   last_name: string
   user_status: string
   when: Date
}

export type UserUpdate = {
   emailaddress: string
   first_name: string
   last_name: string
}

export type UserRead = {
   id: string
   emailaddress: string
   first_name: string
   last_name: string
   enabled: boolean
}

export type UserFacade = {
   id: string
   first_name: string
   last_name: string
   implied: boolean
   status_id: string
   status_name: string
   role_id: string
   role_name: string
}

export type UserProfile = {
   user: UserRead
   status: null
   groups: PermissionGroupFacade
   permissions: string[]
}
