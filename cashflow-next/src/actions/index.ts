import type { Balance, Company } from "@/types";
import type { BalanceInTime, TransactionPerCountry } from "@/types";
import { fetchApi } from "@/utils/api/server";
import { decodeJwtFromCookies } from "@/utils/jwt";

async function externalOrUserIban(companyIban?: string) {
  return companyIban ?? (await decodeJwtFromCookies("iban"));
}

async function getBalanceInTime(companyIban?: string) {
  const iban = await externalOrUserIban(companyIban);
  if (!iban) {
    return undefined;
  }
  return await fetchApi<BalanceInTime[]>("/transactions/chart/balance", {
    iban: iban.toString(),
  });
}
async function getTransactionsPerCountry(companyIban?: string) {
  const iban = await externalOrUserIban(companyIban);
  if (!iban) {
    return undefined;
  }
  return await fetchApi<TransactionPerCountry[]>("/transactions/map", {
    iban: iban.toString(),
  });
}
async function getCompanyInfo(companyIban?: string) {
  const iban = await externalOrUserIban(companyIban);
  if (!iban) {
    return undefined;
  }
  return await fetchApi<Company>(`/companies/iban/${iban}`);
}

async function getBalance(companyIban?: string) {
  const iban = await externalOrUserIban(companyIban);
  if (!iban) {
    return undefined;
  }
  return await fetchApi<Balance>("/transactions/balance", {
    currency: "EUR",
    iban: iban.toString(),
  });
}

export {
  getCompanyInfo,
  getBalanceInTime,
  getTransactionsPerCountry,
  getBalance,
};
