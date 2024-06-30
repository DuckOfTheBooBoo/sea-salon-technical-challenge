import { createApp } from "vue";
import "@/assets/index.css";
import App from "@/App.vue";
import router from "@/router";
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
import 'vue-fullpage.js/dist/style.css'
import VueFullPage from 'vue-fullpage.js'

createApp(App).use(router).use(VueFullPage).use(autoAnimatePlugin).mount("#app");
