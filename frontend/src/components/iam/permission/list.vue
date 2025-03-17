<template>
   <div v-if="node" class="card">
      <PermissionOption
         v-for="child of node.children.sort((a, b) => a.label.localeCompare(b.label))"
         :key="child.key"
         :node="child"
         :multiple="multiple"
         :parent_hovered="false"
         :summarised="false"
      />
   </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch, getCurrentInstance } from "vue"
import { api } from "@/api"
import {
   to_permission_tree,
   type CustomTreeNode,
   type PermissionFacade,
} from "@/api/iam/permission/types"
import type Either from "@/api/core/either"
import PermissionOption from "./permission_option.vue"
import { error } from "@/composables/toast"

const node = ref<CustomTreeNode>()

defineProps<{ multiple: boolean }>()

onMounted(() => {
   api.iam.permission.list().then((result: Either<string, PermissionFacade[]>) => {
      result.fold(
         (error_message: string) => {
            error(error_message)
         },
         (permissions: PermissionFacade[]) => {
            const tree = to_permission_tree(permissions)
            node.value = tree.to_tree_node(null)
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

      if (props.onPicked) {
         if (sorted != undefined && props.picked && !props.multiple && sorted.length > 0) {
            emit("picked", sorted[0])
         } else {
            emit("picked", sorted)
         }
      }
   },
   { deep: true },
)

const emit = defineEmits(["update:modelValue", "picked", "selected"])
</script>
