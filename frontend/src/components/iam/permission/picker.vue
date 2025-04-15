<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="node">
         <div class="border border-gray-200 p-2" @click="show_permissions_dialog = true">
            <label class="block"> {{ $t("iam.pg.fields.permissions") }} </label>
            <div class="rounded gap-2 flex">
               <div
                  v-for="permission of modelValue"
                  :key="permission"
                  @click="remove($event, permission)"
               >
                  <InputGroup class="cursor-pointer select-none text-white">
                     <Badge><Fa icon="times" />{{ permission }}</Badge>
                  </InputGroup>
               </div>
               <Badge :severity="'info'" class="cursor-pointer select-none inline">
                  <Fa icon="plus" class="fa-fw" />
                  {{ $t("actions.add", { entity: $t("iam.p.p") }) }}
               </Badge>
            </div>
         </div>
         <Dialog
            v-model:visible="show_permissions_dialog"
            modal
            @hide="show_permissions_dialog = false"
         >
            <template #default>
               <PermissionOption
                  v-for="child of node.children.sort((a, b) => a.label.localeCompare(b.label))"
                  :key="child.key"
                  :node="child"
                  :multiple="multiple"
                  :parent_hovered="false"
                  :preselected="preselected"
                  :summarised="false"
                  :readonly="(preselected || []).includes(child.key)"
               />
            </template>
            <template #footer>
               <Button @click="show_permissions_dialog = false">{{ $t("actions.close") }} </Button>
            </template>
         </Dialog>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { ref, watch, getCurrentInstance } from "vue"
import { api } from "@/api"
import {
   to_permission_tree,
   type CustomTreeNode,
   type PermissionFacade,
} from "@/api/iam/permission/types"
import PermissionOption from "./permission_option.vue"
import { error } from "@/composables/toast"

const node = ref<CustomTreeNode>()
const show_permissions_dialog = ref(false)
const loading = ref(true)
// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ multiple: boolean; modelValue: string[]; preselected: string[] }>()
async function load() {
   ;(await api.iam.permission.list()).fold(
      (error_message: string) => {
         error(error_message)
         loading.value = false
      },
      (permissions: PermissionFacade[]) => {
         const tree = to_permission_tree(permissions)
         node.value = tree.to_tree_node(null)
         for (const c of props.modelValue.values()) {
            node.value.set_checked([c])
         }

         loading.value = false
      },
   )
}
load()

const instance = getCurrentInstance()
watch(
   node,
   () => {
      // eslint-disable-next-line @typescript-eslint/no-non-null-asserted-optional-chain
      const props = instance?.vnode.props!
      const data = node.value?.get_checked()
      let sorted = data?.sort((a: string, b: string) => a.localeCompare(b))
      if (sorted != null && sorted.length == 1) {
         if (sorted[0] == "") {
            sorted = node.value?.children.map((child) => child.key)
         }
      }

      if (sorted != undefined && props.picked && !props.multiple && sorted.length > 0) {
         handle_picked([sorted[0]])
      } else {
         handle_picked(sorted || [])
      }
   },
   { deep: true },
)

function remove(event: MouseEvent, permission: string) {
   event.stopPropagation()
   if (node.value == null) return
   node.value.uncheck([permission])
   emit(
      "update:modelValue",
      props.modelValue.filter((p) => p != permission),
   )
}
const emit = defineEmits(["update:modelValue"])

function handle_picked(permissions: string[]) {
   emit("update:modelValue", permissions)
}
</script>
