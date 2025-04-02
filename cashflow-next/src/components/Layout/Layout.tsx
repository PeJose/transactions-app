import { ReactNode } from "react";
import Menu from "./Menu";
import Navbar from "./Navbar";

export default function Layout({
  children,
}: {
  children: ReactNode;
}) {
  return (
    <div className="drawer lg:drawer-open">
      <input id="my-drawer" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        <Navbar />
        {children}
      </div>
      <div className="drawer-side max-w-60">
        <label
          htmlFor="my-drawer"
          aria-label="close sidebar"
          className="drawer-overlay"
        ></label>
        <Menu />
      </div>
    </div>
  );
}
