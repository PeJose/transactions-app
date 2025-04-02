"use client";
import type { TransactionPerCountry } from "@/types";
import Chart from "react-google-charts";

type TransactionsMapProps = { data: TransactionPerCountry[] };

export default function TransactionsMap({ data }: TransactionsMapProps) {
  const chartData = [
    ["Country", "Occurances", "Amount"],
    ...data.map(({ country, occurrances, amount }) => [
      country,
      occurrances,
      amount,
    ]),
  ];

  return (
    <Chart
      chartType="GeoChart"
      width="100%"
      height="100%"
      data={chartData}
      options={{
        region: "150", 
        colorAxis: { colors: ["#a3c6ff", "#1d71fa"] }, 
      }}
    />
  );
}
