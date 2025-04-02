import {
  getBalanceInTime,
  getCompanyInfo,
  getTransactionsPerCountry,
} from "@/actions";
import { Container } from "@/components";
import { decodeJwtFromCookies } from "@/utils/jwt";
import { isValid as isValidIban } from "iban-ts";
import CompanyOverview from "./_components/CompanyOverview/CompanyOverview";

type ChartsPageProps = {
  searchParams: {
    companyIban: string;
  };
};

export default async function ChartsPage({ searchParams }: ChartsPageProps) {
  const { companyIban } = await searchParams;
  const iban = companyIban ?? (await decodeJwtFromCookies("iban"))?.toString();
  const companyInfo = await getCompanyInfo(iban);
  const balanceInTime = await getBalanceInTime(iban);
  const transactionsPerCountry = await getTransactionsPerCountry(iban);

  return (
    <Container title="Company overview">
      {companyIban && !isValidIban(companyIban) ? (
        <h3 className="text-error text-lg">Iban you provided is not valid</h3>
      ) : (
        <CompanyOverview
          companyInfo={companyInfo}
          balanceInTime={balanceInTime}
          transactionsPerCountry={transactionsPerCountry}
          iban={iban}
        />
      )}
    </Container>
  );
}
