<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div
         v-else-if="profile"
         v-permission="'iam.permission_group.read'"
         class="grid grid-cols-2 gap-4"
      >
         <Card>
            <template #title> {{ $t("iam.pg.profile.hierarchy") }} </template>
            <template #content>
               <Timeline :value="ancestors">
                  <template #content="slotProps">
                     <router-link
                        :to="{
                           name: 'admin.iam.permission_group.profile',
                           params: { id: slotProps.item.id },
                        }"
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
                     <router-link
                        v-for="p of profile?.permissions.filter((p) => p.implied)"
                        :key="p.name"
                        class="block"
                        :to="{ name: 'admin.iam.permission.profile', params: { name: p.name } }"
                     >
                        {{ p.name }}
                     </router-link>
                  </div>
                  <div>
                     {{ $t("iam.pg.profile.owned") }}
                     <router-link
                        v-for="p of profile?.permissions.filter((p) => !p.implied)"
                        :key="p.name"
                        class="block"
                        :to="{ name: 'admin.iam.permission.profile', params: { name: p.name } }"
                     >
                        {{ p.name }}
                     </router-link>
                  </div>
               </div>
            </template>
         </Card>
         <Card>
            <template #title
               >{{ $t("iam.pg.profile.properties") }}
               <router-link
                  v-permission="'iam.permission_group.write'"
                  :to="{
                     name: 'admin.iam.permission_group.update',
                     params: { id: profile!.permission_group.id },
                  }"
                  class="!inline"
                  ><Fa icon="pencil" /></router-link
            ></template>
            <template #content>
               <ReadPermissionGroup :data="profile!.permission_group" />
            </template>
         </Card>
         <Card>
            <template #title>{{ $t("iam.pg.profile.users") }}</template>
            <template #content>
               <UserFacades v-if="profile.users.length > 0" :users="profile!.users" />
               <span v-else>
                  {{ $t("iam.pg.profile.no_users") }}
               </span>
            </template>
         </Card>
      </div>
      <div v-else>Something went wrong loading this permission group</div>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type {
   PermissionGroupFacade,
   PermissionGroupProfile,
} from "@/api/iam/permission_group/types"
import { error } from "@/composables/toast"
import { computed, ref } from "vue"
import UserFacades from "@/components/pkg/iam/user/facade.vue"
import ReadPermissionGroup from "./read.vue"

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

load_profile()
</script>
