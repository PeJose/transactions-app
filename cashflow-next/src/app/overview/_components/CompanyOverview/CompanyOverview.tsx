"use client";
import type { BalanceInTime, Company, TransactionPerCountry } from "@/types";
import { useState } from "react";
import Transactions from "../Transactions/Transactions";
import TransactionBalanceInTime from "../TransactionsBalanceInTime";
import TransactionsMap from "../TransactionsMap";
import CompanyDetails from "./CompanyDetails";
import TabNavigation, { TABS } from "./TabNavigation";

type CompanyOverviewProps = {
  companyInfo: Company;
  balanceInTime: BalanceInTime[];
  transactionsPerCountry: TransactionPerCountry[];
  iban: string;
};

export default function CompanyOverview({
  balanceInTime,
  companyInfo,
  transactionsPerCountry,
  iban,
}: CompanyOverviewProps) {
  const [activeTab, setActiveTab] = useState("balance");

  return (
    <div>
      <CompanyDetails companyInfo={companyInfo} />
      <TabNavigation activeTab={activeTab} setActiveTab={setActiveTab} />
      <div className="mt-4">
        {activeTab === TABS.BALANCE && (
          <TransactionBalanceInTime data={balanceInTime} />
        )}
        {activeTab === TABS.MAP && (
          <TransactionsMap data={transactionsPerCountry} />
        )}
        {activeTab === TABS.SEPA && (
          <Transactions key={`${iban}-SEPA`} iban={iban} type="SEPA" />
        )}
        {activeTab === TABS.SWIFT && (
          <Transactions key={`${iban}-SWIFT`} iban={iban} type="SWIFT" />
        )}
      </div>
    </div>
  );
}
