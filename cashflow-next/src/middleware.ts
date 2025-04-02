import { NextRequest, NextResponse } from "next/server";
import { jwtVerify } from "jose";

export async function middleware(req: NextRequest) {
  const token = req.cookies.get("jwt-token")?.value;
  const { pathname } = req.nextUrl;

  // Allow login page access without token
  if (pathname.startsWith("/login")) {
    return NextResponse.next();
  }

  // Require token for other pages
  if (!token) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  try {
    // Verify token using jose instead of jsonwebtoken
    const secret = new TextEncoder().encode(process.env.JWT_SECRET);
    const { payload } = await jwtVerify(token, secret);

    const exp = payload.exp;

    if (exp && exp * 1000 > Date.now()) {
      return NextResponse.next();
    } else {
      return NextResponse.redirect(new URL("/login", req.url));
    }
  } catch (error) {
    console.error("Token verification failed:", error);
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

// Define which routes this middleware should run on
export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
};
