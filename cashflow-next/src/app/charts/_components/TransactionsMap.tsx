"use client";
import type { TransactionPerCountry } from "@/types";
import Chart from "react-google-charts";

export default function TransactionsMap({
  data,
}: { data: TransactionPerCountry[] }) {
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
        region: "150", // Set region to Europe
        colorAxis: { colors: ["#a3c6ff", "#1d71fa"] }, // Tailwind-inspired gradient
      }}
    />
  );
}
