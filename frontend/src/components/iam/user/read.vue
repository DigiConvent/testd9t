<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Card
         v-else-if="user_read"
         v-permission="'iam.user.read'"
         v-permission.except="is_loggedin_user"
      >
         <template #title
            >{{ user_read.first_name }} {{ user_read.last_name }}
            <router-link
               v-permission="'iam.user.write'"
               v-permission.except="is_loggedin_user"
               :to="{
                  name: 'admin.iam.user.update',
                  params: { id: is_loggedin_user ? 'me' : user_read.id },
               }"
               outlined
               class="!inline"
               ><Fa icon="pencil"
            /></router-link>
         </template>
         <template #content>
            <table>
               <tr>
                  <td>Vorname</td>
                  <td>{{ user_read.first_name }}</td>
               </tr>
               <tr>
                  <td>Nachname</td>
                  <td>{{ user_read.last_name }}</td>
               </tr>
               <tr>
                  <td>E-Postaddresse</td>
                  <td>{{ user_read.emailaddress }}</td>
               </tr>
            </table>
         </template>
      </Card>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { UserRead } from "@/api/iam/user/types"
import JwtAuthenticator from "@/auth/jwt"
import type { IdOrData } from "@/components/form/form"
import { error } from "@/composables/toast"
import { computed, ref } from "vue"

const props = defineProps<IdOrData<UserRead>>()
const loading = ref(true)
const user_read = ref<UserRead | null>(null)
const is_loggedin_user = computed(() => {
   const user_id = JwtAuthenticator.get_instance().get_token()?.id
   if (props.id != undefined && user_id == props.id) {
      return true
   }
   if (user_read.value != undefined && user_id == user_read.value.id) {
      return true
   }
   return false
})

async function load_user() {
   if (props.id === undefined) {
      user_read.value = props.data
      loading.value = false
      return
   }
   ;(await api.iam.user.get(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: UserRead) => {
         user_read.value = data
         loading.value = false
      },
   )
}
load_user()
</script>
