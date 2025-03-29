<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <NeedsPermission v-else-if="data" permission="iam.permission_group.list">
         <FormTextInput
            v-model="pg_name"
            readonly
            label="iam.pg.fields"
            name="parent"
            @click="show_picker_dialog = true"
         />
         <Dialog v-model:visible="show_picker_dialog" modal :header="$t('iam.pg.fields.parent')">
            <OrganizationChart
               v-if="data"
               :value="data"
               collapsible
               :selection-mode="'single'"
               @node-select="handle_picked($event)"
            >
               <template #default="slotProps">
                  <span :class="{ 'text-white': !slotProps.node.selectable }">
                     <Fa
                        v-if="selected != null && slotProps.node.data.id == selected!.id"
                        icon="circle-check"
                     />
                     {{ slotProps.node.data.name }}</span
                  >
               </template>
            </OrganizationChart>
         </Dialog>
      </NeedsPermission>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { computed, ref } from "vue"

const t = useI18n().t
const loading = ref(true)
const show_picker_dialog = ref(false)

import FormTextInput from "@/components/form/text_input.vue"
import { error } from "@/composables/toast"
import { useI18n } from "vue-i18n"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: string | undefined
   label: string
   only_nodes?: boolean
   only_groups?: boolean
   discriminate_descendants?: string
}>()

const emit = defineEmits(["update:modelValue"])

const data = ref<CustomNode<PermissionGroupFacade & { selectable: boolean }>>()

const selected = ref<PermissionGroupFacade | null>(null)
const pg_name = computed(() => {
   if (selected.value) return selected.value.name
   return t("iam.pg.picker.none")
})

async function load_permission_groups() {
   ;(await api.iam.permission_group.list()).fold(
      (err: string) => {
         error(err)
         loading.value = false
      },
      (permission_groups: PermissionGroupFacade[]) => {
         if (props.modelValue) {
            for (const pg of permission_groups) {
               if (pg.id == props.modelValue) {
                  selected.value = pg
                  break
               }
            }
         }

         const selectable_permission_groups = permission_groups.map((pg) => {
            return {
               ...pg,
               selectable: selected.value == null || pg.id != selected.value.id,
               styleClass: "",
            }
         })

         const root = selectable_permission_groups.find((entry) => entry.parent == null)
         if (!root) return

         const root_node = create_tree_using_parent<
            PermissionGroupFacade & { selectable: boolean; styleClass: string }
         >(root, selectable_permission_groups)

         if (props.discriminate_descendants != undefined) {
            discriminate_descendant(root_node)
         }

         data.value = root_node

         loading.value = false
      },
   )
}

load_permission_groups()

function handle_picked(event: any) {
   if (event.selectable) {
      selected.value = event.data
      emit("update:modelValue", event.key)
      show_picker_dialog.value = false
   }
}

function discriminate_descendant(
   permission_group: CustomNode<
      PermissionGroupFacade & { selectable: boolean; styleClass: string }
   >,
   discriminate: boolean = false,
) {
   if (permission_group.key == props.discriminate_descendants) discriminate = true

   permission_group.selectable = !discriminate && permission_group.data.selectable

   if (!permission_group.selectable) {
      permission_group.styleClass = "!bg-gray-300"
   }
   if (permission_group.data.id == props.discriminate_descendants) {
      permission_group.styleClass = "!bg-sky-500"
   }
   if (selected.value != null && permission_group.data.id == selected.value.id) {
      permission_group.styleClass = "!bg-emerald-500"
   }
   for (const child of permission_group.children) {
      discriminate_descendant(child, discriminate)
   }
}
</script>
