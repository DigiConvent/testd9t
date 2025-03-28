<template>
   <div class="card flex justify-center">
      <Select
         :value="modelValue"
         :options="user_status"
         option-label="name"
         class="w-full"
         :empty-message="$t('iam.user_status.picker.empty')"
         @input="emit('update:modelValue', $event.target.value)"
      >
         <template #value="slotProps">
            <div v-if="slotProps.value" class="flex items-center">
               <div>{{ slotProps.value.name }}</div>
            </div>
            <span v-else>
               {{ slotProps.placeholder }}
            </span>
         </template>
         <template #option="slotProps">
            <div class="flex items-center">
               <div>{{ slotProps.option.name }}</div>
            </div>
         </template>
         <template #footer>
            <div class="p-3">
               <Button
                  :label="$t('iam.user_status.create.title')"
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
         :header="$t('iam.user_status.create.title')"
         :style="{ width: '25rem' }"
      >
         <NewUserStatusForm @created="load_user_status()"></NewUserStatusForm>
      </Dialog>
   </div>
</template>

<script lang="ts" setup>
import { ref, defineProps } from "vue"
import type { UserStatusRead } from "@/api/iam"
import { api } from "@/api"
import NewUserStatusForm from "./create.vue"
import { error } from "@/composables/toast"

const user_status = ref<UserStatusRead[]>([])

// eslint-disable-next-line vue/prop-name-casing
defineProps<{ modelValue: string }>()
const emit = defineEmits(["update:modelValue", "empty"])

async function load_user_status() {
   const result = await api.iam.user_status.list()
   result.fold(
      (err: string) => {
         error(err)
      },
      (result: UserStatusRead[]) => {
         if (result.length == 0) {
            emit("empty")
         } else {
            user_status.value = result
         }
      },
   )
}

load_user_status()

const show_new_user_status_form = ref(false)
</script>
