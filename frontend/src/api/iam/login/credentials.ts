import Either from "@/api/core/either"
import type { ApiCall } from "@/api/core/endpoint"

const credentials: ApiCall<string> = async (emailaddress: string, password: string) => {
   const url = "/api/iam/login/credentials"
   const body = JSON.stringify({ emailaddress, password })
   const request = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      mode: "same-origin",
      body: body,
   })

   const data = await request.json()

   if (request.ok) {
      return new Either<string, string>().right(data.token)
   } else {
      return new Either<string, string>().left(request.status + ": " + data["message"])
   }
}

export default credentials
