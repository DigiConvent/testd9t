<template>
   <div class="card flex justify-center">
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form v-else class="flex flex-col gap-4">
         <FormTextInput v-model="user_status.name" label="iam.ur.fields" name="name" />
         <FormTextInput v-model="user_status.abbr" label="iam.ur.fields" name="abbr" />
         <FormTextarea v-model="user_status.description" label="iam.ur.fields" name="description" />
         <PermissionGroupPicker
            v-model="user_status.parent"
            label="iam.ur.fields"
            :discriminate_meta="['status', 'role']"
            name="parent"
         />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('actions.create', { entity: $t('iam.us.us') })"
               @click="create_user_status"
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
import PermissionGroupPicker from "@/components/pkg/iam/permission_group/picker.vue"
import type { UserStatusCreate } from "@/api/iam/user_status/types"
import router from "@/router"

const loading = ref(false)
const user_status = ref<UserStatusCreate>({
   name: "",
   abbr: "",
   description: "",
   archived: false,
   parent: "",
})

async function create_user_status() {
   ;(await api.iam.user_status.create(user_status.value)).fold(
      (l: string) => {
         error(l)
      },
      (id: string) => {
         router.push({ name: "admin.iam.user_status.profile", params: { id: id } })
      },
   )
}

const props = defineProps<{ parent?: string }>()
function load() {
   loading.value = true
   if (props.parent != undefined) {
      user_status.value.parent = props.parent
   }
   loading.value = false
}

load()
</script>
