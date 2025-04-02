import { Container } from "@/components";
import { decodeJwtFromCookies } from "@/utils/jwt";
import * as IBAN from "iban-ts";
import ChartBalance from "./_components/ChartBalance";
import TransactionsMap from "./_components/TransactionsMap";

import {
  getBalanceInTime,
  getCompanyInfo,
  getTransactionsPerCountry,
} from "@/actions";

export default async function ChartsPage({
  searchParams,
}: { searchParams: { companyIban: string } }) {
  const { companyIban } = await searchParams;
  const userIban = await decodeJwtFromCookies("iban");
  const iban = companyIban ?? userIban;

  if (!IBAN.isValid(iban)) {
    return (
      <Container title="Invalid iban">
        <h3 className="text-error text-lg">Iban you provided is not valid</h3>
      </Container>
    );
  }

  const companyInfo = await getCompanyInfo(companyIban);
  const balanceInTime = await getBalanceInTime(companyIban);
  const transactionsPerCountry = await getTransactionsPerCountry(companyIban);
  const companyTitle = `Companies - ${iban === userIban ? "Your company" : companyInfo.name}`;

  return (
    <Container title={companyTitle}>
      <div className="grid grid-cols-3 gap-4">
        <h2 className="text-lg col-span-3">{companyInfo.name}</h2>
        <ChartBalance data={balanceInTime} />
        <TransactionsMap data={transactionsPerCountry} />
      </div>
    </Container>
  );
}
