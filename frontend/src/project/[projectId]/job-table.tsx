import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge.tsx";

// TODO: click on job to see job infos, like state history change, detailed parameters, etc.

const jobs = [
  {
    id: "0",
    param: {
      key1: "v0",
      key2: "v0",
      key3: "v0",
    },
    status: "PENDING",
  },
  {
    id: "1",
    param: {
      key1: "v1",
      key2: "v1",
      key3: "v1",
    },
    status: "RUNNING",
  },
  {
    id: "2",
    param: {
      key1: "v2",
      key2: "v2",
      key3: "v2",
    },
    name: "P0",
    status: "COMPLETED",
  },
  {
    id: "3",
    param: {
      key1: "v3",
      key2: "v3",
      key3: "v3",
    },
    status: "FAILED",
  },
  {
    id: "4",
    param: {
      key1: "v4",
      key2: "v4",
      key3: "v4",
    },
    status: "CANCELLED",
  },
];

function statusBadge(status: string) {
  switch (status) {
    case "PENDING":
      return <Badge variant="outline">PENDING</Badge>;
    case "RUNNING":
      return <Badge>RUNNING</Badge>;
    case "COMPLETED":
      return <Badge variant="success">COMPLETED</Badge>;
    case "FAILED":
      return <Badge variant="destructive">FAILED</Badge>;
    case "CANCELLED":
      return <Badge variant="secondary">CANCELLED</Badge>;
    default:
      return <Badge variant="outline">UNKNOWN</Badge>;
  }
}

export function JobTable() {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">Job</TableHead>
          <TableHead>Parameter</TableHead>
          <TableHead>Status</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {jobs.map((job) => (
          <TableRow key={job.id}>
            <TableCell className="font-medium">{job.id}</TableCell>
            <TableCell>{/*job.param*/}</TableCell>
            <TableCell>{statusBadge(job.status)}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
