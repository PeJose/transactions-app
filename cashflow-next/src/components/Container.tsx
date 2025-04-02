import { ReactNode } from "react";

type ContainerProps = {
  title?: string;
  children: ReactNode;
} & React.HTMLAttributes<HTMLDivElement>;

export default function Container({
  title,
  children,
  ...rest
}: ContainerProps) {
  return (
    <div className="container p-2 lg:p-8 mx-auto" {...rest}>
      {title && <h1 className="mb-4 pb-4 text-xl border-b">{title}</h1>}
      {children}
    </div>
  );
}
