<script>
    import Home from "../pages/Home.svelte";
    import Register from "../pages/Register.svelte";
    import Router from "svelte-spa-router";
    import PermissionsRouter from "../routers/PermissionsRouter.svelte";

    const routes = {
        "/": Home,
        "/register": Register,
        "/permissions": PermissionsRouter,
        "/permissions/*": PermissionsRouter,
    };
</script>

<Router {routes} />
