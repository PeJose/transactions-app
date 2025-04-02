import type { Transaction } from "@/types";

type TransactionsItemProps = { t: Transaction };

export default function TransactionsItem({ t }: TransactionsItemProps) {
  return (
    <>
      <td>
        {new Date(t.timestamp).toLocaleDateString("en-US", {
          year: "2-digit",
          month: "2-digit",
          day: "2-digit",
        })}
      </td>
      <td>
        {t.amount.toLocaleString("en-US", {
          style: "currency",
          currency: t.currency,
        })}
      </td>
      <td>{t.from}</td>
      <td>{t.to}</td>
    </>
  );
}
