<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="node">
      {{ modelValue }}
      <FormTextarea
         v-model="value_facade"
         label="iam.pg.fields"
         name="permissions"
         @click="show_permissions_dialog = true"
      ></FormTextarea>
      {{ value_facade }}
      <Dialog
         v-model:visible="show_permissions_dialog"
         modal
         @hide="show_permissions_dialog = false"
      >
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
         <template #footer>
            <Button @click="show_permissions_dialog = false">{{ $t("actions.close") }} </Button>
         </template>
      </Dialog>
   </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch, getCurrentInstance, computed } from "vue"
import { api } from "@/api"
import {
   to_permission_tree,
   type CustomTreeNode,
   type PermissionFacade,
} from "@/api/iam/permission/types"
import type Either from "@/api/core/either"
import PermissionOption from "./permission_option.vue"
import { error } from "@/composables/toast"
import FormTextarea from "@/components/form/textarea.vue"

const node = ref<CustomTreeNode>()
const show_permissions_dialog = ref(false)
const loading = ref(true)
// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ multiple: boolean; modelValue: string[]; preselected: string[] }>()

onMounted(() => {
   api.iam.permission.list().then((result: Either<string, PermissionFacade[]>) => {
      result.fold(
         (error_message: string) => {
            error(error_message)
            loading.value = false
         },
         (permissions: PermissionFacade[]) => {
            const tree = to_permission_tree(permissions)
            node.value = tree.to_tree_node(null)
            loading.value = false
         },
      )
   })
})

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

const value_facade = computed({
   get: () => {
      return (props.modelValue || []).join(", ")
   },
   set: () => {
      // do nothing
   },
})
const emit = defineEmits(["update:modelValue"])

function handle_picked(permissions: string[]) {
   emit("update:modelValue", permissions)
}
</script>
