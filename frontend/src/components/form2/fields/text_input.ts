import { type VNode, h } from "vue"
import { BaseField } from "../base"

export class TextInputField extends BaseField {
   render_input(): VNode {
      return h("input", {
         type: this.options.type || "text",
         value: this.value.value,
         onInput: (e: Event) => {
            this.value.value = (e.target as HTMLInputElement).value
         },
         disabled: this.options.disabled,
         placeholder: this.options.placeholder,
         class: "border py-1 px-3 rounded shadow-md",
      })
   }
}
