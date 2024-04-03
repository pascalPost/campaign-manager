"use client";

import { JobTable } from "./job-table.tsx";

export default function ProjectPage({
  params,
}: {
  params: { projectId: string };
}) {
  return (
    <>
      <h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        Project {params.projectId}
      </h1>
      <div className="mt-4">
        <JobTable />
      </div>
    </>
  );
}
