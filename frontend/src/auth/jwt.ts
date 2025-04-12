import { api } from "@/api"
import get_web_app from "./telegram"
import type Either from "@/api/core/either"
import { ref, type Ref } from "vue"
import type { PermissionFacade } from "@/api/iam/permission/types"
import { info, warn } from "@/composables/toast"

export default class JwtAuthenticator {
   get_permissions() {
      return this.permissions
   }
   private static instance: JwtAuthenticator | undefined
   private _token: string | undefined
   public is_authenticated: Ref<boolean> = ref<boolean>(false)
   private permissions = ref<string[]>([])
   public countdown = ref<number>(0)
   private timeout: NodeJS.Timeout | undefined

   constructor() {}

   public get token() {
      if (this._token == undefined) {
         console.warn("Token is undefined, the user cannot authenticate for requests")
      }
      return this._token == undefined ? "" : this._token
   }

   recover_session(): boolean {
      const t = localStorage.getItem("token")
      if (t != null) {
         this._token = t
         if (this.seconds_remaining() < 0) {
            this.logout()
            return false
         }
         this.refresh_token()
         return true
      }
      this.is_authenticated.value = false
      return false
   }

   private seconds_remaining(): number {
      const token = this.get_token()
      if (token == null) {
         return 0
      }
      // subtract 5 seconds just to be sure
      const expiration = token.exp - Math.floor(new Date().getTime() / 1000) - 5
      return expiration
   }

   private refresh_token() {
      if (this.seconds_remaining() < 0) {
         this.logout()
         return
      }

      this.timeout = setTimeout(async () => {
         const result = await this.login(api.iam.jwt.refresh())
         if (result) {
            info("Token refreshed", "")
         } else {
            warn("Could not refresh token", "")
         }
      }, this.seconds_remaining() * 1000)
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
         enabled: boolean
      }
   } | null {
      if (this._token != undefined) {
         try {
            return JSON.parse(atob(this._token.split(".")[1]))
         } catch {
            this.logout()
         }
      }
      return null
   }

   public has_permission(permission: Exclude<string, "admin">): boolean {
      if (this.permissions.value.includes("admin")) {
         return true
      }
      // maybe there's a permission that has the prefix of permission
      for (const p of this.permissions.value) {
         if (permission.startsWith(p)) {
            return true
         }
      }
      return false
   }

   has_permissions(permissions: string[]) {
      for (const p of permissions) {
         if (this.has_permission(p)) {
            return true
         }
      }
      return false
   }

   public async load_permissions() {
      const result = await api.iam.user.list_permissions()
      result.fold(
         () => {},
         (permissions: PermissionFacade[]) => {
            this.permissions.value = permissions.map((p) => p.name)
            this.is_authenticated.value = true
         },
      )
   }

   static get_instance(): JwtAuthenticator {
      if (this.instance == undefined) {
         this.instance = new JwtAuthenticator()
      }
      return this.instance
   }

   async login_using_telegram(): Promise<string> {
      const data = get_web_app().initData
      return this.login(api.iam.login.telegram(data))
   }

   async login_using_credentials(emailaddress: string, password: string): Promise<string> {
      return this.login(api.iam.login.credentials(emailaddress, password))
   }

   async connect_telegram_user(): Promise<boolean> {
      const result = await api.iam.user.connect_telegram(get_web_app().initData)
      return result.is_right()
   }

   async login(response: Promise<Either<string, string>>): Promise<string> {
      const result = await response
      if (result.is_right()) {
         const token = result.get_right()
         if (token == undefined) {
            return "token is undefined"
         } else {
            this._token = token
            localStorage.setItem("token", token)
            this.refresh_token()
            await this.load_permissions()
            return ""
         }
      }
      return result.get_left()!
   }

   async login_using_token() {
      this.login(api.iam.jwt.refresh())
   }

   logout() {
      clearTimeout(this.timeout)
      this._token = undefined
      localStorage.removeItem("token")
      this.is_authenticated.value = false
   }
}
