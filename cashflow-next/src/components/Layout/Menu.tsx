import Link from "next/link";
import { ReactNode } from "react";
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
  return (
    <ul className="menu bg-base-200 text-base-content min-h-full w-80 p-4">
      {MenuList.map((item) => (
        <li key={item.href}>
          <Link
            className="flex gap-2 items-center text-md btn btn-ghost mb-2 justify-start [&>svg]:w-6 [&>svg]:h-6"
            href={item.href}
          >
            {item.icon}
            {item.title}
          </Link>
        </li>
      ))}
    </ul>
  );
}
