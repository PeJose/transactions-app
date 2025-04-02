export type BalanceInTime = {
  amount: number;
  date: Date | string;
};

export type JWT = {
  email: string;
  user_id: string;
  exp: number;
  iban: string;
};

export type Company = {
  id: number;
  ibans: string[];
  name: string;
  address: string;
};

export type TransactionPerCountry = {
  amount: number;
  occurrances: number;
  country: string;
};
