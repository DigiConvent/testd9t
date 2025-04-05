import { type VNode, h } from "vue"
import { BaseField } from "./base"

export class TextInput extends BaseField {
   render_field(): VNode {
      return h("input", {
         type: this.options.type || "text",
         value: this.value.value,
         onInput: (e: Event) => {
            this.value.value = (e.target as HTMLInputElement).value
         },
         ...this.options,
      })
   }
}
