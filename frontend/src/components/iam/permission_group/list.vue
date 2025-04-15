<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <OrganizationChart
         v-else-if="data"
         v-permission="'iam.permission_group.read'"
         :value="data"
         collapsible
         @update:collapsed-keys="toggle_collapse($event)"
      >
         <template #toggleicon="slotProps">
            <Fa :icon="slotProps.expanded ? 'minus' : 'plus'" />
         </template>
         <template #default="slotProps">
            <span
               ><Fa :icon="get_icon(slotProps.node.data.meta)" class="mr-2" />{{
                  slotProps.node.data.name
               }}</span
            >
            <Button
               class="p-button-rounded p-button-text"
               @click="emit('click', { event: $event, pg: slotProps.node.data })"
            >
               <Fa icon="fa-ellipsis-v"></Fa>
            </Button>
            <div v-if="collapsed.includes(slotProps.node.key)" class="">
               <InputGroup class="text-sm text-gray-500">
                  <Tag rounded severity="secondary"
                     >{{ inventory(slotProps.node, "meta", ["role", "status"]).other }}
                     <Fa icon="folders" />
                  </Tag>
                  <Tag rounded severity="secondary"
                     >{{ inventory(slotProps.node, "meta", ["role", "status"]).role }}
                     <Fa icon="user-shield" />
                  </Tag>
                  <Tag rounded severity="secondary"
                     >{{ inventory(slotProps.node, "meta", ["role", "status"]).status }}
                     <Fa icon="user-tag" />
                  </Tag>
               </InputGroup>
            </div>
         </template>
      </OrganizationChart>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, inventory, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { ref, watch } from "vue"
import { get_icon } from "@/api/iam/permission_group/utils"

const collapsed = ref<string[]>([])
function toggle_collapse(event: any) {
   collapsed.value = Object.keys(event)
}

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
</script>
