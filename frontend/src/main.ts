import "./assets/main.css"

import { createApp } from "vue"
import PrimeVue from "primevue/config"
import Aura from "@primevue/themes/aura"

declare global {
   interface Window {
      debug: boolean
      enable_gesture_navigation: boolean
   }
}

window.debug = process.env.NODE_ENV === "development"
window.enable_gesture_navigation = true
// window.debug = false

import { library, config } from "@DigiConvent/ff/fontawesome-svg-core"
import { FontAwesomeIcon, FontAwesomeLayers } from "@DigiConvent/ff/vue-fontawesome"
import { fasds } from "@DigiConvent/ff/sharp-duotone-solid-svg-icons"
import { faGithub } from "@DigiConvent/ff/free-brands-svg-icons"

import App from "./app/App.vue"
import router from "./router"

import Toast from "primevue/toast"
import { Form } from "@primevue/forms"
import {
   Accordion,
   AccordionContent,
   AccordionHeader,
   AccordionPanel,
   Badge,
   Breadcrumb,
   Button,
   Card,
   Checkbox,
   Column,
   DataTable,
   Dialog,
   Drawer,
   Fieldset,
   FloatLabel,
   InputGroup,
   InputGroupAddon,
   InputMask,
   InputText,
   Listbox,
   Menu,
   Menubar,
   Message,
   MeterGroup,
   OrganizationChart,
   Popover,
   ProgressBar,
   ProgressSpinner,
   Ripple,
   Select,
   Skeleton,
   Splitter,
   SplitterPanel,
   Step,
   StepItem,
   StepPanel,
   Stepper,
   Tag,
   Textarea,
   Timeline,
   ToastService,
   ToggleButton,
   ToggleSwitch,
   TreeSelect,
} from "primevue"

const app = createApp(App)

app.use(PrimeVue, { theme: { preset: Aura } })

import { is_mini_app } from "./auth/telegram"
import JwtAuthenticator from "./auth/jwt"
import { permission } from "./directives/permission"
import { i18n } from "./locales/i18n"
import { info } from "./composables/toast"

app.use(i18n)

config.familyDefault = "sharp-duotone"
library.add(fasds)
library.add(faGithub)

app.use(router)

app.component("Accordion", Accordion)
app.component("AccordionPanel", AccordionPanel)
app.component("AccordionHeader", AccordionHeader)
app.component("AccordionContent", AccordionContent)
app.component("Badge", Badge)
app.component("Breadcrumb", Breadcrumb)
app.component("Button", Button)
app.component("Card", Card)
app.component("Column", Column)
app.component("DataTable", DataTable)
app.component("Dialog", Dialog)
app.component("Drawer", Drawer)
app.component("Checkbox", Checkbox)
app.component("Fieldset", Fieldset)
app.component("FloatLabel", FloatLabel)
app.component("Form", Form)
app.component("InputGroupAddon", InputGroupAddon)
app.component("InputGroup", InputGroup)
app.component("InputMask", InputMask)
app.component("InputText", InputText)
app.component("Listbox", Listbox)
app.component("Menu", Menu)
app.component("Menubar", Menubar)
app.component("MeterGroup", MeterGroup)
app.component("Message", Message)
app.component("OrganizationChart", OrganizationChart)
app.component("Popover", Popover)
app.component("ProgressBar", ProgressBar)
app.component("ProgressSpinner", ProgressSpinner)
app.component("Tag", Tag)
app.component("Textarea", Textarea)
app.component("Timeline", Timeline)
app.component("Toast", Toast)
app.component("ToggleButton", ToggleButton)
app.component("ToggleSwitch", ToggleSwitch)
app.component("TreeSelect", TreeSelect)
app.component("Select", Select)
app.component("Skeleton", Skeleton)
app.component("Splitter", Splitter)
app.component("SplitterPanel", SplitterPanel)
app.component("Stepper", Stepper)
app.component("Step", Step)
app.component("StepItem", StepItem)
app.component("StepPanel", StepPanel)

app.directive("ripple", Ripple)
app.directive("permission", permission)

app.component("Fa", FontAwesomeIcon).component("fal", FontAwesomeLayers)

let reuser = window.location.href.replace(window.location.origin, "")
app.use(ToastService)

const auth = JwtAuthenticator.get_instance()
if (is_mini_app()) {
   // I don't want to recover a session, I want to start a new one
   auth.login_using_telegram().then(() => {
      mount(reuser)
   })
} else if (auth.sessions.length == 0) {
   mount(reuser)
} else if (auth.sessions.length == 1) {
   auth.recover_session().then((is_authenticated) => {
      if (is_authenticated) {
         mount(reuser)
         info("Session recovered", "")
      } else {
         info("Session not recovered", "")
      }
   })
} else if (auth.sessions.length > 1) {
   mount("/auth/session/")
}

function mount(path_to_load: string = "/app/") {
   app.mount("#app")
   try {
      router.replace({ path: path_to_load })
   } catch (e) {
      console.log(e)
   }
}
