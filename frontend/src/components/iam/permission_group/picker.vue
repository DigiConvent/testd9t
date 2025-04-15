<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="data" v-permission="'iam.permission_group.read'">
         <FormTextInput
            v-model="pg_name"
            readonly
            :label="label"
            :name="name"
            class="!cursor-pointer"
            @click="show_picker_dialog = true"
         />
         <Dialog v-model:visible="show_picker_dialog" modal :header="$t(label + '.' + name)">
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
                     <Fa :icon="get_icon(slotProps.node.data.meta)" />
                     {{ slotProps.node.data.name }}</span
                  >
               </template>
            </OrganizationChart>
         </Dialog>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { create_tree_using_parent, type CustomNode } from "@/api/core/node"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { computed, ref } from "vue"
import { get_icon } from "@/api/iam/permission_group/utils"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: string | undefined
   label: string
   name: string
   discriminate_meta?: (null | "role" | "status")[]
   discriminate_descendants?: string
}>()

const t = useI18n().t
const loading = ref(true)
const show_picker_dialog = ref(false)

import FormTextInput from "@/components/form/text_input.vue"
import { error } from "@/composables/toast"
import { useI18n } from "vue-i18n"

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

         discriminate_permission_groups(root_node)

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

function discriminate_permission_groups(
   permission_group: CustomNode<
      PermissionGroupFacade & { selectable: boolean; styleClass: string }
   >,
   discriminate: boolean = false,
) {
   if (
      props.discriminate_descendants != undefined &&
      permission_group.key == props.discriminate_descendants
   ) {
      permission_group.styleClass = "!bg-sky-500"
      discriminate = true
   }
   permission_group.selectable = !discriminate && permission_group.data.selectable

   if (
      props.discriminate_meta != undefined &&
      props.discriminate_meta.includes(permission_group.data.meta)
   ) {
      permission_group.selectable = false
   }

   if (!permission_group.selectable) {
      permission_group.styleClass = "!bg-gray-300"
   }

   if (selected.value != null && permission_group.data.id == selected.value.id) {
      permission_group.styleClass = "!bg-emerald-500"
   }
   for (const child of permission_group.children) {
      discriminate_permission_groups(child, discriminate)
   }
}
</script>
