<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <DataTable v-else-if="user_list" v-permission="'iam.user.list'" :value="user_list.items">
      <Column field="first_name" :header="$t('iam.user.fields.first_name')" />
      <Column field="last_name" :header="$t('iam.user.fields.last_name')" />
      <Column :header="$t('iam.user.fields.user_status')">
         <template #body="slotProps">
            <router-link
               v-if="slotProps.data.status_id != null"
               :to="{
                  name: 'admin.iam.user_status.profile',
                  params: { id: slotProps.data.status_id },
               }"
               >{{ slotProps.data.status_name }}</router-link
            >
            <Fa v-else icon="circle-exclamation" class="text-red-500" />
         </template>
      </Column>
      <Column :header="$t('iam.user.fields.roles')">
         <template #body="slotProps">
            <div class="flex shrink gap-4">
               <router-link
                  v-for="(role, i) of slotProps.data.roles"
                  :key="role.id"
                  :to="{ name: 'admin.iam.user_role.profile', params: { id: role.id } }"
                  :style="`background-color: ${colours[i]}`"
                  class="white-text p-1 rounded"
                  >{{ role.name }}</router-link
               >
            </div>
         </template>
      </Column>
      <Column v-permission="'iam.user.read'">
         <template #body="slotProps">
            <Button class="p-button-text" @click="handle_row_click(slotProps)">
               <Fa icon="id-card" class="text-xl"></Fa>
            </Button>
         </template>
      </Column>
   </DataTable>
   <router-link v-permission="'iam.user.create'" :to="{ name: 'admin.iam.user.create' }"
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

function handle_row_click(data: any) {
   router.push({ name: "admin.iam.user.profile", params: { id: data.data.id } })
}

const colours = [
   "#ef4444",
   "#f59e0b",
   "#84cc16",
   "#10b981",
   "#06b6d4",
   "#3b82f6",
   "#8b5cf6",
   "#d946ef",
   "#f43f5e",
   "#f97316",
   "#eab308",
   "#22c55e",
   "#14b8a6",
   "#0ea5e9",
   "#6366f1",
   "#a855f7",
   "#ec4899",
]
</script>
