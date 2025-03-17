import { api } from "@/api"
import get_web_app from "./telegram"
import type Either from "@/api/core/either"
import { ref, type Ref } from "vue"
import type { PermissionFacade } from "@/api/iam/permission/types"
import { info, warn } from "@/composables/toast"

export default class JwtAuthenticator {
   private static instance: JwtAuthenticator | undefined
   private token: string | undefined
   public is_authenticated: Ref<boolean> = ref<boolean>(false)
   private permissions = ref<string[]>([])
   public countdown = ref<number>(0)

   constructor() {
      const t = localStorage.getItem("token")
      if (t != null) {
         this.token = t
         this.refresh_token()
         this.load_permissions().then(() => {
            this.is_authenticated.value = true
         })
      } else {
         this.is_authenticated.value = false
      }
   }

   private refresh_token() {
      const expiration = this.get_token()?.exp
      const now = Math.floor(new Date().getTime() / 1000)
      if (expiration == undefined || now > expiration - 5) {
         this.logout()
         return
      }

      const timeout = expiration! - now

      this.countdown.value = (timeout - 5) * 1000
      setTimeout(
         async () => {
            const result = await this.login(api.iam.jwt.refresh())
            if (result) {
               info("Token refreshed", "")
            } else {
               warn("Could not refresh token", "")
            }
         },
         (timeout - 5) * 1000,
      )
   }

   public get_token(): {
      id: string
      exp: number
      tgid: number
      iat: number
      user: {
         id: string
         emailaddress: string
         first_name: string
         last_name: string
         date_of_birth: string
         enabled: boolean
      }
   } | null {
      if (this.token != undefined) return JSON.parse(atob(this.token.split(".")[1]))
      return null
   }

   public has_permission(permission: Exclude<string, "super">): boolean {
      if (this.permissions.value.includes("super")) {
         return true
      }
      return this.permissions.value.includes(permission)
   }

   public async load_permissions() {
      const result = await api.iam.user.list_permissions()
      result.fold(
         (error: string) => {
            console.error(error)
         },
         (permissions: PermissionFacade[]) => {
            this.permissions.value = permissions.map((p) => p.name)
         },
      )
   }

   static get_instance(): JwtAuthenticator {
      if (this.instance == undefined) {
         this.instance = new JwtAuthenticator()
      }
      return this.instance
   }

   async login_using_telegram(): Promise<boolean> {
      const data = get_web_app().initData
      return this.login(api.iam.login.telegram(data))
   }

   async login_using_credentials(emailaddress: string, password: string): Promise<boolean> {
      return this.login(api.iam.login.credentials(emailaddress, password))
   }

   async connect_telegram_user(): Promise<boolean> {
      const result = await api.iam.user.connect_telegram(get_web_app().initData)
      return result.is_right()
   }

   async login(response: Promise<Either<string, string>>): Promise<boolean> {
      const result = await response
      if (result.is_right()) {
         const token = result.get_right()
         this.token = token
         localStorage.setItem("token", token!)
         this.refresh_token()
         await this.load_permissions()
         this.is_authenticated.value = true
         return true
      } else {
         return false
      }
   }

   async login_using_token() {
      this.login(api.iam.jwt.refresh())
   }

   logout() {
      this.token = undefined
      localStorage.removeItem("token")
      this.is_authenticated.value = false
   }
}
