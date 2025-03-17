<template>
   <div v-if="auth.has_permission('iam.permission_group.update')" class="card flex justify-center">
      <Form v-if="pg != null" class="flex flex-col gap-4 w-full sm:w-56" @submit="handle_submit">
         <div class="flex flex-col gap-1">
            <FloatLabel variant="in">
               <InputText id="name" v-model="pg.name" name="name" type="text" fluid />
               <label for="name">{{ $t("iam.pg.update.name") }}</label>
            </FloatLabel>
            <Message v-if="errors.name">{{ errors.name }}</Message>
         </div>
         <div class="flex flex-col gap-1">
            <FloatLabel variant="in">
               <InputText id="abbr" v-model="pg.abbr" name="abbr" type="text" fluid />
               <label for="abbr">{{ $t("iam.pg.update.abbr") }}</label>
            </FloatLabel>
            <Message v-if="errors.name">{{ errors.name }}</Message>
         </div>
         <div class="flex flex-col gap-1">
            <FloatLabel variant="in">
               <InputText
                  id="description"
                  v-model="pg.description"
                  name="description"
                  type="text"
                  fluid
               />
               <label for="description">{{ $t("iam.pg.update.description") }}</label>
            </FloatLabel>
            <Message v-if="errors.name">{{ errors.name }}</Message>
         </div>
         <Button type="submit" severity="secondary" :label="$t('iam.auth.login_form.submit')" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead, PermissionGroupWrite } from "@/api/iam/permission_group/types"
import { to_permission_group_write } from "@/api/iam/permission_group/utils"
import JwtAuthenticator from "@/auth/jwt"
import { error } from "@/composables/toast"
import { ref } from "vue"

const auth = JwtAuthenticator.get_instance()

// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ modelValue: string }>()

const pg = ref<PermissionGroupWrite | null>(null)

const errors = ref<{ name: string; abbr: string; description: string }>({
   name: "",
   abbr: "",
   description: "",
})

const handle_submit = async () => {}

const load_permission_group = async () => {
   ;(await api.iam.permission_group.get(props.modelValue)).fold(
      (error_message: string) => {
         error(error_message, "")
      },
      (permission_group: PermissionGroupRead) => {
         pg.value = to_permission_group_write(permission_group)
      },
   )
}

load_permission_group()
</script>
