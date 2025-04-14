import JwtAuthenticator from "@/auth/jwt"
import router from "@/router"

import { type Directive } from "vue"

const auth = JwtAuthenticator.get_instance()
export const permission: Directive = {
   beforeMount(el, binding) {
      const { modifiers, value } = binding
      if (typeof value == "string" && modifiers.except === undefined) {
         if (!auth.has_permission(value)) {
            el.style.display = "none"
         }
      } else if (typeof value == "boolean" && value) el.style.display = null
   },
   mounted(el, binding) {
      if (window.debug)
         if (Object.keys(binding.modifiers).length == 0) {
            const wrapper = document.createElement(binding.arg || "fieldset")

            if (el.className) {
               wrapper.className = el.className
            }

            if (el.attrs) {
               Object.entries(el.attrs).forEach(([key, value]: any) => {
                  wrapper.setAttribute(key, value)
               })
            }

            const parent = el.parentNode
            parent.insertBefore(wrapper, el)
            wrapper.appendChild(el)
            const legend = document.createElement("legend")

            const { modifiers, value } = binding
            if (typeof value == "string" && modifiers.except === undefined) {
               legend.innerText = value
               legend.classList.add("!cursor-pointer")
               legend.classList.add("!select-none")
               legend.classList.add("!border")
               legend.classList.add("!text-xs")
               legend.ondblclick = () => {
                  router.push({ name: "admin.iam.permission.profile", params: { name: value } })
               }
            }
            wrapper.appendChild(legend)

            el._wrapDirectiveWrapper = wrapper
         }
   },
}
