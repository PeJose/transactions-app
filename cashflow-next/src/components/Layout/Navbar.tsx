import { getBalance } from "@/actions/transaction";
import { decodeJwtFromCookies } from "@/utils/jwt";
import type { ReactNode } from "react";
import { PiHamburger } from "react-icons/pi";

type NavbarProps = {
  title: ReactNode;
  isLogin?: boolean;
};

export default async function Navbar({ title, isLogin = false }: NavbarProps) {
  const iban = isLogin ? undefined : await decodeJwtFromCookies("iban");
  const balance = isLogin ? undefined : await getBalance(iban!.toString());
  return (
    <header className="p-2">
      <div className="navbar rounded-xl bg-secondary px-8 text-secondary-content">
        <div className="flex-none">
          <label
            htmlFor="my-drawer"
            className="btn btn-square btn-ghost lg:hidden"
          >
            <PiHamburger />
          </label>
          <h1 className="font-bold text-2xl">{title}</h1>
        </div>
        <div className="flex-1" />
        {balance && (
          <div className="h-full flex-0 text-nowrap rounded-xl bg-base-100 px-4 py-2 font-semibold text-base-content">
            <span className="inline">Balance: </span>
            {balance.toLocaleString("pl-PL", {
              style: "currency",
              currency: "EUR",
            })}
          </div>
        )}
      </div>
    </header>
  );
}
