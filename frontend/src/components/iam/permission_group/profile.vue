<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <NeedsPermission v-else-if="profile" permission="iam.permission_group.read">
      <div class="grid grid-cols-2 gap-4">
         <Card>
            <template #title> {{ $t("iam.pg.profile.hierarchy") }} </template>
            <template #content>
               <Timeline :value="ancestors">
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
            <template #title>{{ $t("iam.pg.profile.permissions") }}</template>
            <template #content>
               <div class="grid grid-cols-2">
                  <div>
                     {{ $t("iam.pg.profile.inherited") }}
                     <div v-for="p of profile?.permissions.filter((p) => p.implied)" :key="p.name">
                        {{ p.name }}
                     </div>
                  </div>
                  <div>
                     {{ $t("iam.pg.profile.owned") }}
                     <div v-for="p of profile?.permissions.filter((p) => !p.implied)" :key="p.name">
                        {{ p.name }}
                     </div>
                  </div>
               </div>
            </template>
         </Card>
         <Card>
            <template #title>{{ $t("iam.pg.profile.properties") }}</template>
            <template #content>
               <UpdatePermissionGroup
                  v-model="profile!.permission_group.id"
               ></UpdatePermissionGroup>
            </template>
         </Card>
         <Card>
            <template #title>{{ $t("iam.pg.profile.members") }}</template>
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
import type {
   PermissionGroupFacade,
   PermissionGroupProfile,
} from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { computed, ref } from "vue"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"
import UserFacades from "@/components/iam/user/facade.vue"

const props = defineProps<{ id: string }>()

const loading = ref(true)
const profile = ref<PermissionGroupProfile | null>(null)
const ancestors = computed({
   get: () => {
      const ancestors: PermissionGroupFacade[] = []
      for (const pg of profile.value!.ancestors) {
         ancestors.push(pg)
      }
      return ancestors.reverse()
   },
   set: () => {},
})
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
