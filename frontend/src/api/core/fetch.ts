import router from "@/router"
import Either from "./either"
import JwtAuthenticator from "@/auth/jwt"

interface FromJSON<T> {
   (data: any): T
}

export async function api_get<T>(
   url: string,
   format_data: FromJSON<T>,
): Promise<Either<string, T>> {
   const result = new Either<string, T>()
   const request = await fetch(url, {
      method: "GET",
      mode: "same-origin",
      headers: {
         Authorization: "" + (localStorage.getItem("token") || ""),
         page: router.currentRoute.value.fullPath,
      },
   })

   if (request.status == 401) {
      JwtAuthenticator.get_instance().logout()
   }

   if (request.ok) {
      let data: any = {}
      try {
         data = await request.json()
      } catch (e: any) {
         return result.left("Could not read json from request " + e)
      }
      const formatted_data = format_data(data)
      return result.right(formatted_data)
   } else {
      return result.left(request.status + ": ")
   }
}
export async function api_post<T>(
   url: string,
   payload: any,
   format_data?: FromJSON<T>,
   expects?: number,
): Promise<Either<string, T>> {
   const result = new Either<string, T>()
   const body = JSON.stringify(payload)

   const request = await fetch(url, {
      method: "POST",
      headers: {
         "Content-Type": "application/json",
         Authorization: localStorage.getItem("token") || "",
         page: router.currentRoute.value.fullPath,
      },
      body: body,
   })

   if (expects != undefined) {
      if (expects != request.status) {
         return result.left(request.status + ": ")
      } else {
         return result.right(true as T)
      }
   }

   let data: any = {}
   if (request.status >= 200 && request.status < 300) {
      try {
         data = await request.json()
      } catch (e: any) {
         return result.left("Could not read json from request " + e)
      }
   }

   if (request.status == 401) {
      JwtAuthenticator.get_instance().logout()
   }

   if (request.ok) {
      if (expects && expects == request.status) {
         return result.right(true as T)
      } else if (format_data != undefined) {
         const formatted_data = format_data(data)
         return result.right(formatted_data)
      } else {
         return result.left("Malfunction")
      }
   } else {
      return result.left(request.status + ": " + data["error"])
   }
}

export async function api_multipart<T>(
   url: string,
   body: any,
   files: Map<string, File>,
   expects?: number,
): Promise<Either<string, T>> {
   const form_data = new FormData()
   for (const key in body) {
      form_data.set(key, body[key])
   }

   for (const key of files.keys()) {
      form_data.set(key, files.get(key)!)
   }

   const result = new Either<string, T>()
   const request = await fetch(url, {
      method: "POST",
      headers: {
         Authorization: "" + (localStorage.getItem("token") || ""),
         page: router.currentRoute.value.fullPath,
      },
      body: form_data,
   })

   const data: any = await request.json()
   if (expects != undefined) {
      if (expects == request.status) return result.right(true as T)
      else return result.left(data["error"])
   }

   if (request.ok) {
      return result.right(data)
   } else {
      return result.left(request.status + ": " + data["message"])
   }
}

export async function api_delete(url: string, expects?: number): Promise<Either<string, boolean>> {
   const result = new Either<string, boolean>()
   const request = await fetch(url, {
      method: "DELETE",
      headers: {
         Authorization: "" + (localStorage.getItem("token") || ""),
      },
   })

   if (expects && expects == request.status) {
      return result.right(true)
   } else {
      return result.left(request.status + ": ")
   }
}
