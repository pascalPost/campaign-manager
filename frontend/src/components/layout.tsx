import { ReactNode } from "react";
import Header from "@/components/header";
import { Footer } from "@/components/footer.tsx";
import { Toaster } from "@/components/ui/sonner.tsx";
import { Sidebar } from "@/components/sidebar.tsx";

export const Layout = ({ children }: { children: ReactNode }) => {
  return (
    <div className="flex h-screen flex-1 flex-col">
      <Header />
      <div className="flex h-full flex-row">
        <aside className="hidden h-full min-w-64 border-r px-6 py-6 lg:flex">
          <Sidebar />
        </aside>
        <div className="w-full px-6 py-6">
          <main className="flex w-full flex-col">{children}</main>
        </div>
      </div>
      <Toaster richColors />
      <Footer />
    </div>
  );
};
