import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./state/store";
import { Quasar } from "quasar";
import quasarUserOptions from "./quasar-user-options";
import "./assets/style.css";

createApp(App)
  .use(Quasar, quasarUserOptions)
  .use(router)
  .use(store)
  .mount("#app");
