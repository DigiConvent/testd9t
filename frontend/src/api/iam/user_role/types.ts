export type UserRoleRead = {
   id: string
   name: string
   abbr: string
   description: string
   archived: boolean
}

export type UserRoleWrite = {
   name: string
   abbr: string
   description: string
   archived: boolean
}

export type UserRoleCreate = {
   name: string
   abbr: string
   description: string
   archived: boolean
   parent: string
}
