import { join } from "node:path";
import { cookies } from "next/headers";
import { baseFetchApi } from "./base";

export async function fetchApi<T>(
  url: string,
  query?: Record<string, string>,
  config?: RequestInit,
): Promise<T> {
  const cookieStore = await cookies();
  const fetchUrl = join(process.env.NEXT_PUBLIC_API_URL!, url);
  return await baseFetchApi(fetchUrl, cookieStore.toString(), query, config);
}
