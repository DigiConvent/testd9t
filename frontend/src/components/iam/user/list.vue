<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <NeedsPermission v-else-if="user_list" :permission="'iam.user.list'">
      <table>
         <tr v-for="user of user_list.items" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.name }}</td>
            <td>{{ user.implied }}</td>
            <td>{{ user.status_id }}</td>
            <td>{{ user.status_name }}</td>
         </tr>
      </table>
      <NeedsPermission permission="iam.user.create">
         <router-link :to="{ name: 'iam.user.create' }"
            ><Fa icon="user-plus" /> {{ $t("iam.user.create.title") }}</router-link
         ></NeedsPermission
      >
   </NeedsPermission>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { Page } from "@/api/core/page"
import type { UserFacade } from "@/api/iam/user/types"
import { useToast } from "primevue"
import { ref } from "vue"

const toast = useToast()

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
      (error: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: error,
         })
      },
      (users: Page<UserFacade>) => {
         user_list.value = users
         loading.value = false
      },
   )
}

load_users()
</script>
