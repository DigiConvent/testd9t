<template>
   <div class="card flex justify-center">
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <Form v-else class="flex flex-col gap-4">
         <FormTextInput v-model="name" label="iam.user_status.create" name="name" />
         <FormTextInput v-model="abbr" label="iam.user_status.create" name="abbr" />
         <FormTextarea v-model="description" label="iam.user_status.create" name="description" />
         <FormSwitch
            v-model="archived"
            label_on="iam.user_status.create.archived"
            label_off="iam.user_status.create.unarchived"
            name="archived"
         />
         <PermissionGroupPicker v-model="us_parent" label="iam.user_status.create" name="parent" />
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('iam.us.create.submit')"
               @click="create_user_status"
            ></Button>
         </div>
      </Form>
      <component :is="form.v_node"></component>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import { api } from "@/api"
import { error, warn } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormTextarea from "@/components/form/textarea.vue"
import FormSwitch from "@/components/form/switch.vue"
import PermissionGroupPicker from "@/components/iam/permission_group/picker.vue"
import { Form } from "@/components/form2/form"
import type { UserStatusWrite } from "@/api/iam/user_status/types"

const t = useI18n().t

const form = new Form<UserStatusWrite>({
   t: "iam.user_status.create",
   save_endpoint: api.iam.user_status.create,
   data: { id: "", name: "", abbr: "", description: "", archived: false, parent: "" },
})
   .add_text_input({ key: "name", label: "iam.user_status.new.name" })
   .add_text_input({
      key: "abbr",
      label: "iam.user_status.new.abbr",
   })
   .add_text_input({ key: "description", label: "iam.user_status.new.description" })
   .add_switch_input({
      key: "archived",
      label: "iam.user_status.new.archived",
   })

const loading = ref(false)
const name = ref<string>("")
const name_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.us.new.required", { field: t("iam.us.new.name") })),
)

const abbr = ref<string>("")
const abbr_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.us.new.required", { field: t("iam.us.new.abbr") })),
)

const description = ref<string>("")
const description_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.us.new.required", { field: t("iam.us.new.description") })),
)

const archived = ref<boolean>(false)
const archived_check = v.boolean()

const us_parent = ref<string>()
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
         parent: us_parent.value,
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
         (l: string) => {
            error(l)
         },
         (id: string) => {
            emit("created", id)
         },
      )
   } else {
      warn(re.issues[0].message)
   }
}

const props = defineProps<{ parent?: string }>()
function load() {
   loading.value = true
   if (props.parent != undefined) {
      us_parent.value = props.parent
   }
   loading.value = false
}

load()
</script>
