import type { ReactNode } from "react";
import Menu from "./Menu";

export default function Layout({
  children,
}: {
  children: ReactNode;
}) {
  return (
    <div className="drawer lg:drawer-open">
      <input id="my-drawer" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">{children}</div>
      <div className="drawer-side">
        <label
          htmlFor="my-drawer"
          aria-label="close sidebar"
          className="drawer-overlay"
        />
        <Menu />
      </div>
    </div>
  );
}
