<template>
   <Toast />
   <div v-if="node" class="card">
      <PermissionOption
         v-for="child of node.children"
         :key="child.key"
         :node="child"
         :parent_hovered="false"
         :multiple="multiple"
         :picker="true"
         @selected="console.log($event)"
      />
   </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch, getCurrentInstance } from "vue"
import { useToast } from "primevue/usetoast"
import { api } from "@/api"
import {
   to_permission_tree,
   type CustomTreeNode,
   type PermissionFacade,
} from "@/api/iam/permission/types"
import type Either from "@/api/core/either"
import PermissionOption from "./permission_option.vue"

const node = ref<CustomTreeNode>()
const toast = useToast()

defineProps<{ multiple: boolean }>()

onMounted(() => {
   api.iam.permission.list().then((result: Either<string, PermissionFacade[]>) => {
      result.fold(
         (error: string) => {
            toast.add({ severity: "error", summary: "Error", detail: error, life: 3000 })
         },
         (permissions: PermissionFacade[]) => {
            const tree = to_permission_tree(permissions)
            node.value = tree.to_tree_node()
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
      if (props.onPicked) {
         if (data != undefined && props.picked && props.multiple && data.length > 0) {
            emit("picked", data[0])
         } else {
            emit("picked", data)
         }
      }

      if (instance?.vnode.props) if (data == undefined) return
   },
   { deep: true },
)

const emit = defineEmits(["update:modelValue", "picked", "selected"])
</script>
