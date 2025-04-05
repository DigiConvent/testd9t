import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "../../core/fetch"

const telegram: ApiCall<string> = (data_string: string) => {
   return api_post<string>(
      "/api/iam/login/telegram",
      { payload: data_string },
      (data) => {
         return data.token
      },
      200,
   )
}

export default telegram
