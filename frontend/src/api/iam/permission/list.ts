import type Either from "@/api/core/either"
import { api_get } from "@/api/core/fetch"
import type { PermissionFacade } from "./types"

export default async function list_permissions(): Promise<Either<string, PermissionFacade[]>> {
   return api_get<PermissionFacade[]>("/api/iam/permission", (data: any) => {
      return data.map((entry: any) => {
         return {
            name: entry.name,
            implied: entry.implied,
         }
      })
   })
}
