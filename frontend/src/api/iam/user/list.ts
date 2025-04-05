import { api_get } from "@/api/core/fetch"
import type { UserFacade } from "./types"
import type { Page } from "@/api/core/page"
import { to_user_facade } from "./utils"
import type { ApiCall } from "@/api/core/endpoint"

const list_users: ApiCall<Page<UserFacade>> = () => {
   return api_get<Page<UserFacade>>("/api/iam/user", (data: any) => {
      return {
         items: data.items.map((entry: any) => to_user_facade(entry)),
         page: data.page,
         items_per_page: data.items_per_page,
         total_items: data.total_items,
      }
   })
}

export default list_users
