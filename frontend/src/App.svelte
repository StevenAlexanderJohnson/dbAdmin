<script>
  import Home from "./pages/Home.svelte";
  import Permissions from "./pages/PermissionsDashboard.svelte";
  import Register from "./pages/Register.svelte";
  import Router from "svelte-spa-router";
  import PermissionsRouter from "./routers/PermissionsRouter.svelte";

  const routes = {
    "/": Home,
    "/register": Register,
    "/permissions": PermissionsRouter,
    "/permissions/*": PermissionsRouter
  };
</script>

<main>
  <Router {routes} />
</main>

<style>
  :global(html, body, #app, #app main) {
    height: 100%;
    padding: 0;
    margin: 0;
    box-sizing: border-box;
    font-family: "Courier Prime";
    @apply text-text;
  }
</style>
