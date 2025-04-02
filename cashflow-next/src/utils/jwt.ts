import { JWT } from "@/types";
import { jwtVerify } from "jose";
import { cookies } from "next/headers";

export async function decodeJwtFromCookies(
  field: keyof JWT,
): Promise<JWT[keyof JWT]> {
  const cookieStore = await cookies();
  const token = cookieStore.get("jwt-token");
  try {
    const secret = new TextEncoder().encode(process.env.JWT_SECRET);
    const { payload } = await jwtVerify<JWT>(token?.value ?? "", secret);
    return payload[field];
  } catch (error) {
    console.error(error);
    throw new Error("Invalid JWT token");
  }
}
