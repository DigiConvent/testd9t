<template>
   <div v-if="!loading">
      <FormTextInput
         v-model="selected!.name"
         readonly
         label="iam.pg.fields"
         name="parent"
         @click="show_permission_group_dialog = true"
      />
      <Dialog
         v-model:visible="show_permission_group_dialog"
         modal
         :header="$t('iam.pg.fields.parent')"
      >
         <OrganizationChart
            v-if="data"
            :value="data"
            collapsible
            :selection-mode="'single'"
            @node-select="handle_picked($event)"
         >
            <template #default="slotProps">
               <i v-if="slotProps.node.data.id == selected!.id" class="pi pi-check-circle mr-1"></i>
               {{ slotProps.node.selectable }}
               <span
                  :class="
                     !slotProps.node.selectable || slotProps.node.data.id == selected!.id
                        ? 'text-gray-300'
                        : ''
                  "
               >
                  {{ slotProps.node.data.name }}</span
               >
            </template>
         </OrganizationChart>
      </Dialog>
   </div>
   <ProgressBar v-else mode="indeterminate"></ProgressBar>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { useToast } from "primevue"
import { ref } from "vue"

const loading = ref(true)
const show_permission_group_dialog = ref(false)

import FormTextInput from "@/components/form/text_input.vue"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: string | undefined
   label: string
   only_nodes?: boolean
   only_groups?: boolean
   discriminate_descendants?: string
}>()

const emit = defineEmits(["picked"])

const toast = useToast()

const data = ref<CustomNode<PermissionGroupFacade & { selectable: boolean }>>()

async function load_permission_groups() {
   ;(await api.iam.permission_group.list()).fold(
      (error: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: error,
         })
         loading.value = false
      },
      (permission_groups: PermissionGroupFacade[]) => {
         const selectable_permission_groups = permission_groups.map((pg) => {
            return {
               ...pg,
               selectable: true,
            }
         })

         const root = selectable_permission_groups.find((entry) => entry.parent == null)
         if (!root) return

         const root_node = create_tree_using_parent<
            PermissionGroupFacade & { selectable: boolean }
         >(root, selectable_permission_groups)

         if (props.discriminate_descendants != undefined) {
            discriminate_descendant(root_node)
         }

         data.value = root_node

         if (props.modelValue) {
            for (const pg of permission_groups) {
               if (pg.id == props.modelValue) {
                  selected.value = pg
                  break
               }
            }
         }

         loading.value = false
      },
   )
}

load_permission_groups()

const selected = ref<PermissionGroupFacade>({ name: "" } as any)
function handle_picked(event: any) {
   if (event.selectable) {
      selected.value = event.data
      emit("picked", event.key)
      show_permission_group_dialog.value = false
   }
}

function discriminate_descendant(
   permission_group: CustomNode<PermissionGroupFacade & { selectable: boolean }>,
   discriminate: boolean = false,
) {
   if (permission_group.key == props.discriminate_descendants) {
      discriminate = true
   }
   permission_group.selectable = !discriminate
   permission_group.styleClass = "bg-red-500"
   for (const child of permission_group.children) {
      discriminate_descendant(child, discriminate)
   }
}
</script>
