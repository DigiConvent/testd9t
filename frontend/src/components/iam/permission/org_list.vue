<template>
   <div>
      <OrganizationChart v-if="data" :value="data" collapsible>
         <template #default="slotProps">
            <span>{{ slotProps.node.data.name }}</span>
            <Button
               icon="pi pi-pencil"
               class="p-button-rounded p-button-text"
               @click="console.log(slotProps.node.data.id)"
            ></Button>
         </template>
      </OrganizationChart>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { type CustomNode } from "@/api/core/node"
import { to_permission_tree, type PermissionFacade } from "@/api/iam/permission/types"
import { useToast } from "primevue"
import { ref } from "vue"

const data = ref<CustomNode<PermissionFacade>>()

const toast = useToast()
api.iam.permission.list().then((result) =>
   result.fold(
      (err: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: err,
         })
      },
      (permissions: PermissionFacade[]) => {
         const permission_tree = to_permission_tree(permissions)
         data.value = permission_tree
      },
   ),
)
</script>
