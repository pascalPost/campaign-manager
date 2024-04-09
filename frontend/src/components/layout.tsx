import { ReactNode } from "react";
import Header from "@/components/header";
import { Footer } from "@/components/footer.tsx";
import { Toaster } from "@/components/ui/sonner.tsx";

export const Layout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <Header />
      <div className="w-full flex-1 px-6">
        <aside className="fixed"></aside>
        <main className="py-6">{children}</main>
      </div>
      <Toaster richColors />
      <Footer />
    </>
  );
};
