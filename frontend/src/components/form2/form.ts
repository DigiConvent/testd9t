import { type Ref, ref, type VNode, h, watch } from "vue"
import { SelectField } from "./fields/select"
import { TextInputField } from "./fields/text_input"
import type { Field, FieldOptions, FormInput } from "./types"
import { CheckboxField } from "./fields/checkbox"
import { error } from "@/composables/toast"
import { Button, ProgressBar } from "primevue"

export class Form<T extends Record<string, any>> {
   private t: string = "form"
   private fields: Field[] = []
   private entity: Ref<T | null> = ref(null)
   private loading = ref(true)
   private error = ref<string | null>(null)
   public v_node: Ref<VNode> = ref(h("fieldset"))

   constructor(private input: FormInput<T>) {
      this.initialize_entity()
      watch(
         () => this.entity.value,
         () => this.update_field_values(),
         { deep: true },
      )
   }

   add_text_input(options: FieldOptions): this {
      return this.add_field(new TextInputField(`${this.t}.${options.key}`, "text", options))
   }

   add_email_input(options: FieldOptions): this {
      return this.add_field(new TextInputField(`${this.t}.${options.key}`, "email", options))
   }

   add_select(options: FieldOptions & { options: Array<{ value: any; label: string }> }): this {
      if (!this.validate_entity_key(options.key)) {
         throw new Error(`Key "${options.key}" does not exist in entity`)
      }
      return this.add_field(new SelectField(`${this.t}.${options.key}`, options.label, options))
   }

   add_checkbox(options: FieldOptions): this {
      return this.add_field(new CheckboxField(`${this.t}.${options.key}`, options.label, options))
   }

   private add_field(field: Field): this {
      this.fields.push(field)

      watch(field.value, (new_val) => {
         if (this.entity.value) {
            ;(this.entity.value as any)[field.key] = new_val
         }
      })

      this.update_field_values()
      this.trigger_render()
      return this
   }

   private validate_entity_key(key: string): boolean {
      if (!this.entity.value) return true
      return key in this.entity.value
   }

   private async initialize_entity() {
      if (this.input.id) {
         await this.load_entity(this.input.id)
      } else if (this.input.data) {
         this.entity.value = this.input.data
         this.update_field_values()
      }
      this.loading.value = false
   }

   private async load_entity(id: string) {
      this.loading.value = true
      try {
         ;(await this.input.fetch_endpoint!(id)).fold(
            (err: string) => {
               error(err)
            },
            (data: T) => {
               setTimeout(() => {
                  this.entity.value = data
                  this.update_field_values()
                  this.loading.value = false
               }, 3000)
            },
         )
         this.update_field_values()
      } catch (err: any) {
         error(err)
      } finally {
         this.trigger_render()
      }
   }

   async save_entity(): Promise<void> {
      if (!this.entity.value) return

      this.loading.value = true
      this.trigger_render()

      if (this.input.update_endpoint) {
         ;(await this.input.update_endpoint(this.input.id!, this.entity.value)).fold(
            (err: string) => {
               error(err)
            },
            (success: boolean) => {
               if (this.input.on_success && success) this.input.on_success()
            },
         )
      } else {
         ;(await this.input.save_endpoint(this.entity.value)).fold(
            (err: string) => {
               error(err)
            },
            (id: string) => {
               if (this.input.on_success && id) this.input.on_success(id)
            },
         )
      }
   }

   private update_field_values() {
      if (!this.entity.value) return

      this.fields.forEach((field) => {
         if (this.entity.value && field.key in this.entity.value) {
            field.update_value(this.entity.value[field.key])
         }
      })
      this.trigger_render()
   }

   private trigger_render() {
      this.v_node.value = this.render()
   }

   private render(): VNode {
      return h("fieldset", { class: { "!animate-pulse": this.loading.value } }, [
         h(
            "legend",
            {
               class: {
                  "!px-4": true,
               },
            },
            this.loading.value
               ? h(ProgressBar, {
                    style: { width: "7rem", height: "2px" },
                    mode: "indeterminate",
                 })
               : this.t,
         ),
         ...this.fields.map((field) => field.v_node.value),
         h(
            Button,
            {
               type: "submit",
               disabled: this.loading.value,
               onClick: () => this.save_entity(),
            },
            this.loading.value ? "Saving..." : "Save",
         ),
         h("pre", {}, [JSON.stringify(this.entity.value, null, 3)]),
      ])
   }
}
