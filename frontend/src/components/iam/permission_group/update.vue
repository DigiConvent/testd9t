<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form
         v-else-if="pg"
         v-permission="'iam.permission_group.write'"
         class="flex flex-col gap-4"
         @submit="handle_submit"
      >
         <FormTextInput
            v-if="pg.meta == null"
            v-model="pg.name"
            label="iam.pg.fields"
            name="name"
         />
         <FormTextInput
            v-if="pg.meta == null"
            v-model="pg.abbr"
            label="iam.pg.fields"
            name="abbr"
         />
         <FormTextareaInput v-model="pg.description" label="iam.pg.fields" name="description" />
         <PermissionGroupPicker
            v-if="pg.parent != null"
            v-model="pg.parent"
            label="iam.pg.fields"
            name="parent"
            :discriminate_descendants="id()"
            :discriminate_meta="['role', 'status']"
         ></PermissionGroupPicker>
         <PermissionPicker
            v-model="permissions"
            :multiple="true"
            :preselected="inherited_permissions"
         ></PermissionPicker>
         <Button @click="handle_submit">{{
            $t("actions.save", { entity: $t("iam.pg.pg") })
         }}</Button>
      </Form>
      <div v-else>Could not load permission group</div>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead, PermissionGroupWrite } from "@/api/iam/permission_group/types"
import { ref, watch } from "vue"

import FormTextInput from "@/components/form/text_input.vue"
import FormTextareaInput from "@/components/form/textarea.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"
import PermissionPicker from "@/components/iam/permission/picker.vue"
import { error, success } from "@/composables/toast"
import type { IdOrData } from "@/components/form/form"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const loading = ref(true)

const props = defineProps<IdOrData<PermissionGroupRead>>()
const id = () => {
   return props.id || props.data!.id
}

const pg = ref<PermissionGroupWrite | null>(null)
const parent = ref<string>("")
const permissions = ref<string[]>([])

const handle_submit = async () => {
   if (pg.value == null) return
   pg.value.permissions = permissions.value
   ;(await api.iam.permission_group.update(id(), pg.value!)).fold(
      (err: string) => {
         error(t("feedback.-.update", { entity: t("iam.pg.pg") }), err)
      },
      (result: boolean) => {
         if (!result) error("Failed to update permission group")
         else success(t("feedback.+.update", { entity: t("iam.pg.pg") }))
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
         pg.value = {
            ...permission_group,
            permissions: permission_group.permissions.map((e) => e.name),
         }
         permissions.value = permission_group.permissions
            .filter((e) => !e.implied)
            .map((e) => e.name)
         inherited_permissions.value = permission_group.permissions
            .filter((e) => e.implied)
            .map((e) => e.name)

         parent.value = permission_group.parent!
         loading.value = false
      },
   )
}

load_permission_group()

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

watch(
   pg,
   (d_old) => {
      if (d_old!.parent != parent.value) {
         load_inherited_permissions()
         parent.value = d_old!.parent!
      }
   },
   { deep: true },
)
</script>
