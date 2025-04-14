import { api_post } from "@/api/core/fetch"
import type { UserUpdate } from "./types"
import type { ApiSaveById } from "@/api/core/endpoint"

const update_user: ApiSaveById<UserUpdate, boolean> = (
   id: string,
   user_update_data: UserUpdate,
) => {
   return api_post<boolean>(
      `/api/iam/user${id ? "/" + id : "/me"}`,
      user_update_data,
      undefined,
      204,
   )
}

export default update_user
