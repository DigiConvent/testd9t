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
            :label="$t('iam.user_status.new')"
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
      header="Edit Profile"
      :style="{ width: '25rem' }"
    >
      <NewUserStatusForm @created="load_user_status()"></NewUserStatusForm>
    </Dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, defineProps } from "vue"
import type { UserStatusRead } from "@/api/iam"
import { useToast } from "primevue"
import { api } from "@/api"
import NewUserStatusForm from "./create.vue"

const toast = useToast()
const user_status = ref<UserStatusRead[]>([])

defineProps(["modelValue"])
const emit = defineEmits(["update:modelValue"])

async function load_user_status() {
  const result = await api.iam.user_status.list()
  result.fold(
    (error: string) => {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: error
      })
    },
    (result: UserStatusRead[]) => {
      user_status.value = result
    }
  )
}

load_user_status()

const show_new_user_status_form = ref(false)
</script>
