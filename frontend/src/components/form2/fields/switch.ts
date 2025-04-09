import { h, type VNode } from "vue"
import { BaseField } from "../base"

export class SwitchField extends BaseField {
   render_input(): VNode {
      return h(
         "label",
         {
            class: "switch",
            style: {
               "--active-color": this.options.activeColor,
               "--inactive-color": this.options.inactiveColor,
            },
         },
         [
            h("input", {
               type: "checkbox",
               checked: this.value.value,
               onChange: (e: Event) => {
                  this.value.value = (e.target as HTMLInputElement).checked
               },
               disabled: this.options.disabled,
            }),
            h("span", { class: "slider round" }),
         ],
      )
   }
}
