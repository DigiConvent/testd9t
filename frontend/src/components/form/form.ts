import { h, ref, type Ref, type VNode } from "vue"
import type { ApiGetById } from "@/api/core/endpoint"
import type Either from "@/api/core/either"
import { error, success } from "@/composables/toast"
import { Button, Fieldset, ProgressBar } from "primevue"
import type Field from "./field"

export default class Form<T> {
   private _fields: Field[] = []
   private _field_models: Map<string, Ref<any>> = new Map()
   private _initialised = false
   private _loading = false
   private _value_instance?: T
   private _form_config: FormConfig<T>

   private _vnode: Ref<VNode> = ref(h("div", "hallo"))
   public get vnode() {
      this.render()
      return this._vnode
   }

   constructor(config: FormConfig<T>) {
      this._form_config = config
      if (config.id !== undefined) {
         this._loading = true
         this._form_config.get(config.id).then((either: Either<string, T>) => {
            either.fold(
               (err: string) => {
                  this._loading = false
                  error(err)
                  this.render()
               },
               (data: T) => {
                  success("loaded")
                  this._value_instance = data
                  this._loading = false
                  this.render()
               },
            )
         })
         return
      }
      this._value_instance = config.data
      this.render()
   }

   public add_field(class_reference: FieldReference, args: FieldArgs): Form<T> {
      args.name = this._form_config.t = args.name
      const instance = new class_reference(args)
      this._fields.push(instance)
      return this
   }

   public async done() {
      // make sure that the data is loaded
      while (this._value_instance === undefined) {
         await new Promise((resolve) => {
            setTimeout(() => {
               resolve(true)
            }, 100)
         })
      }

      const keys = Object.getOwnPropertyNames(this._value_instance)
      const entries = Object.entries(this._value_instance!)

      for (const field in this._fields) {
         const key = this._fields[field].name
         if (keys.includes(key)) {
            const value = entries.find((e) => e[0] == key)![1]
            this._field_models.set(key, ref(value))
            this._fields[field].set_value(this._field_models.get(key)!)
         }
      }

      this._initialised = true
      this.render()
   }

   private render(): void {
      if (!this._initialised) {
         return
      }

      const child_nodes = this._fields.map((e) => e.vnode.value)
      const child_v_nodes = this._fields.map((e) => e.vnode)
      const v_models = Object.fromEntries(this._field_models.entries())

      this._vnode.value = h("div", [
         h(
            Fieldset,
            {
               class: {
                  "animate-pulse": this._loading,
               },
            },
            {
               legend: () =>
                  this._loading
                     ? h(ProgressBar, {
                          style: { width: "7rem", height: ".3rem" },
                          mode: "indeterminate",
                       })
                     : h(
                          "div",
                          {
                             class: "border-t border-emerald-500",
                             style: { width: "7rem" },
                          },
                          child_nodes.map((e) => h(e)),
                       ),
               default: () =>
                  h("div", [
                     h(
                        "span",
                        child_v_nodes.map((e) => h(e)),
                     ),
                     h("pre", JSON.stringify(v_models, null, 3)),
                     h(
                        Button,
                        {
                           onClick: () => {},
                        },
                        "done",
                     ),
                     h("div", `${this._field_models.size}`),
                  ]),
            },
         ),
      ])
   }

   public to_json(): any {
      return {
         value_source: this._form_config,
         fields: this._fields.map((e) => e.to_json()),
      }
   }
}

export type IdOrData<T> =
   | { id: string; data?: undefined }
   | { id?: undefined; data: T }
   | { id?: undefined; data: T & { id: string } } // sometimes the id is also present in the data
export type FormConfig<T> = IdOrData<T> & {
   get: ApiGetById<T>
   t: string
}

type FieldReference = { new (...args: any[]): any }
type FieldArgs = { label: string; name: string }
