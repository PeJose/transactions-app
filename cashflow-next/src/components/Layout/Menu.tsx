"use client";
import { getPathname } from "@/utils/navigation";
import classnames from "classnames";
import Link from "next/link";
import { usePathname } from "next/navigation";
import type { ReactNode } from "react";
import { PiHouse, PiListMagnifyingGlass, PiMoney } from "react-icons/pi";

type MenuItem = {
  title: string;
  icon: ReactNode;
  href: string;
};

const MenuList: MenuItem[] = [
  {
    href: "/",
    title: "Dashboard",
    icon: <PiHouse />,
  },
  {
    href: "/charts",
    title: "Charts",
    icon: <PiMoney />,
  },
  {
    href: "/transactions",
    title: "Transactions",
    icon: <PiListMagnifyingGlass />,
  },
];

export default function Menu() {
  const path = usePathname();
  console.log(path);
  return (
    <div className="h-full p-2">
      <div className="menu h-full rounded-xl bg-base-200 text-base-content">
        <Link href="/" className="btn btn-ghost mb-2 text-xl">
          CashFlow
        </Link>
        <ul>
          {MenuList.map((item) => {
            const itemApperance = classnames(
              "btn btn-ghost mb-2 flex items-center justify-start gap-2 text-md [&>svg]:h-6 [&>svg]:w-6",
              { "btn-active": path === item.href },
            );
            return (
              <li key={item.href}>
                <Link className={itemApperance} href={item.href}>
                  {item.icon}
                  {item.title}
                </Link>
              </li>
            );
          })}
        </ul>
      </div>
    </div>
  );
}
