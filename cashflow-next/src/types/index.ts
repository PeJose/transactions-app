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

export type Balance = number;
export type CompanyWithBalance = Balance & Company;

export type TransactionPerCountry = {
  amount: number;
  occurrances: number;
  country: string;
};

export type InfiniteTransaction<T> = {
  hasNext: boolean;
  lastCursor: string;
  transactions: T[];
};

export type Transaction = {
  id: string;
  amount: number;
  currency: string;
  timestamp: string;
  from: string;
  to: string;
};

export type TransactionType = "SEPA" | "SWIFT";
export type TransactionTarget = "from" | "to";
