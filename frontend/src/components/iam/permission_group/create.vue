<template>
   <Form
      v-permission="'iam.permission_group.create'"
      class="flex flex-col gap-4"
      @submit="handle_submit"
   >
      <FormTextInput v-model="pg.name" label="iam.pg.fields" name="name" />
      <FormTextInput v-model="pg.abbr" label="iam.pg.fields" name="abbr" />
      <FormTextareaInput v-model="pg.description" label="iam.pg.fields" name="description" />
      <PermissionGroupPicker
         v-model="pg.parent"
         label="iam.pg.fields"
         name="parent"
      ></PermissionGroupPicker>
      <PermissionPicker
         v-model="pg.permissions"
         :multiple="true"
         :preselected="inherited_permissions"
      ></PermissionPicker>
      <Button @click="handle_submit">{{ $t("actions.save") }}</Button>
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

const emit = defineEmits(["created"])

const props = defineProps<{ parent: string }>()
const pg = ref<PermissionGroupWrite>({
   name: "",
   abbr: "",
   description: "",
   parent: props.parent,
   is_group: false,
   is_node: false,
   permissions: [],
})

const handle_submit = async () => {
   ;(await api.iam.permission_group.create(pg.value!)).fold(
      (err: string) => {
         error(err)
      },
      (id: string) => {
         emit("created", id)
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
