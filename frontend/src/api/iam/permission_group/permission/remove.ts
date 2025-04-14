import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

const remove_permission: ApiCall<boolean> = (pgid: string, permission: string) => {
   return api_post<boolean>(
      "/api/iam/permission-group/" + pgid + "/permission/",
      { add: "", remove: permission },
      undefined,
      204,
   )
}

export default remove_permission
