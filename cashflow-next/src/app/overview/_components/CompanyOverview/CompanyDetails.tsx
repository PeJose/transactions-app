"use client";
import type { Company } from "@/types";
import Link from "next/link";

type CompanyDetailsProps = {
  companyInfo: Company;
};

export default function CompanyDetails({ companyInfo }: CompanyDetailsProps) {
  return (
    <div>
      <h2 className="text-lg">Company: {companyInfo.name}</h2>
      <h3 className="text-md">Address: {companyInfo.address}</h3>
      <h3 className="mb-4 text-md">
        Ibans: [
        {companyInfo.ibans.map((iban) => (
          <Link
            className="link inline"
            key={iban}
            href={`/overview?companyIban=${iban}`}
          >
            {iban}
          </Link>
        ))}
        ]
      </h3>
    </div>
  );
}
