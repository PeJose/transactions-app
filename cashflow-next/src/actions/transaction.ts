import type { Balance } from "@/types";
import type { BalanceInTime, TransactionPerCountry } from "@/types";
import { fetchApi } from "@/utils/api/server";

async function getBalanceInTime(iban: string) {
  return await fetchApi<BalanceInTime[]>("/transactions/charts/balance", {
    iban,
  });
}
async function getTransactionsPerCountry(iban: string) {
  return await fetchApi<TransactionPerCountry[]>("/transactions/charts/map", {
    iban,
  });
}

async function getBalance(iban: string) {
  return await fetchApi<Balance>("/transactions/balance", {
    currency: "EUR",
    iban,
  });
}

export { getBalanceInTime, getTransactionsPerCountry, getBalance };
