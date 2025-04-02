import type { InfiniteTransaction, Transaction } from "@/types";
import { type NextRequest, NextResponse } from "next/server";

export const GET = async (request: NextRequest) => {
  const { searchParams } = new URL(request.url);
  const iban = searchParams.get("iban");
  const cursor = searchParams.get("cursor");
  const type = searchParams.get("type");
  const target = searchParams.get("target");

  if (!iban || !type) {
    return NextResponse.json(
      { error: "Missing required parameters: iban and cursor" },
      { status: 400 },
    );
  }

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/transactions/${type}?iban=${iban}&cursor=${cursor}&limit=10&target=${target}`,
    {
      credentials: "include",
      headers: {
        Cookie: request.headers.get("cookie") || "", 
      },
    },
  );

  if (!response.ok) {
    return NextResponse.json(
      { error: "Failed to fetch transactions" },
      { status: response.status },
    );
  }

  const data = (await response.json()).data;

  const mappedData: InfiniteTransaction<Transaction> = {
    hasNext: data.hasNext,
    lastCursor: data.lastCursor,
    // biome-ignore lint/suspicious/noExplicitAny: <explanation>
    transactions: data.transactions.map((transaction: any) => {
      if (type === "SEPA") {
        return {
          id: transaction.id,
          amount: transaction.amount,
          currency: transaction.currency,
          timestamp: transaction.timestamp,
          from: transaction.payer,
          to: transaction.receiver,
        };
      }

      return {
        id: transaction.id,
        amount: transaction.amount,
        currency: transaction.currency,
        timestamp: transaction.timestamp,
        from: transaction.sender,
        to: transaction.beneficiary,
      };
    }),
  };

  return NextResponse.json(mappedData);
};
