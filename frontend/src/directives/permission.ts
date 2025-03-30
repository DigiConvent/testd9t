import JwtAuthenticator from "@/auth/jwt"

import { type Directive } from "vue"

const auth = JwtAuthenticator.get_instance()
export const permission: Directive = {
   beforeMount(el, binding) {
      const { modifiers, value } = binding
      if (typeof value == "string" && modifiers.except === undefined && !auth.has_permission(value))
         el.style.display = "none"
      else if (typeof value == "boolean" && value) el.style.display = null
   },
}
