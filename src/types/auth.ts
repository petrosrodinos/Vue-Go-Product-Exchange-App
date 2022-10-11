export type RegisterState = {
  firstName: string;
  lastName: string;
  phoneNumber: string;
  email: string | null;
  password: string;
  confirmPassword: string;
  address: {
    city: string;
    area: string;
  };
};

export type PasswordValidator = {
  length: boolean;
  capital: boolean;
  number: boolean;
  symbol: boolean;
};
