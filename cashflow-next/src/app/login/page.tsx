import { login } from "@/actions/auth";
import { Container } from "@/components";

export default function LoginPage() {
  return (
    <Container title="Sign In" isLogin>
      <form className="mx-auto flex max-w-max flex-col gap-4" action={login}>
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
