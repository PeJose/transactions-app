import { Container } from "@/components";
import type { Company } from "@/types";
import { fetchApi } from "@/utils/api/server";
import CompanyTable from "./_components/CompanyTable";

const LIMIT = 10;

export default async function CompaniesPage({
  searchParams,
}: {
  searchParams: { afterId?: string };
}) {
  const { afterId = "0" } = await searchParams;

  const limit = LIMIT;

  const companies = await fetchApi<Company[]>("/companies", {
    limit: (limit + 1).toString(),
    "after-id": afterId,
  });
 
  const hasMore = companies.length > limit;

  const displayCompanies = hasMore ? companies.slice(0, limit) : companies;

  const nextCursor = hasMore ? companies[limit - 1].id : null;
  const prevCursor = Number(companies[0].id) - (limit + 1);

  return (
    <Container title="Companies">
      <CompanyTable
        data={displayCompanies}
        nextCursor={nextCursor}
        currentCursor={prevCursor.toString()}
      />
    </Container>
  );
}
