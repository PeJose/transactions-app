import { jwtVerify } from "jose";
import { type NextRequest, NextResponse } from "next/server";

export async function middleware(req: NextRequest) {
  const token = req.cookies.get("jwt-token")?.value;
  const { pathname } = req.nextUrl;

  if (pathname.startsWith("/login")) {
    return NextResponse.next();
  }

  if (!token) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  try {
    const secret = new TextEncoder().encode(process.env.JWT_SECRET);
    const { payload } = await jwtVerify(token, secret);

    const exp = payload.exp;

    if (exp && exp * 1000 > Date.now()) {
      const headers = new Headers(req.headers);
      headers.set("x-current-path", req.nextUrl.pathname);
      return NextResponse.next({ headers });
    }
    return NextResponse.redirect(new URL("/login", req.url));
  } catch (error) {
    console.error("Token verification failed:", error);
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
};
