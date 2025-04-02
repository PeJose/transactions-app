import { join } from "node:path";
import { cookies } from "next/headers";

export async function fetchApi<T>(
  url: string,
  query?: Record<string, string>,
  config?: RequestInit,
): Promise<T> {
  const cookieStore = await cookies();
  const queryString = query
    ? "?" +
      Object.entries(query)
        .map(
          ([key, value]: [string, string]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(value)}`,
        )
        .join("&")
    : "";

  const fetchUrl = join(process.env.API_URL!, url + queryString);
  console.log(fetchUrl);

  const res = await fetch(fetchUrl, {
    ...config,
    headers: {
      ...config?.headers,
      "Content-Type": "application/json",
      Cookie: cookieStore.toString(),
    },
    credentials: "include",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  const data = (await res.json()) as ApiSuccessResponse<T>;

  if (data.status !== "success") {
    throw new Error(data.message);
  }
  return data.data;
}

type ApiSuccessResponse<T> = {
  message: string;
  data: T;
  status: "success" | "error";
};
