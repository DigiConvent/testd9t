<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <table v-else-if="user_list" v-permission="'iam.user.list'">
      <tr
         v-for="user of user_list.items"
         :key="user.id"
         @click="router.push({ name: 'iam.user.profile', params: { id: user.id } })"
      >
         <td>{{ user.id }}</td>
         <td>{{ user.name }}</td>
         <td>{{ user.implied }}</td>
         <td>{{ user.status_id }}</td>
         <td>{{ user.status_name }}</td>
      </tr>
   </table>
   <router-link v-permission="'iam.user.create'" :to="{ name: 'iam.user.create' }"
      ><Fa icon="user-plus" /> {{ $t("iam.user.create.title") }}</router-link
   >
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { Page } from "@/api/core/page"
import type { UserFacade } from "@/api/iam/user/types"
import { error } from "@/composables/toast"
import router from "@/router"
import { ref } from "vue"

const loading = ref(true)
const user_list = ref<Page<UserFacade>>({
   items: [],
   page: 0,
   items_per_page: 0,
   total_items: 0,
})

async function load_users() {
   loading.value = true
   ;(await api.iam.user.list()).fold(
      (err: string) => {
         error(err)
      },
      (users: Page<UserFacade>) => {
         user_list.value = users
         loading.value = false
      },
   )
}

load_users()
</script>
