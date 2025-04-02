import type { ReactNode } from "react";
import Navbar from "./Layout/Navbar";

type ContainerProps = {
  title: ReactNode;
  children: ReactNode;
} & React.HTMLAttributes<HTMLDivElement>;

export default function Container({
  title,
  children,
  ...rest
}: ContainerProps) {
  return (
    <>
      <Navbar title={title} />
      <div className="container mx-auto p-2 lg:p-8" {...rest}>
        {children}
      </div>
    </>
  );
}
