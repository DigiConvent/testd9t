import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import type { UserRead } from "./types"

export default async function get_user(id: string): Promise<Either<string, UserRead>> {
   return api_get<UserRead>("/api/iam/user/" + id, (data: any) => {
      return {
         id: data.id,
         emailaddress: data.emailaddress,
         first_name: data.first_name,
         last_name: data.last_name,
         date_of_birth: data.date_of_birth,
      }
   })
}
