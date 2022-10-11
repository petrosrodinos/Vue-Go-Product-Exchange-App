import { Notify } from "quasar";

export const displayErrorMessage = (error, color = "negative") => {
  Notify.create({
    message: error,
    color: color,
    position: "bottom",
    timeout: 2500,
  });
};
