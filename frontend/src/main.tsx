import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./globals.css";
import { Route, Switch } from "wouter";
import DashboardPage from "./dashboard/page";
import CreateProjectPage from "@/project/create/page";
import ProjectPage from "@/project/[projectId]/page.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Switch>
      <Route path="/" component={DashboardPage} />
      <Route path="/project/create" component={CreateProjectPage} />
      <Route path="/project/:projectId" component={ProjectPage} />

      {/*/!* Default route in a switch *!/*/}
      <Route>404: No such page!</Route>
    </Switch>
  </StrictMode>,
);
