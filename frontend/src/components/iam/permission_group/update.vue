<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <Form
      v-else-if="pg != null"
      v-permission="'iam.permission_group.update'"
      class="flex flex-col gap-4"
      @submit="handle_submit"
   >
      <FormTextInput v-model="pg.name" label="iam.pg.fields" name="name" />
      <FormTextInput v-model="pg.abbr" label="iam.pg.fields" name="abbr" />
      <FormTextareaInput v-model="pg.description" label="iam.pg.fields" name="description" />
      <PermissionGroupPicker
         v-if="pg.parent != null"
         v-model="pg.parent"
         label="iam.pg.fields"
         name="parent"
         :discriminate_descendants="id()"
      ></PermissionGroupPicker>
      <PermissionPicker
         v-model="permissions"
         :multiple="true"
         :preselected="inherited_permissions"
      ></PermissionPicker>
      <Button @click="handle_submit">{{ $t("actions.save") }}</Button>
   </Form>
   <div v-else>Could not load permission group</div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead } from "@/api/iam/permission_group/types"
import { ref, watch } from "vue"

import FormTextInput from "@/components/form/text_input.vue"
import FormTextareaInput from "@/components/form/textarea.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"
import PermissionPicker from "@/components/iam/permission/picker.vue"
import { error } from "@/composables/toast"
import type { IdOrData } from "@/components/form/form"

const loading = ref(true)

const props = defineProps<IdOrData<PermissionGroupRead>>()
const id = () => {
   return props.id || props.data!.id
}

const pg = ref<PermissionGroupRead | null>(null)
const permissions = ref<string[]>([])
const emit = defineEmits(["updated"])

const handle_submit = async () => {
   ;(await api.iam.permission_group.update(id(), pg.value!)).fold(
      (err: string) => {
         error(err)
      },
      (result: boolean) => {
         if (!result) error("Failed to update permission group")
         else emit("updated", true)
      },
   )
}

const load_permission_group = async () => {
   loading.value = true
   ;(await api.iam.permission_group.get(id())).fold(
      (err: string) => {
         error(err)
      },
      (permission_group: PermissionGroupRead) => {
         pg.value = permission_group
         permissions.value = permission_group.permissions.map((e) => e.name)
         loading.value = false
      },
   )
}

load_permission_group()
watch(props, () => {
   load_permission_group()
})

const inherited_permissions = ref<string[]>([])
function load_inherited_permissions() {
   if (pg.value == null || pg.value.parent == null) return
   api.iam.permission_group.get(pg.value.parent).then((result) => {
      result.fold(
         (err: string) => {
            error(err)
         },
         (pg_parent: PermissionGroupRead) => {
            inherited_permissions.value = pg_parent.permissions.map((e) => e.name)
         },
      )
   })
}

load_inherited_permissions()
watch(pg, () => {
   load_inherited_permissions()
})
</script>
