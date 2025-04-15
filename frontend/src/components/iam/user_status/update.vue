<template>
   <div class="card flex justify-center">
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form v-else-if="data" class="flex flex-col gap-4">
         <FormTextInput v-model="data.name" label="iam.user_status.create" name="name" />
         <FormTextInput v-model="data.abbr" label="iam.user_status.create" name="abbr" />
         <FormTextarea
            v-model="data.description"
            label="iam.user_status.create"
            name="description"
         />
         <FormSwitch
            v-model="data.archived"
            label_on="iam.user_status.create.archived"
            label_off="iam.user_status.create.unarchived"
            name="archived"
         />
         <PermissionGroupPicker
            v-model="data.parent"
            label="iam.user_status.create"
            name="parent"
         />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('iam.us.create.submit')"
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
import FormSwitch from "@/components/form/switch.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"
import type { UserStatusRead, UserStatusWrite } from "@/api/iam/user_status/types"
import { type IdOrData } from "@/components/form/form"

const loading = ref(true)
const data = ref<UserStatusRead>()
const props = defineProps<IdOrData<UserStatusWrite> & { parent?: string }>()
const us_parent = ref<string>()

const emit = defineEmits(["updated"])

async function create_user_status() {
   ;(await api.iam.user_status.create()).fold(
      (l: string) => {
         error(l)
      },
      (id: string) => {
         emit("updated", id)
      },
   )
}

async function load() {
   loading.value = true

   if (props.id != undefined) {
      ;(await api.iam.user_status.get(props.id)).fold(
         (err: string) => {
            error(err)
         },
         (user_status: UserStatusRead) => {
            data.value = user_status
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
