export type Trade = {
  status: "cancelled" | "completed" | "pending";
  name1: string;
  name2: string;
  quantity1: string;
  quantity2: string;
  traderName: string;
  traderPhone: string;
  userId: string;
  traderId: string;
  id: string;
};
