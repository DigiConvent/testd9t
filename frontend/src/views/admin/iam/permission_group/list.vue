<template>
   <div>
      PermissionGroup List
      <OrganizationChart v-if="data" :value="data" collapsible>
         <template #default="slotProps">
            <span>{{ slotProps.node.data.name }}</span>
            <Button
               icon="pi pi-pencil"
               class="p-button-rounded p-button-text"
               @click="inspect_permission_group(slotProps.node.data.id)"
            ></Button>
         </template>
      </OrganizationChart>
      <Dialog
         v-model:visible="visible"
         :dismissable-mask="true"
         pt:root:class="!border-0 !bg-transparent"
         pt:mask:class="backdrop-blur-sm"
      >
         <template #header> {{ focussed_permission_group }} </template>
         <template #container="{ closeCallback }">
            <div
               class="flex flex-col px-8 py-8 gap-6 rounded-2xl"
               style="
                  background-image: radial-gradient(
                     circle at left top,
                     var(--p-primary-400),
                     var(--p-primary-700)
                  );
               "
            >
               {{ focussed_permission_group }}
               <PermissionGroupUpdate v-model="focussed_permission_group"></PermissionGroupUpdate>
               <div class="flex items-center gap-4">
                  <Button
                     label="Cancel"
                     text
                     class="!p-4 w-full !text-primary-50 !border !border-white/30 hover:!bg-white/10"
                     @click="closeCallback"
                  ></Button>
                  <Button
                     label="Sign-In"
                     class="!p-4 w-full !text-primary-50 !border !border-white/30 hover:!bg-white/10"
                     text
                     @click="closeCallback"
                  ></Button>
               </div>
            </div>
         </template>
      </Dialog>
      <PermissionPicker @picked="handle_picked" :multiple="true"></PermissionPicker>
      <!-- <PermissionOrgChart></PermissionOrgChart> -->
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { useToast } from "primevue"
import { ref } from "vue"
import PermissionGroupUpdate from "@/components/iam/permission_group/update.vue"
import PermissionPicker from "@/components/iam/permission/list.vue"

const data = ref<CustomNode<PermissionGroupFacade>>()
const visible = ref(false)
const toast = useToast()

const handle_picked = function (data: any) {
   console.log("Picked")
   console.dir(data)
}

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

const focussed_permission_group = ref<string>("")
const inspect_permission_group = async (id: string) => {
   focussed_permission_group.value = id
   visible.value = true
}
</script>
