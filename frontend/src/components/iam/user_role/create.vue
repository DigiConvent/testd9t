<template>
   <div class="card flex justify-center">
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form v-else class="flex flex-col gap-4">
         <FormTextInput v-model="user_role.name" label="iam.ur.fields" name="name" />
         <FormTextInput v-model="user_role.abbr" label="iam.ur.fields" name="abbr" />
         <FormTextarea v-model="user_role.description" label="iam.ur.fields" name="description" />
         <PermissionGroupPicker v-model="user_role.parent" label="iam.ur.fields" name="parent" />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('actions.create', { entity: $t('iam.us.us') })"
               @click="create_user_role"
            ></Button>
         </div>
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { api } from "@/api"
import { error } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextarea from "@/components/form/textarea.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"
import type { UserRoleCreate } from "@/api/iam/user_role/types"
import router from "@/router"

const loading = ref(false)
const user_role = ref<UserRoleCreate>({
   name: "",
   abbr: "",
   description: "",
   parent: "",
})

const emit = defineEmits(["created"])

async function create_user_role() {
   ;(await api.iam.user_role.create(user_role.value)).fold(
      (l: string) => {
         error(l)
      },
      (id: string) => {
         router.push({ name: "admin.iam.user_role.profile", params: { id: id } })
      },
   )
}

const props = defineProps<{ parent?: string }>()
function load() {
   loading.value = true
   if (props.parent != undefined) {
      user_role.value.parent = props.parent
   }
   loading.value = false
}

load()
</script>
