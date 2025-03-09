<template>
   <div></div>
</template>

<script lang="ts" setup>
import {
   CustomTreeNode,
   // to_permission_tree,
   // type PermissionFacade,
} from "@/api/iam/permission/types"
// import { api } from "@/api"
// import type Either from "@/api/core/either"
import { onMounted, ref } from "vue"
// import { useToast } from "primevue/usetoast"

const props = defineProps<{ node: CustomTreeNode }>()

// const toast = useToast()

const node = ref(props.node)
onMounted(() => {
   console.log(node.value)
   // if (node.value) return
   // api.iam.permission.list().then((result: Either<string, PermissionFacade[]>) => {
   //    result.fold(
   //       (error: string) => {
   //          toast.add({ severity: "error", summary: "Error", detail: error, life: 3000 })
   //       },
   //       (permissions: PermissionFacade[]) => {
   //          const tree = to_permission_tree(permissions)
   //          node.value = tree.to_tree_node()
   //       },
   //    )
   // })
})
</script>
