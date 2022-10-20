import "quasar/dist/quasar.css";
import iconSet from "quasar/icon-set/material-icons-outlined.js";
import "@quasar/extras/material-icons-outlined/material-icons-outlined.css";
import { Notify } from "quasar";
// To be used on app.use(Quasar, { ... })
export default {
  config: {},
  plugins: { Notify },
  iconSet: iconSet,
};
