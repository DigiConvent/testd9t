import { api } from "@/api"
import get_web_app from "./telegram"
import type Either from "@/api/core/either"
import { ref, type Ref } from "vue"
import type { PermissionFacade } from "@/api/iam/permission/types"
import { info, warn } from "@/composables/toast"

export type Session = {
   user: {
      id: string
      first_name: string
      last_name: string
      emailaddress: string
   }
   token: string
   expires: number
}

export default class JwtAuthenticator {
   get_permissions() {
      return this.permissions
   }
   private static instance: JwtAuthenticator | undefined
   public is_authenticated: Ref<boolean> = ref<boolean>(false)
   public countdown = ref<number>(0)
   private permissions = ref<string[]>([])
   private timeout: NodeJS.Timeout | undefined

   constructor() {}

   private _token: string | undefined
   public get token() {
      if (this._token == undefined) {
         console.warn("Token is undefined, the user cannot authenticate for requests")
      }
      return this._token == undefined ? "" : this._token
   }

   public get first_name(): string {
      if (this._token == undefined) return ""
      return JSON.parse(atob(this._token.split(".")[1]))["first_name"]
   }

   public get sessions(): Session[] {
      const raw_sessions = localStorage.getItem("sessions") || "[]"
      const parsed_sessions = JSON.parse(raw_sessions)
      const sessions: Session[] = []
      for (const session_data of parsed_sessions) {
         sessions.push({
            user: session_data.user,
            token: session_data.token,
            expires: session_data.expires,
         })
      }
      return sessions
   }

   async recover_session(user_id: string = ""): Promise<boolean> {
      if (user_id == "") user_id = this.sessions[0].user.id
      const session = this.sessions.find((s) => s.user.id == user_id)
      if (session == undefined) return false
      const t = session.token
      if (t != null) {
         this._token = t
         if (this.seconds_remaining() < 0) {
            this.logout()
            return false
         }
         await this.load_permissions()
         this.refresh_token()
         return true
      }
      this.is_authenticated.value = false
      return false
   }

   private seconds_remaining(raw_token: string = ""): number {
      const token = this.get_token(raw_token)
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

   public get_token(token: string = ""): {
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
      if (token == "") {
         token = this._token!
      }
      if (token != undefined) {
         try {
            return JSON.parse(atob(token.split(".")[1]))
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
      const sessions_without_email = this.sessions.filter(
         (s) => s.user.emailaddress != emailaddress,
      )
      localStorage.setItem("sessions", JSON.stringify(sessions_without_email))
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
            this.add_session({
               user: this.get_token()!.user,
               token: token,
               expires: this.get_token()!.exp,
            })
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

   logout(emailaddress: string = ""): boolean {
      if (emailaddress == "") {
         if (this._token == undefined) {
            return false
         }

         emailaddress = this.get_token()?.user.emailaddress!
      }

      if (emailaddress == this.get_token()?.user.emailaddress!) {
         clearTimeout(this.timeout)
         this._token = undefined
         this.is_authenticated.value = false
      }

      const sessions = this.sessions
      const clean_sessions = sessions.filter((s) => s.user.emailaddress != emailaddress)
      localStorage.setItem("sessions", JSON.stringify(clean_sessions))
      return sessions.length != clean_sessions.length
   }

   add_session(session: Session) {
      const sessions = this.sessions
      for (let i = 0; i < sessions.length; i++) {
         if (sessions[i].user.emailaddress == session.user.emailaddress) {
            sessions[i] = session
            localStorage.setItem("sessions", JSON.stringify(sessions))
            return
         }
      }

      sessions.push(session)
      localStorage.setItem("sessions", JSON.stringify(sessions))
   }
}
