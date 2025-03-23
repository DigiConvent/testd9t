<template>
   <div>
      <OrganizationChart v-if="data" :value="data" collapsible>
         <template #default="slotProps">
            <i
               v-if="
                  slotProps.node.data.id == modelValue && slotProps.node.data.parent == undefined
               "
               class="pi pi-check-circle"
            ></i>
            <span>{{ slotProps.node.data.name }}</span>
            <Button
               class="p-button-rounded p-button-text"
               @click="emit('picked', slotProps.node.data.id)"
            ></Button>
         </template>
      </OrganizationChart>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { useToast } from "primevue"
import { ref } from "vue"

defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: string | undefined
   only_nodes?: boolean
   only_groups?: boolean
}>()

const emit = defineEmits(["picked"])

const toast = useToast()

const data = ref<CustomNode<PermissionGroupFacade>>()

async function load_permission_groups() {
   ;(await api.iam.permission_group.list()).fold(
      (error: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: error,
         })
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
