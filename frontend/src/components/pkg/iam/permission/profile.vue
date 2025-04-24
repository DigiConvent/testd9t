<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="profile" class="flex flex-col gap-4">
         <Card>
            <template #content>
               <router-link
                  v-for="permission of ancestors"
                  :key="permission.name"
                  class="block"
                  :to="{ name: 'admin.iam.permission.profile', params: { name: permission.name } }"
               >
                  {{ permission.name }}
               </router-link>
               <h1 class="text-2xl">
                  {{ profile.permission.name }}
                  <small v-if="profile.permission.archived">{{
                     $t("iam.p.profile.archived")
                  }}</small>
                  <p>{{ profile.permission.description }}</p>
               </h1>
               <router-link
                  v-for="permission of profile.descendants"
                  :key="permission.name"
                  class="block"
                  :to="{ name: 'admin.iam.permission.profile', params: { name: permission.name } }"
               >
                  {{ permission.name }}
               </router-link>
            </template>
         </Card>
         <Card>
            <template #title>
               {{ $t("iam.p.profile.users") }}
            </template>
            <template #content>
               <UserFacades :users="profile.users"></UserFacades>
            </template>
         </Card>
         <Card>
            <template #title>
               {{ $t("iam.p.profile.groups") }}
               <a @click="show_permission_group_picker = true"><Fa icon="plus" /></a>
            </template>
            <template #content>
               <div v-for="pg of profile.permission_groups" :key="pg.id">
                  <Button
                     class="p-button-text"
                     severity="danger"
                     @click="handle_remove_permission({ event: $event, pg: pg })"
                  >
                     <Fa icon="times" /></Button
                  ><Fa :icon="get_icon(pg.meta)" class="fa-fw" /> {{ pg.name }}
               </div>
               <Dialog v-model:visible="show_permission_group_picker" modal>
                  <PermissionGroupPicker
                     v-model="permission_group_to_add"
                     label="iam.p.fields"
                     name="group"
                  />
                  <Button @click="handle_add_permission_group">{{
                     $t("actions.add", { entity: $t("iam.pg.pg") })
                  }}</Button>
               </Dialog>
            </template>
         </Card>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { type PermissionProfile } from "@/api/iam/permission/types"
import UserFacades from "@/components/pkg/iam/user/facade.vue"
import { error, success } from "@/composables/toast"
import { computed, ref } from "vue"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import PermissionGroupPicker from "../permission_group/picker.vue"
import { useI18n } from "vue-i18n"
import { get_icon } from "@/api/iam/permission_group/utils"

const loading = ref<boolean>(true)
const props = defineProps<{ id: string }>()
const profile = ref<PermissionProfile>()
const show_permission_group_picker = ref(false)
const permission_group_to_add = ref<string>()

const t = useI18n().t

const ancestors = computed(() => {
   if (!profile.value) return []
   const ancestors = []
   const segments = profile.value.permission.name
   let result = ""
   for (const segment of segments.split(".")) {
      result += `${result ? "." : ""}${segment}`
      if (result == profile.value.permission.name) break
      ancestors.push({ name: result })
   }
   return ancestors
})

async function load() {
   ;(await api.iam.permission.read_profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: PermissionProfile) => {
         profile.value = data
         loading.value = false
      },
   )
}

load()

async function handle_remove_permission(event: { event: Event; pg: PermissionGroupFacade }) {
   ;(await api.iam.permission_group.remove_permission(event.pg.id, props.id)).fold(
      (err: string) => {
         error(t("iam.p.profile.remove_from_group.error"), err)
      },
      () => {
         success(t("iam.p.profile.remove_from_group.success"))
         load()
      },
   )
}
async function handle_add_permission_group() {
   show_permission_group_picker.value = false
   ;(await api.iam.permission_group.add_permission(permission_group_to_add.value, props.id)).fold(
      (err: string) => {
         error(t("iam.p.profile.add_to_group.error"), err)
      },
      () => {
         success(t("iam.p.profile.add_to_group.success"))
         load()
      },
   )
}
</script>
