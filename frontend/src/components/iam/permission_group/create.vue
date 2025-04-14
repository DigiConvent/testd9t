<template>
   <Form
      v-permission="'iam.permission_group.write'"
      class="flex flex-col gap-4"
      @submit="handle_create"
   >
      <FormTextInput v-model="pg.name" label="iam.pg.fields" name="name" />
      <FormTextInput v-model="pg.abbr" label="iam.pg.fields" name="abbr" />
      <FormTextareaInput v-model="pg.description" label="iam.pg.fields" name="description" />
      <PermissionGroupPicker
         v-model="pg.parent"
         label="iam.pg.fields"
         name="parent"
         :discriminate_meta="['role', 'status']"
      ></PermissionGroupPicker>
      <PermissionPicker
         v-model="pg.permissions"
         :multiple="true"
         :preselected="inherited_permissions"
      ></PermissionPicker>
      <Button @click="handle_create">{{ $t("actions.save") }}</Button>
   </Form>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead, PermissionGroupWrite } from "@/api/iam/permission_group/types"
import { ref, watch } from "vue"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextareaInput from "@/components/form/textarea.vue"
import PermissionGroupPicker from "./picker.vue"
import PermissionPicker from "@/components/iam/permission/picker.vue"
import { error } from "@/composables/toast"
import router from "@/router"

const props = defineProps<{ parent: string }>()
const pg = ref<PermissionGroupWrite>({
   name: "",
   abbr: "",
   description: "",
   parent: "",
   is_group: false,
   is_node: false,
   permissions: [],
})

if (props.parent != undefined) {
   pg.value.parent = props.parent
}

const handle_create = async () => {
   ;(await api.iam.permission_group.create(pg.value!)).fold(
      (err: string) => {
         error(err)
      },
      () => {
         router.back()
      },
   )
}

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
watch(
   pg,
   (new_pg, old_pg) => {
      if (new_pg.parent != old_pg.parent) load_inherited_permissions()
   },
   { deep: true },
)
</script>
