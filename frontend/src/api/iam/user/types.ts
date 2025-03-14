export type UserCreate = {
   emailaddress: string
   first_name: string
   last_name: string
   date_of_birth: string
}

export type UserRead = {
   id: string
   emailaddress: string
   first_name: string
   last_name: string
   date_of_birth: string
}

export type UserFacade = {
   id: string
   name: string
   implied: boolean
   status_id: string
   status_name: string
}
