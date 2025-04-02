import type { Company } from "@/types";
import Link from "next/link";

type CompanyTableProps = {
  data: Company[];
  nextCursor: number | null;
  currentCursor: string | null;
};

export default function CompanyTable({
  data,
  nextCursor,
  currentCursor,
}: CompanyTableProps) {
  return (
    <div className="space-y-4">
      <table className="table w-full border-collapse">
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
                    <Link
                      className="link inline"
                      key={iban}
                      href={`/overview?companyIban=${iban}`}
                    >
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
