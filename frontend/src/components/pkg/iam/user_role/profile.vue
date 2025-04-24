<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="profile" v-permission="'iam.user_role.read'">
         <Card>
            <template #content>
               <MyTimeline
                  :show_now="true"
                  :height="Math.max(400, t_window.innerHeight * 0.4)"
                  :width="t_window.innerWidth"
                  :data="[
                     ...profile.history.map((e) => ({
                        label: e.user.last_name,
                        start: e.start,
                        end: e.end,
                     })),
                  ]"
               ></MyTimeline>
            </template>
         </Card>
         <PermissionGroupProfile :id="profile.role.id"></PermissionGroupProfile>
      </div>
   </div>
</template>
<script lang="ts" setup>
const t_window = window
// profile.users_became_role.map((e) => ({
//    label: e.user.last_name,
//    start: e.start,
//    end: e.end,
// }))
import { api } from "@/api"
import type { UserRoleProfile } from "@/api/iam/user_role/types"
import PermissionGroupProfile from "@/components/pkg/iam/permission_group/profile.vue"
import { error } from "@/composables/toast"
import { ref } from "vue"
import MyTimeline from "@/components/shared/timeline.vue"

const loading = ref<boolean>(true)

const props = defineProps<{ id: string }>()
const profile = ref<UserRoleProfile>()

async function load() {
   loading.value = true
   ;(await api.iam.user_role.read_profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: UserRoleProfile) => {
         profile.value = data
         loading.value = false
      },
   )
   loading.value = false
}

load()
</script>
