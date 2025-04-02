import { Container } from "@/components";
import { fetchApi } from "@/utils/api/server";
import Link from "next/link";

type Company = {
  id: string;
  name: string;
  ibans: string[];
  address: string;
};

function PaginatedTable({
  data,
  nextCursor,
  currentCursor,
}: {
  data: Company[];
  nextCursor: string | null;
  currentCursor: string | null;
}) {
  return (
    <div className="space-y-4">
      <table className="w-full border-collapse">
        <thead>
          <tr className="bg-base-100">
            <th className="p-2 text-left">ID</th>
            <th className="p-2 text-left">Name</th>
            <th className="p-2 text-left">Address</th>
            <th className="p-2 text-left">Ibans</th>
          </tr>
        </thead>
        <tbody>
          {data.length === 0 ? (
            <tr>
              <td colSpan={2} className="p-4 text-center">
                No companies found
              </td>
            </tr>
          ) : (
            data.map((item) => (
              <tr key={item.id} className="border-t">
                <td className="p-2">{item.id}</td>
                <td className="p-2">{item.name}</td>
                <td className="p-2">{item.address}</td>
                <td className="p-2">
                  {item.ibans.map((iban) => (
                    <Link key={iban} href={`/charts?companyIban=${iban}`}>
                      {iban}
                    </Link>
                  ))}
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>

      <div className="flex items-center justify-between">
        <Link
          href={
            currentCursor && currentCursor > "0"
              ? `/?afterId=${currentCursor}`
              : "/?afterId=0"
          }
          className={"btn btn-primary"}
        >
          Previous
        </Link>
        <Link
          href={nextCursor ? `/?afterId=${nextCursor}` : "#"}
          className={`btn btn-primary ${
            !nextCursor ? "pointer-events-none opacity-50" : ""
          }`}
        >
          Next
        </Link>
      </div>
    </div>
  );
}

const LIMIT = 10;

export default async function CompaniesPage({
  searchParams,
}: {
  searchParams: { afterId?: string };
}) {
  const { afterId = "0" } = await searchParams;

  const limit = LIMIT;

  // Fetch one more item than needed to check if there are more pages
  const companies = await fetchApi<Company[]>("/companies", {
    limit: (limit + 1).toString(),
    "after-id": afterId,
  });

  // Check if there are more items
  const hasMore = companies.length > limit;

  // Remove the extra item used for "hasMore" check
  const displayCompanies = hasMore ? companies.slice(0, limit) : companies;

  // Get the ID of the last item for the next cursor
  const nextCursor = hasMore ? companies[limit - 1].id : null;
  const prevCursor = Number(companies[0].id) - (limit + 1);

  return (
    <Container title="Companies">
      <PaginatedTable
        data={displayCompanies}
        nextCursor={nextCursor}
        currentCursor={prevCursor.toString()}
      />
    </Container>
  );
}
