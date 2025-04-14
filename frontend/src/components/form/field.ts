import { h, ref, type Ref, type VNode } from "vue"

export default abstract class Field {
   protected _v_model?: Ref<any>
   protected _old_value: any
   protected _new_value: any
   protected _name: string
   protected _label: string
   protected _vnode: Ref<VNode> = ref(h("div"))

   public get name(): string {
      return this._name
   }

   public get vnode() {
      return this._vnode
   }

   constructor(arg: { label: string; name: string; value: any; old_value?: any }) {
      this._old_value = arg.old_value || arg.value
      this._new_value = arg.value
      this._name = arg.name
      this._label = arg.label
      this.render()
   }

   public set_value(value: Ref) {
      this._v_model = value
   }
   public abstract to_json(): any
   public abstract render_component(): void
   public render() {
      this.render_component()
   }
}
