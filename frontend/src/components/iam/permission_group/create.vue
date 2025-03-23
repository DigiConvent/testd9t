<template>
   <div v-if="auth.has_permission('iam.permission_group.update')" class="card flex justify-center">
      create pg
      <Form class="flex flex-col gap-4 w-full sm:w-56" @submit="handle_submit">
         <FormTextInput v-model="pg.name" label="iam.pg.fields" :error="errors.name" name="name" />
         <FormTextInput v-model="pg.abbr" label="iam.pg.fields" :error="errors.abbr" name="abbr" />
         <FormTextAreaInput
            v-model="pg.name"
            label="iam.pg.fields"
            :error="errors.name"
            name="description"
         />
         <PermissionGroupPicker
            v-model="pg.parent"
            @picked="pg.parent = $event"
         ></PermissionGroupPicker>
         <PermissionPicker
            :multiple="true"
            @picked="console.log($event)"
            @selected="console.log($event)"
         ></PermissionPicker>
         <Button type="submit" severity="secondary" :label="$t('iam.auth.login_form.submit')" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupWrite } from "@/api/iam/permission_group/types"
import JwtAuthenticator from "@/auth/jwt"
import { ref } from "vue"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextAreaInput from "@/components/form/textarea.vue"
import PermissionGroupPicker from "./picker.vue"
import PermissionPicker from "@/components/iam/permission/list.vue"

const auth = JwtAuthenticator.get_instance()

const pg = ref<PermissionGroupWrite>({
   name: "",
   abbr: "",
   description: "",
   parent: undefined,
   is_group: true,
   is_node: false,
   permissions: [],
})

const errors = ref<{ name: string; abbr: string; description: string }>({
   name: "",
   abbr: "",
   description: "",
})

const handle_submit = async () => {
   ;(await api.iam.permission_group.create(pg.value!)).fold(
      (error: string) => {
         console.log(error)
      },
      (id: string) => {
         console.log(id)
      },
   )
}
</script>
