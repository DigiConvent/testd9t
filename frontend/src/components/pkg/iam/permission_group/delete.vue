<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Card v-else-if="profile" class="flex flex-col gap-2">
         <template #content>
            <div v-if="profile.descendants.length > 0 || profile.users.length > 0">
               {{ $t("iam.pg.delete.warning") }}
               <div v-if="profile.users.length > 0">
                  <h1>{{ $t("iam.user.user", { count: profile.users.length }) }}</h1>
                  <UserFacades :users="profile.users" />
               </div>
               <br />
               <div v-if="profile.descendants.filter((e) => e.meta == 'role').length > 0">
                  <h1>
                     {{
                        $t("iam.ur.ur", {
                           count: profile.descendants.filter((e) => e.meta == "role").length,
                        })
                     }}
                  </h1>
                  <PermissionGroupFacades
                     :permission_groups="profile.descendants.filter((e) => e.meta == 'role')"
                  />
               </div>
               <div v-if="profile.descendants.filter((e) => e.meta == 'status').length > 0">
                  <h1>
                     {{
                        $t("iam.us.us", {
                           count: profile.descendants.filter((e) => e.meta == "status").length,
                        })
                     }}
                  </h1>
                  <PermissionGroupFacades
                     :permission_groups="profile.descendants.filter((e) => e.meta == 'status')"
                  />
               </div>
               <div v-if="profile.descendants.filter((e) => e.meta == null).length > 0">
                  <h1>
                     {{
                        $t("iam.pg.pg", {
                           count: profile.descendants.filter((e) => e.meta == null).length,
                        })
                     }}
                  </h1>
                  <PermissionGroupFacades
                     :permission_groups="profile.descendants.filter((e) => e.meta == null)"
                  />
               </div>
            </div>
            <div v-else class="gap-2 flex flex-wrap">
               {{
                  $t("prompt.delete", {
                     entity: t("iam.pg.pg"),
                     name: profile.permission_group.name,
                  })
               }}
               <br />
               <Button severity="danger" @click="handle_delete">
                  {{ $t("answers.yes") }}
               </Button>
               <Button @click="router.back()">{{ $t("answers.no") }}</Button>
            </div>
         </template>
      </Card>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupProfile } from "@/api/iam/permission_group/types"
import UserFacades from "@/components/pkg/iam/user/facade.vue"
import PermissionGroupFacades from "./facade.vue"
import { error, success } from "@/composables/toast"
import router from "@/router"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const loading = ref<boolean>(true)
const t = useI18n().t
const props = defineProps<{ id: string }>()

async function handle_delete() {
   ;(await api.iam.permission_group.delete(props.id)).fold(
      (err: string) => {
         error(t("feedback.-.delete", { entity: t("iam.pg.pg") }), err)
      },
      () => {
         success(t("feedback.+.delete", { entity: t("iam.pg.pg") }))
         router.back()
      },
   )
}

const profile = ref<PermissionGroupProfile>()
async function load() {
   loading.value = true
   ;(await api.iam.permission_group.read_profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (pg: PermissionGroupProfile) => {
         profile.value = pg
         loading.value = false
      },
   )
}

load()
</script>
