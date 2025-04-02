import { headers } from "next/headers";

export async function getPathname() {
  const headerList = await headers();
  return headerList.get("x-current-path");
}
