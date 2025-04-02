"use client";

import { getCookie } from "cookies-next/client";
import { baseFetchApi } from "./base";

export function fetchApi<T>(
  url: string,
  query?: Record<string, string>,
  config?: RequestInit,
): Promise<T> {
  const cookie = getCookie("jwt-token")!;
  return baseFetchApi(url, cookie?.toString(), query, config);
}
