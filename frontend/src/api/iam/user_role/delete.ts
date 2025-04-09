import type { ApiCall } from "@/api/core/endpoint"
import { api_delete } from "@/api/core/fetch"

const delete_user_role: ApiCall<boolean> = (pid: string) => {
   return api_delete("/api/iam/user-role/" + pid, 204)
}

export default delete_user_role
