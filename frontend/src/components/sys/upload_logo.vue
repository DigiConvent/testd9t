<template>
   <label>
      <img :src="`@/assets/${variant}.png`" alt="" />
      <input type="file" name="logo" accept="image/jpg" @input="on_upload" />
   </label>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { info } from "@/composables/toast"

const props = defineProps<{ variant: "small" | "large" }>()

let file: File
const on_upload = (event: any) => {
   file = event.target.files[0]
   info("Success", "File Uploaded" + event.files, 3000)
   api.sys.upload_logo(props.variant, file)
}
</script>
