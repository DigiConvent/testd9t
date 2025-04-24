<template>
   <div class="card flex justify-center">
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form v-else-if="data" class="flex flex-col gap-4">
         <FormTextInput v-model="data.name" label="iam.user_role.fields" name="name" />
         <FormTextInput v-model="data.abbr" label="iam.user_role.fields" name="abbr" />
         <FormTextarea v-model="data.description" label="iam.user_role.fields" name="description" />
         <FormSwitch
            v-model="data.archived"
            label_on="iam.user_role.fields.archived"
            label_off="iam.user_role.fields.unarchived"
            name="archived"
         />
         <PermissionGroupPicker v-model="data.parent" label="iam.user_role.fields" name="parent" />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('iam.us.create.submit')"
               @click="update_user_role"
            ></Button>
         </div>
         <PermissionPicker v-model="data.permissions" :multiple="true" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { api } from "@/api"
import { error, success } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextarea from "@/components/form/textarea.vue"
import FormSwitch from "@/components/form/switch.vue"
import PermissionGroupPicker from "@/components/pkg/iam/permission_group/picker.vue"
import PermissionPicker from "@/components/pkg/iam/permission/picker.vue"
import { type IdOrData } from "@/components/form/form"
import type { UserRoleRead, UserRoleWrite } from "@/api/iam/user_role/types"

const loading = ref(true)
const data = ref<UserRoleWrite>()
const props = defineProps<IdOrData<UserRoleWrite> & { parent?: string }>()
const us_parent = ref<string>()

async function update_user_role() {
   ;(await api.iam.user_role.update()).fold(
      (l: string) => {
         error(l)
      },
      (result: boolean) => {
         if (result) {
            success("Successfully updated user role")
         }
      },
   )
}

async function load() {
   loading.value = true

   if (props.id != undefined) {
      ;(await api.iam.user_role.read(props.id)).fold(
         (err: string) => {
            error(err)
         },
         (user_role: UserRoleRead) => {
            data.value = user_role
            loading.value = false
         },
      )
   } else {
      data.value = props.data
   }

   if (props.parent != undefined) {
      us_parent.value = props.parent
   }
   loading.value = false
}

load()
</script>
