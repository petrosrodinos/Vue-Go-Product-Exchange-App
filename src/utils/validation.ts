import { PasswordValidator } from "@/types/auth";
import { reactive } from "vue";

const validPassword = reactive<PasswordValidator>({
  length: false,
  capital: false,
  number: false,
  symbol: false,
});

export function validateEmail(email: string) {
  return /[a-z0-9]+@[a-z]+\.[a-z]{2,3}/.test(email);
}

export function validatePassword(password: string) {
  validPassword.length = password.length >= 12;
  validPassword.capital = /^(?=.*[A-Z]).*$/.test(password);
  validPassword.number = /^(?=.*[0-9]).*$/.test(password);
  validPassword.symbol = /^(?=.*[!@#$%^&*()\-_+=]).*$/.test(password);
  return (
    validPassword.length &&
    validPassword.capital &&
    validPassword.number &&
    validPassword.symbol
  );
}
