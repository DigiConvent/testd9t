import { type Ref, ref, type VNode, h } from "vue"

export abstract class BaseField implements Field {
   public value: Ref<any> = ref("")
   public error: Ref<string | null> = ref(null)
   public v_node: Ref<VNode> = ref(h("div"))

   constructor(
      public key: string,
      public label: string,
      public options: Record<string, any> = {},
   ) {
      this.trigger_render()
   }

   update_value(new_value: any): void {
      this.value.value = new_value
      this.trigger_render()
   }

   abstract render_field(): VNode

   trigger_render(): void {
      this.v_node.value = h("div", { class: "form-field" }, [
         h("div", "hallo"),
         this.render_label(),
         // this.render_field(),
      ])
   }

   private render_label(): VNode {
      return h("label", this.label)
   }

   private render_error(): VNode | null {
      return this.error.value ? h("div", { class: "field-error" }, this.error.value) : null
   }
}
