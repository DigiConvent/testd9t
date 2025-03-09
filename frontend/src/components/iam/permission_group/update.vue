<template>
   <Toast />
   <div v-if="auth.has_permission('iam.permission_group.update')" class="card flex justify-center">
      <Form v-if="pg != null" class="flex flex-col gap-4 w-full sm:w-56" @submit="handle_submit">
         <FormTextInput v-model="pg.name" label="iam.pg.update" :error="errors.name" name="name" />
         <FormTextInput v-model="pg.abbr" label="iam.pg.update" :error="errors.abbr" name="abbr" />
         <FormTextInput
            v-model="pg.description"
            label="iam.pg.update"
            :error="errors.description"
            name="description"
         />
         <div class="flex flex-col gap-1">
            <PermissionPicker></PermissionPicker>
         </div>
         <div class="flex flex-col gap-1">
            <PermissionGroupPicker v-model="pg.parent" @picked="pg.parent = $event" />
         </div>
         <Button type="submit" severity="secondary" :label="$t('iam.auth.login_form.submit')" />
      </Form>
   </div>
   <div v-else class="card flex justify-center">
      {{ $t("iam.auth.no_permission") }}
      {{ auth.has_permission("super") }}
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead, PermissionGroupWrite } from "@/api/iam/permission_group/types"
import { to_permission_group_write } from "@/api/iam/permission_group/utils"
import JwtAuthenticator from "@/auth/jwt"
import { useToast } from "primevue"
import { ref } from "vue"
import PermissionGroupPicker from "./picker.vue"
import PermissionPicker from "./../permission/list.vue"

import FormTextInput from "@/components/form/text_input.vue"

const auth = JwtAuthenticator.get_instance()

const props = defineProps(["modelValue"])

const pg = ref<PermissionGroupWrite | null>(null)

const errors = ref<{ name: string; abbr: string; description: string }>({
   name: "",
   abbr: "",
   description: "",
})

const handle_submit = async () => {}

const toast = useToast()
const load_permission_group = async () => {
   ;(await api.iam.permission_group.get(props.modelValue)).fold(
      (error: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: error,
         })
      },
      (permission_group: PermissionGroupRead) => {
         pg.value = to_permission_group_write(permission_group)
      },
   )
}

load_permission_group()
</script>
