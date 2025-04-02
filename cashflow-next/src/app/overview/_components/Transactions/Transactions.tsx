"use client";

import type {
  InfiniteTransaction,
  Transaction,
  TransactionTarget,
  TransactionType,
} from "@/types";
import { useInfiniteQuery } from "@tanstack/react-query";
import { getCookie } from "cookies-next";
import { useEffect, useState } from "react";
import { useInView } from "react-intersection-observer";
import TransactionBottomIcon from "./TransactionBottomIcon";
import TransactionToggle from "./TransactionToggle";
import TransactionsList from "./TransactionsList";

type TransactionsProps = { iban: string; type: TransactionType };

async function getSepa(
  iban: string,
  cursor: string,
  type: TransactionType = "SEPA",
  target: TransactionTarget = "from",
): Promise<InfiniteTransaction<Transaction>> {
  const cookie = getCookie("jwt-token");
  const response = await fetch(
    `/api/transactions?iban=${iban}&cursor=${cursor}&type=${type}&target=${target}`,
    {
      credentials: "include",
      headers: new Headers(cookie ? { Cookie: cookie.toString() } : {}),
    },
  );
  if (!response.ok) {
    throw new Error("Failed to fetch SEPA transactions");
  }
  return await response.json();
}

export default function Transactions({ iban, type }: TransactionsProps) {
  const { ref, inView } = useInView();
  const [target, setTarget] = useState<TransactionTarget>("from");

  const {
    data,
    error,
    isLoading,
    hasNextPage,
    fetchNextPage,
    isSuccess,
    isFetchingNextPage,
  } = useInfiniteQuery({
    queryFn: ({ pageParam = "" }) => getSepa(iban, pageParam, type, target),
    queryKey: ["transactions", iban, type, target],
    initialPageParam: "",
    getNextPageParam: (lastPage) => {
      return !lastPage.hasNext ? undefined : lastPage.lastCursor;
    },
  });

  useEffect(() => {
    if (inView && hasNextPage) {
      fetchNextPage();
    }
  }, [hasNextPage, inView, fetchNextPage]);

  if (error)
    return (
      <div className="alert alert-error mt-10">
        <span>{`An error has occurred: ${error.message}`}</span>
      </div>
    );

  return (
    <div className="relative col-span-full mt-10 rounded-xl border border-base-300 shadow-lg">
      <TransactionToggle target={target} setTarget={setTarget} />
      <div className="max-h-[600px] overflow-y-scroll">
        <table className="table w-full">
          <thead className="sticky top-0 bg-white">
            <tr>
              <th style={{ width: 120 }}>Date</th>
              <th style={{ width: 180 }}>Amount</th>
              <th style={{ width: 300 }}>Payer</th>
              <th style={{ width: 300 }}>Receiver</th>
            </tr>
          </thead>
          <tbody>
            {isSuccess && data?.pages[0].transactions.length === 0 && (
              <tr>
                <td colSpan={4} className="text-center text-gray-500">
                  No transactions available.
                </td>
              </tr>
            )}
            {isSuccess && <TransactionsList data={data} ref={ref} />}
          </tbody>
        </table>
      </div>
      <TransactionBottomIcon
        inView={inView}
        isLoading={isLoading}
        isFetchingNextPage={isFetchingNextPage}
      />
    </div>
  );
}
