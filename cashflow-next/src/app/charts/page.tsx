import ChartBalance from "./_components/ChartBalance";
import { decodeJwtFromCookies } from "@/utils/jwt";
import { Container } from "@/components";
import TransactionsMap from "./_components/TransactionsMap";
import {
  getBalanceInTime,
  getCompanyInfo,
  getTransactionsPerCountry,
} from "./actions";
import * as IBAN from "iban-ts";

export default async function ChartsPage({
  searchParams,
}: { searchParams: { companyIban: string } }) {
  const { companyIban } = await searchParams;
  const userIban = await decodeJwtFromCookies("iban");
  const iban = companyIban ?? userIban;

  if (!IBAN.isValid(iban)) {
    return (
      <Container title="Invalid iban">
        <h3 className="text-lg text-error">Iban you provided is not valid</h3>
      </Container>
    );
  }

  const companyInfo = await getCompanyInfo(iban);
  const balanceInTime = await getBalanceInTime(iban);
  const transactionsPerCountry = await getTransactionsPerCountry(iban);

  return (
    <Container title={companyInfo.name}>
      <div className="grid grid-cols-3 gap-4">
        <ChartBalance data={balanceInTime} />
        <TransactionsMap data={transactionsPerCountry} />
      </div>
    </Container>
  );
}
