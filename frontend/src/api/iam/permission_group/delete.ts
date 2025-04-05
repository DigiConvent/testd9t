import type { ApiCall } from "@/api/core/endpoint"
import { api_delete } from "@/api/core/fetch"

const delete_permission_group: ApiCall<boolean> = (pid: string) => {
   return api_delete("/api/iam/permission-group/" + pid, 204)
}

export default delete_permission_group
