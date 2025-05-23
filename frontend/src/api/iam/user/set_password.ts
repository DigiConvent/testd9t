import type Either from "@/api/core/either"
import { api_post } from "@/api/core/fetch"

export default function set_password(
   id: string,
   password: string,
): Promise<Either<string, boolean>> {
   return api_post<boolean>(
      "/api/iam/user/" + id + "/set-password",
      { password: password },
      undefined,
      204,
   )
}
