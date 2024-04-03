import { Link } from "wouter";
import { Button } from "@/components/ui/button.tsx";
import { ProjectTable } from "./project-table.tsx";

export default function DashboardPage() {
  return (
    <>
      <h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        Dashboard
      </h1>
      <div className="mt-4 flex gap-5">
        <Button className="w-full" asChild>
          <Link href="/project/create">Create New Project</Link>
        </Button>
      </div>
      <div className="mt-4">
        <ProjectTable />
      </div>
    </>
  );
}
