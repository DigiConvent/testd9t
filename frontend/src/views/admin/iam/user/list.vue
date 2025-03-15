<template>
   <div>
      <h2 class="text-2xl">{{ $t("iam.user.list.title") }}</h2>
      <table>
         <tr v-for="user of user_list.items" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.name }}</td>
            <td>{{ user.implied }}</td>
            <td>{{ user.status_id }}</td>
            <td>{{ user.status_name }}</td>
         </tr>
      </table>
      <router-view />
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { Page } from "@/api/core/page"
import type { UserFacade } from "@/api/iam/user/types"
import { useToast } from "primevue"
import { ref } from "vue"

const toast = useToast()

const user_list = ref<Page<UserFacade>>({
   items: [],
   page: 0,
   items_per_page: 0,
   total_items: 0,
})

const load_users = async () => {
   const users = api.iam.user.list()

   ;(await users).fold(
      (error: string) => {
         toast.add({
            severity: "error",
            summary: "Error",
            detail: error,
         })
      },
      (users: Page<UserFacade>) => {
         console.log(users)
         user_list.value = users
      },
   )
}

load_users()
</script>
