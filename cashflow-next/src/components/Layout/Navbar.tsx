import { getBalance } from "@/actions";
import type { ReactNode } from "react";
import { PiHamburger } from "react-icons/pi";

type NavbarProps = {
  title: ReactNode;
};

export default async function Navbar({ title }: NavbarProps) {
  const balance = await getBalance();
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
          <div className="flex-0 bg-base-100 rounded-xl px-4 py-2 h-full text-base-content font-semibold">
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
