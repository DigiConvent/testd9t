import type Either from "../../core/either"
import { api_post } from "../../core/fetch"

export default async function refresh(): Promise<Either<string, string>> {
   return api_post<string>("/api/iam/jwt/refresh", {  }, (data) => {
      return data.token
   })
}
