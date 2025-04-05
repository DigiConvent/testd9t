import { type VNode, h } from "vue"
import { BaseField } from "./base"

export class SelectField extends BaseField {
   render_field(): VNode {
      return h(
         "select",
         {
            value: this.value.value,
            onChange: (e: Event) => {
               this.value.value = (e.target as HTMLSelectElement).value
            },
            ...this.options,
         },
         this.options.options?.map((option: any) =>
            h("option", { value: option.value }, option.label),
         ),
      )
   }
}
