"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

interface QuerryWrapperProps {
  children: React.ReactNode;
}

const queryClient = new QueryClient();

const QueryWrapper = ({ children }: QuerryWrapperProps) => {
  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
};

export default QueryWrapper;
