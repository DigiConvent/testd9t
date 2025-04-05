import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

const set_permission_group_parent: ApiCall<boolean> = (pid: string, parent: string) => {
   return api_post<boolean>("/api/iam/permission-group/" + pid, { parent: parent }, undefined, 204)
}

export default set_permission_group_parent
