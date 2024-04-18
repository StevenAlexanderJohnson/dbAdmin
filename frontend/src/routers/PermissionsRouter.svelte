<script>
    import PermissionsDashboard from "../pages/PermissionsDashboard.svelte";
    import Router from "svelte-spa-router";
    import { pop } from "svelte-spa-router";
    import { GetConnections } from "../lib/wailsjs/go/main/App";
    import {selectedConnection} from '../store.js';
    import Users from "../pages/Users.svelte";

    let connections = [];
    const onLoad = async () => {
        try {
            connections = await GetConnections();
        } catch (ex) {
            alert(ex);
        }
    };
    onLoad();
    let selectedValue = "";
    const prefix = "/permissions";
    const routes = {
        "/": PermissionsDashboard,
        "/users": Users,
    };
</script>

<div>
    <div
        class="
        h-20 w-full bg-gradient-to-r from-primary to-accent pb-2 relative flex items-center
        before:absolute before:top-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
        "
    >
        <div
            class="h-20 w-full pt-3 px-5 bg-gradient-to-r from-primary to-accent bg-opacity-30 flex items-center justify-between mb-2 relative"
        >
            <div>
                <label
                    for="connection-selection"
                    class="px-3 text-background text-xl font-extrabold"
                >
                    Connections
                </label>
                <select
                    title="connection-selection"
                    class="bg-background p-5 rounded-full"
                    bind:value={selectedValue}
                    on:change={() => selectedConnection.set(selectedValue)}
                >
                    {#each connections as connection}
                        <option value={connection} class="rounded-full">
                            {connection}
                        </option>
                    {/each}
                </select>
            </div>
            <button
                on:click={() => pop()}
                class="bg-background w-32 h-16 rounded-2xl">Back</button
            >
        </div>
    </div>
    <Router {routes} {prefix} />
</div>
