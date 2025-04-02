import type { BalanceInTime, Company, TransactionPerCountry } from "@/types";
import { fetchApi } from "@/utils/api";

async function getBalanceInTime(iban: string) {
  return await fetchApi<BalanceInTime[]>("/transactions/chart/balance", {
    iban,
  });
}
async function getTransactionsPerCountry(iban: string) {
  return await fetchApi<TransactionPerCountry[]>("/transactions/map", {
    iban,
  });
}

async function getCompanyInfo(iban: string) {
  return await fetchApi<Company>(`/companies/iban/${iban}`);
}

export { getBalanceInTime, getCompanyInfo, getTransactionsPerCountry };
