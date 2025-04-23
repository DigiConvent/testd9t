<template>
   <div>
      <Form class="flex flex-col gap-4">
         <FormTextInput v-model="password" label="iam.user.fields" name="password" />
         <Button @click="handle_submit">
            {{ $t("iam.user.set_password.submit") }}
         </Button>
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { error } from "@/composables/toast"
import { ref } from "vue"
import FormTextInput from "@/components/form/text_input.vue"

const props = defineProps<{ id?: string }>()
const emit = defineEmits(["success"])
const password = ref("")

async function handle_submit() {
   ;(await api.iam.user.set_password(props.id || "me", password.value)).fold(
      (err: string) => {
         console.log(err)
         error(err)
      },
      () => {
         emit("success")
      },
   )
}
</script>
