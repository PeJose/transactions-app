import type { InfiniteTransaction, Transaction } from "@/types";
import type { InfiniteData } from "@tanstack/react-query";
import TransactionsItem from "./TransactionItem";

type TransactionsListProps = {
  data: InfiniteData<InfiniteTransaction<Transaction>> | undefined;
  ref: React.Ref<HTMLTableRowElement>;
};

export default function TransactionsList({ data, ref }: TransactionsListProps) {
  return (
    <>
      {data?.pages.map((page) =>
        page.transactions.map((transaction, index) => {
          if (page.transactions.length === index + 1) {
            return (
              <tr ref={ref} key={transaction.id}>
                <TransactionsItem t={transaction} />
              </tr>
            );
          }
          return (
            <tr key={transaction.id}>
              <TransactionsItem t={transaction} />
            </tr>
          );
        }),
      )}
    </>
  );
}
