export type ProductType = {
  id?: string;
  name: number;
  description: string;
  quantity: number;
  image: File;
  type: string;
  listed: boolean;
  exchangeFor: {
    name: string;
    quantity: number;
  };
  createdAt?: string;
};
