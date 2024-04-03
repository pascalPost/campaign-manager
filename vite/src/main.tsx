import React from "react";
import ReactDOM from "react-dom/client";
import "./globals.css";
import { Route, Switch } from "wouter";
import Dashboard from "./dashboard/page.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Switch>
      <Route path="/" component={Dashboard} />

      {/*<Route path="/users/:name">*/}
      {/*    {(params) => <>Hello, {params.name}!</>}*/}
      {/*</Route>*/}

      {/*/!* Default route in a switch *!/*/}
      <Route>404: No such page!</Route>
    </Switch>
  </React.StrictMode>,
);
