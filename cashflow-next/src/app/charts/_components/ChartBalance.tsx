"use client";
import type { BalanceInTime } from "@/types";
import {
  Area,
  AreaChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

type ChartBalanceProps = {
  data: BalanceInTime[];
};

const gradientOffset = (data: BalanceInTime[]) => {
  const dataMax = Math.max(...data.map((i) => i.amount));
  const dataMin = Math.min(...data.map((i) => i.amount));

  if (dataMax <= 0) {
    return 0;
  }
  if (dataMin >= 0) {
    return 1;
  }

  return dataMax / (dataMax - dataMin);
};

export default function ChartBalance({ data }: ChartBalanceProps) {
  const off = gradientOffset(data);
  console.log(off);
  return (
    <ResponsiveContainer
      width="100%"
      height="100%"
      className="min-h-[500px] col-span-2"
    >
      <AreaChart
        width={500}
        height={300}
        data={data}
        margin={{
          top: 5,
          right: 30,
          left: 20,
          bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis
          dataKey="date"
          tickFormatter={(val) => new Date(val).toDateString()}
          tickCount={20}
        />
        <YAxis
          dataKey="amount"
          tickFormatter={(val) =>
            val.toLocaleString("en-US", { style: "currency", currency: "EUR" })
          }
          width={100}
          domain={["dataMin", "dataMax"]}
        />
        <Tooltip
          labelFormatter={(val) => new Date(val).toDateString()}
          formatter={(value: number) =>
            value.toLocaleString("en-US", {
              style: "currency",
              currency: "EUR",
            })
          }
        />
        <defs>
          <linearGradient id="splitColor" x1="0" y1="0" x2="0" y2="1">
            <stop offset={off} stopColor="rgb(34 197 94)" stopOpacity={1} />{" "}
            {/* Tailwind green-500 */}
            <stop offset={off} stopColor="rgb(239 68 68)" stopOpacity={1} />{" "}
            {/* Tailwind red-500 */}
          </linearGradient>
        </defs>
        <Area
          type="linear"
          dataKey="amount"
          stroke="#000"
          fill="url(#splitColor)"
          name="Amount(EUR)"
        />
      </AreaChart>
    </ResponsiveContainer>
  );
}
