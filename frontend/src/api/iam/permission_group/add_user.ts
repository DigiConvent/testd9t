import type { ApiCall } from "@/api/core/endpoint"
import { api_post } from "@/api/core/fetch"

const add_user: ApiCall<boolean> = (pid: string, uid: string) => {
   return api_post<boolean>("/api/iam/permission-group/" + pid, { uid: uid }, undefined, 204)
}

export default add_user
