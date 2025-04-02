import { ReactNode } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <div className="drawer">
      <input id="my-drawer" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        <Header />
        {children}
      </div>
      <nav className="drawer-side">
        <label htmlFor="my-drawer" aria-label="close sidebar" className="drawer-overlay"></label>
        <ul className="menu bg-base-200 text-base-content min-h-full w-80 p-4">
          <Menu />
        </ul>
      </nav>
    </div>

  );
}

function Header() {
  return <header className="navbar bg-base-100 shadow-sm">
    <div className="flex-none">
      <label htmlFor="my-drawer" className="btn btn-ghost drawer-button">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" className="inline-block h-5 w-5 stroke-current"> <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16"></path> </svg>
      </label>
    </div>
    <div className="flex-1">
      <a className="btn btn-ghost text-xl">daisyUI</a>
    </div>
  </header>
}

function Menu() {
  const items = [
    { name: "Home", href: "/" },
    { name: "Companies", href: "/about" },
  ]

  return items.map((item) => (
    <li key={item.name}>
      <a href={item.href} className="hover:bg-base-100">
        {item.name}
      </a>
    </li>
  ))
}
