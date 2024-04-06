import { ReactNode } from "react";
import Header from "@/components/header";
import { ScrollArea } from "@/components/ui/scroll-area.tsx";
import { Footer } from "@/components/footer.tsx";

export const Layout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <Header />
      <div className="w-full flex-1 px-6">
        <aside className="fixed">
          {/*<ScrollArea className="h-full py-6 pr-6 lg:py-8"></ScrollArea>*/}
        </aside>
        <main className="py-6">{children}</main>
      </div>
      <Footer />
    </>
  );
};
