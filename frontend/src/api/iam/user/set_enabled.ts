import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

const set_enabled: ApiCall<boolean> = (id: string, enabled: boolean) => {
   return api_post<boolean>(
      "/api/iam/user/" + id + "/enabled",
      { enabled: enabled },
      undefined,
      200,
   )
}

export default set_enabled
