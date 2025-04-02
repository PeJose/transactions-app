import type { Company } from "@/types";
import { fetchApi } from "@/utils/api/server";

export async function getCompanyInfo(iban?: string) {
  return await fetchApi<Company>(`/companies/iban/${iban}`);
}
