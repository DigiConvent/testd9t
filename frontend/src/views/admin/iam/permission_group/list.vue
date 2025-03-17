<template>
   <div>
      <PermissionGroupList @click="console.log($event)"></PermissionGroupList>
      <Dialog
         v-model:visible="visible"
         :dismissable-mask="true"
         pt:root:class="!border-0 !bg-transparent"
         pt:mask:class="backdrop-blur-sm"
      >
         <template #header> {{ focussed_permission_group }} </template>
         <template #container="{ closeCallback }">
            <div
               class="flex flex-col px-8 py-8 gap-6 rounded-2xl"
               style="
                  background-image: radial-gradient(
                     circle at left top,
                     var(--p-primary-400),
                     var(--p-primary-700)
                  );
               "
            >
               {{ focussed_permission_group }}
               <PermissionGroupUpdate v-model="focussed_permission_group"></PermissionGroupUpdate>
               <div class="flex items-center gap-4">
                  <Button
                     label="Cancel"
                     text
                     class="!p-4 w-full !text-primary-50 !border !border-white/30 hover:!bg-white/10"
                     @click="closeCallback"
                  ></Button>
                  <Button
                     label="Sign-In"
                     class="!p-4 w-full !text-primary-50 !border !border-white/30 hover:!bg-white/10"
                     text
                     @click="closeCallback"
                  ></Button>
               </div>
            </div>
         </template>
      </Dialog>
      <!-- <PermissionPicker :multiple="true" @picked="handle_picked"></PermissionPicker> -->
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupList from "@/components/iam/permission_group/list.vue"
import { ref } from "vue"
import PermissionGroupUpdate from "@/components/iam/permission_group/update.vue"

const focussed_permission_group = ref<string>("")
</script>
