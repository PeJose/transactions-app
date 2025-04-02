import { Container } from "@/components";
import { fetchApi } from "@/utils/api/server";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default function LoginPage() {
  return (
    <Container title="login">
      <form className="flex max-w-max flex-col gap-4 mx-auto" action={login}>
        <label className="input">
          Email:
          <input type="text" name="email" required />
        </label>
        <label className="input">
          Password:
          <input type="password" name="password" required />
        </label>
        <button className="btn btn-primary" type="submit">
          Login
        </button>
      </form>
    </Container>
  );
}

async function login(formData: FormData) {
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

  console.log("Login successful:", token);

  // Handle the response, e.g., set cookies or redirect
  const cookieStore = await cookies();
  cookieStore.set("jwt-token", token);

  redirect("/");
}
