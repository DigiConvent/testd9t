import { type Ref, ref, type VNode, h } from "vue"
import type { Field, FieldOptions } from "./types"

export abstract class BaseField implements Field {
   public value: Ref<any> = ref("")
   public error: Ref<string | null> = ref(null)
   public v_node: Ref<VNode> = ref(h("div"))

   constructor(
      public key: string,
      public label: string,
      public options: FieldOptions,
   ) {
      this.trigger_render()
   }

   update_value(new_value: any): void {
      this.value.value = new_value
      this.trigger_render()
   }

   validate(): boolean {
      if (this.options.required && !this.value.value) {
         this.error.value = "This field is required"
         return false
      }
      this.error.value = null
      return true
   }

   abstract render_input(): VNode

   protected trigger_render(): void {
      const input_wrapper = h("div", {}, [this.render_input()])
      this.v_node.value = h("div", {}, [input_wrapper, h("label", this.label)])
   }
}
