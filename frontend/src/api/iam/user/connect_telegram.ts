import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

const connect_telegram: ApiCall<boolean> = (init_data: string) => {
   return api_post<boolean>(
      "/api/iam/login/telegram/connect",
      { payload: init_data },
      undefined,
      200,
   )
}

export default connect_telegram
