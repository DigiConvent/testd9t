<template>
   <template v-if="allowed">
      <Fieldset v-if="debug">
         <template #legend>
            <span class="text-xs">{{ props.permission }}</span>
         </template>
         <slot />
      </Fieldset>
      <template v-else>
         <slot></slot>
      </template>
   </template>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import { ref, watch } from "vue"

const props = defineProps<{ permission: string }>()

const auth = JwtAuthenticator.get_instance()
const is_authenticated = auth.is_authenticated

const allowed = ref(false)
const debug = ref(true || auth.has_permission("super"))

watch(is_authenticated, () => {
   check_permission()
})

function check_permission() {
   if (is_authenticated.value) {
      allowed.value = auth.has_permission(props.permission)
      debug.value = true || auth.has_permission("super")
   } else {
      allowed.value = false
   }
}
check_permission()
</script>
