<template>
   <div class="card flex justify-center">
      <Form class="flex flex-col gap-4">
         <FormTextInput v-model="name" label="iam.user_status.create" name="name" />
         <FormTextInput v-model="abbr" label="iam.user_status.create" name="abbr" />
         <FormTextarea v-model="description" label="iam.user_status.create" name="description" />
         <FormSwitch v-model="archived" label="iam.user_status.create" name="archived" />
         <PermissionGroupPicker
            v-model="parent"
            label="iam.user_status.create"
            name="parent"
            @picked="console.log('wowie')"
         />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('iam.user_status.create.submit')"
               @click="create_user_status"
            ></Button>
         </div>
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import { api } from "@/api"
import { error } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextarea from "@/components/form/textarea.vue"
import FormSwitch from "@/components/form/switch.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"

const t = useI18n().t

const name = ref<string>("")
const name_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.name") })),
)

const abbr = ref<string>("")
const abbr_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.abbr") })),
)

const description = ref<string>("")
const description_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.description") })),
)

const archived = ref<boolean>(false)
const archived_check = v.boolean()

const parent = ref<string>()
const parent_check = v.string()

const emit = defineEmits(["created"])

async function create_user_status() {
   const re = v.safeParse(
      v.object({
         name: name_check,
         abbr: abbr_check,
         description: description_check,
         archived: archived_check,
         parent: parent_check,
      }),
      {
         name: name.value,
         abbr: abbr.value,
         description: description.value,
         archived: archived.value,
         parent: parent.value,
      },
   )

   if (re.success) {
      ;(
         await api.iam.user_status.create({
            name: re.output["name"],
            abbr: re.output["abbr"],
            description: re.output["description"],
            archived: re.output["archived"],
            parent: re.output["parent"],
         })
      ).fold(
         (l) => {
            error(l)
         },
         (id: string) => {
            emit("created", id)
         },
      )
   }
}
</script>
