import type { ReactNode } from "react";
import Menu from "./Menu";

type LayoutProps = {
  children: ReactNode;
};
export default function Layout({ children }: LayoutProps) {
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
