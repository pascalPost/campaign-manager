import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./globals.css";
import { Route, Switch } from "wouter";
import DashboardPage from "./dashboard/page";
import CreateProjectPage from "@/project/create/page";
import ProjectPage from "@/project/[projectId]/page.tsx";
import { ThemeProvider } from "@/components/theme-provider.tsx";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from "@/components/ui/navigation-menu.tsx";
import { ModeToggle } from "@/components/mode-toggle.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <div className="m-4 flex-col">
        <div className="pb-2 pt-2">
          <NavigationMenu className="justify-end">
            <NavigationMenuList className="justify-end">
              <NavigationMenuItem>
                <ModeToggle />
              </NavigationMenuItem>
            </NavigationMenuList>
          </NavigationMenu>
        </div>

        <Switch>
          <Route path="/" component={DashboardPage} />
          <Route path="/project/create" component={CreateProjectPage} />
          <Route path="/project/:projectId" component={ProjectPage} />

          {/*/!* Default route in a switch *!/*/}
          <Route>404: No such page!</Route>
        </Switch>
      </div>
    </ThemeProvider>
  </StrictMode>,
);
