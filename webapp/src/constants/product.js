export const productColumns = [
  {
    name: "name",
    align: "left",
    label: "Name",
    field: "name",
  },
  { name: "quantity", label: "Quantity", field: "quantity" },
  { name: "type", label: "Type", field: "type" },
  { name: "exchangeFor", label: "Exchange For", field: "exchangeFor" },
  { name: "createdAt", label: "Created At", field: "createdAt" },
];

const columns = [
  {
    name: "name1",
    align: "left",
    label: "Name 1",
    field: "name1",
  },
  { name: "quantity1", label: "Quantity 1", field: "quantity1" },
  { name: "name2", label: "Name 2", field: "name2" },
  { name: "quantity2", label: "Quantity 2", field: "quantity2" },
  { name: "traderName", label: "Trader Name", field: "traderName" },
  { name: "traderPhone", label: "Trader Phone", field: "traderPhone" },
];

export const CancelledColumns = columns;

export const TradeColumns = [
  ...columns,
  { name: "action", label: "Action", field: "action" },
];

export const productTypes = ["Food", "Clothes", "Electronics", "Other"];
