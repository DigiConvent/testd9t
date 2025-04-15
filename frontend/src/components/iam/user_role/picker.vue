<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Select
         v-else-if="user_status"
         v-model="value_facade"
         class="w-full"
         :options="user_status"
         option-label="name"
         option-value="id"
         :empty-message="$t('iam.us.picker.empty')"
      >
         <template #footer>
            <div class="p-3">
               <Button
                  :label="$t('iam.us.create.title')"
                  fluid
                  severity="secondary"
                  text
                  @click="show_new_user_status_form = true"
               />
            </div>
         </template>
      </Select>
      <Dialog
         v-model:visible="show_new_user_status_form"
         modal
         :header="$t('iam.us.create.title')"
         :style="{ width: '25rem' }"
      >
         <NewUserStatusForm @created="load_user_status($event)"></NewUserStatusForm>
      </Dialog>
   </div>
</template>

<script lang="ts" setup>
import { ref, defineProps, computed } from "vue"
import { api } from "@/api"
import NewUserStatusForm from "./create.vue"
import { error } from "@/composables/toast"
import type { UserStatusRead } from "@/api/iam/user_status/types"

const loading = ref(true)
const user_status = ref<UserStatusRead[]>([])
const show_new_user_status_form = ref(false)

// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ modelValue: string }>()
const emit = defineEmits(["update:modelValue"])
const value_facade = computed({
   get: () => props.modelValue,
   set: (value: string) => emit("update:modelValue", value),
})

async function load_user_status(selected: string = "") {
   show_new_user_status_form.value = false
   loading.value = true
   const result = await api.iam.user_status.list()
   result.fold(
      (err: string) => {
         error(err)
      },
      (result: UserStatusRead[]) => {
         user_status.value = result
         loading.value = false
         if (selected != "") {
            emit("update:modelValue", selected)
         }
      },
   )
}

load_user_status()
</script>
