import { h } from "vue"
import Field from "../field"

export default class TextField extends Field {
   public render_component(): void {
      if (this._v_model == undefined) return
      this._vnode.value = h("div", [
         h("span", this._label + this._old_value),
         // h("input", {
         //    onInput: (e: InputEvent) => {
         //       if (e.target == null) return
         //       this._v_model!.value = (e.target as HTMLInputElement).value!
         //       this.render_component()
         //    },
         //    name: this._name,
         //    vModel: this._new_value,
         //    type: "text",
         //    placeholder: this._label,
         //    class: "w-full border border-red-300 p-3 block",
         // }),
         // h(
         //    Button,
         //    {
         //       default: "hallo",
         //       onClick: () => {
         //          this._v_model!.value = "hallo"
         //       },
         //    },
         //    ["awdawd"],
         // ),
         // h("a", this._v_model!.value),
         // h("pre", JSON.stringify(Object(this._v_model), null, 3)),
      ])
   }

   constructor(args: { label: string; name: string; value: string; old_value?: string }) {
      super(args)
   }

   public to_json() {}
}
