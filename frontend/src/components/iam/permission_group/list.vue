<template>
   <OrganizationChart v-if="data" :value="data" collapsible>
      <template #default="slotProps">
         <span>{{ slotProps.node.data.name }}</span>
         <Button
            icon="pi pi-pencil"
            class="p-button-rounded p-button-text"
            @click="emit('click', slotProps.node.data.id)"
         ></Button>
      </template>
   </OrganizationChart>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { ref } from "vue"

const emit = defineEmits(["click"])

const data = ref<CustomNode<PermissionGroupFacade>>()
async function load_permission_groups() {
   ;(await api.iam.permission_group.list()).fold(
      (error_message: string) => {
         error(error_message)
      },
      (permission_groups: PermissionGroupFacade[]) => {
         const root = permission_groups.find((entry) => entry.parent == null)
         if (!root) return
         const root_node = create_tree_using_parent<PermissionGroupFacade>(root, permission_groups)

         data.value = root_node
      },
   )
}

load_permission_groups()
</script>
