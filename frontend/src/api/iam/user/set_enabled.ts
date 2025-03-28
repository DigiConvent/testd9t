import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"

export default async function set_enabled(
   id: string,
   enabled: boolean,
): Promise<Either<string, boolean>> {
   return api_post<boolean>(
      "/api/iam/user/" + id + "/enabled",
      { enabled: enabled },
      undefined,
      200,
   )
}
