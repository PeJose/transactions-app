export async function baseFetchApi<T>(
  url: string,
  cookiesString: string,
  query?: Record<string, string>,
  config?: RequestInit,
): Promise<T> {
  const queryString = query
    ? `?${Object.entries(query)
        .map(
          ([key, value]: [string, string]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(value)}`,
        )
        .join("&")}`
    : "";

  const fetchUrl = url + queryString;

  const res = await fetch(fetchUrl, {
    ...config,
    headers: {
      ...config?.headers,
      "Content-Type": "application/json",
      Cookie: cookiesString,
    },
    credentials: "include",
  });

  if (!res.ok) {
    throw new Error(`Failed to fetch data on ${url}`);
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
