import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const token = request.cookies.get("token")?.value;
  const isAuthPage = request.nextUrl.pathname.startsWith("/auth");
  const isHomePage = request.nextUrl.pathname === "/home";

  // 未ログインユーザーがホームページにアクセスした場合
  if (!token && isHomePage) {
    return NextResponse.redirect(new URL("/auth/login", request.url));
  }

  // ログイン済みユーザーが認証ページにアクセスした場合
  if (token && isAuthPage) {
    return NextResponse.redirect(new URL("/home", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/home", "/auth/:path*"],
};
