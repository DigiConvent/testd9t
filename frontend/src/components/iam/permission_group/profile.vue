<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <NeedsPermission v-else-if="profile" permission="iam.permission_group.read">
      <div class="grid grid-cols-2 gap-4">
         <Card>
            <template #title> Hierarchy </template>
            <template #content>
               <Timeline :value="profile?.ancestors.reverse()">
                  <template #content="slotProps">
                     <router-link
                        :to="{ name: 'iam.pg.profile', params: { id: slotProps.item.id } }"
                        >{{ slotProps.item.name }}</router-link
                     >
                  </template>
               </Timeline>
            </template>
         </Card>
         <Card>
            <template #title>Permissions</template>
            <template #content>
               <div class="grid grid-cols-2">
                  <div>
                     Inherited
                     <div
                        v-for="p of profile?.permissions.filter((p) => p.implied)"
                        :key="'inherited' + p.name"
                     >
                        {{ p.name }}
                     </div>
                  </div>
                  <div>
                     Owned
                     <div
                        v-for="p of profile?.permissions.filter((p) => !p.implied)"
                        :key="'inherited' + p.name"
                     >
                        {{ p.name }}
                     </div>
                  </div>
               </div>
            </template>
         </Card>
         <Card>
            <template #title>Data</template>
            <template #content>
               <UpdatePermissionGroup
                  v-model="profile!.permission_group.id"
               ></UpdatePermissionGroup>
            </template>
         </Card>
         <Card>
            <template #title>Members</template>
            <template #content>
               <UserFacades v-if="profile.members.length > 0" :users="profile!.members" />
               <span v-else>
                  {{ $t("iam.pg.profile.no_members") }}
               </span>
            </template>
         </Card>
      </div>
   </NeedsPermission>
   <div v-else>Something went wrong loading this permission group</div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupProfile } from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { ref } from "vue"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"
import UserFacades from "@/components/iam/user/facade.vue"

const props = defineProps<{ id: string }>()

const loading = ref(true)
const profile = ref<PermissionGroupProfile | null>(null)
async function load_profile() {
   loading.value = true
   ;(await api.iam.permission_group.get_profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (pg: PermissionGroupProfile) => {
         profile.value = pg
         loading.value = false
      },
   )
}

load_profile()
</script>
