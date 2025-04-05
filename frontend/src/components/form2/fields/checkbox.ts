import { type VNode, h } from "vue"
import { BaseField } from "../base"

export class CheckboxField extends BaseField {
   render_input(): VNode {
      return h("input", {
         type: "checkbox",
         checked: this.value.value,
         onChange: (e: Event) => {
            this.value.value = (e.target as HTMLInputElement).checked
         },
         disabled: this.options.disabled,
         class: "checkbox-input",
      })
   }
}
