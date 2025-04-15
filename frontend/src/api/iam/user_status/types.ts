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
