import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "../../core/fetch"

const refresh: ApiCall<string> = async () => {
   return api_post<string>("/api/iam/jwt/refresh", {}, (data) => {
      return data.token
   })
}

export default refresh
