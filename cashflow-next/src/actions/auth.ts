import { fetchApi } from "@/utils/api/server";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function login(formData: FormData) {
  "use server";
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;

  const token = await fetchApi<string>(
    "auth/login",
    {},
    {
      method: "POST",
      body: JSON.stringify({ email, password }),
    },
  );

  const cookieStore = await cookies();
  cookieStore.set("jwt-token", token);

  redirect("/");
}
