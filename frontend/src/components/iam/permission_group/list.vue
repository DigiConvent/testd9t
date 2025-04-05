<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <OrganizationChart
         v-else-if="data"
         v-permission="'iam.permission_group.list'"
         :value="data"
         collapsible
      >
         <template #default="slotProps">
            <span
               ><Fa :icon="get_icon(slotProps.node.data)" class="mr-2" />{{
                  slotProps.node.data.name
               }}</span
            >
            <Button
               class="p-button-rounded p-button-text"
               @click="emit('click', { event: $event, pg: slotProps.node.data })"
            >
               <Fa icon="fa-ellipsis-v"></Fa>
            </Button>
         </template>
      </OrganizationChart>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { ref, watch } from "vue"

const props = defineProps<{ refresh: number }>()
const emit = defineEmits(["click"])

const loading = ref(true)
const data = ref<CustomNode<PermissionGroupFacade>>()
async function load_permission_groups() {
   loading.value = true
   ;(await api.iam.permission_group.list()).fold(
      (error_message: string) => {
         error(error_message)
      },
      (permission_groups: PermissionGroupFacade[]) => {
         const root = permission_groups.find((entry) => entry.parent == null)
         if (!root) return
         const root_node = create_tree_using_parent<PermissionGroupFacade>(root, permission_groups)

         data.value = root_node
         loading.value = false
      },
   )
}

load_permission_groups()

watch(
   () => props.refresh,
   () => load_permission_groups(),
)

function get_icon(d: PermissionGroupFacade): string {
   if (d.is_group) return "users"
   if (d.is_node) return "user-tag"
   return "folders"
}
</script>
