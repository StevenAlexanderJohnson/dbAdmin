<script>
    import PermissionsDashboard from "../pages/PermissionsDashboard.svelte";
    import Router from "svelte-spa-router";
    import { pop } from "svelte-spa-router";
    import { GetConnections, GetUsers } from "../lib/wailsjs/go/main/App";
    import { selectedConnection, selectedUser } from "../store.js";
    import { onMount } from "svelte";

    let connections = [];
    let users = [];

    let selectedConnectionValue = "";
    selectedConnection.subscribe((value) => {
        selectedConnectionValue = value;
    });
    let selectedUserValue = "";
    selectedUser.subscribe((value) => {
        selectedUserValue = value;
    });

    $: GetUsers(selectedConnectionValue, "admin")
        .then((data) => users = JSON.parse(data)["Data"])
        .catch((err) => console.error(err));

    onMount(async () => {
        try {
            connections = await GetConnections();
        } catch (ex) {
            alert(ex);
        }
    });

    const prefix = "/permissions";
    const routes = {
        "/": PermissionsDashboard,
    };
</script>

<div>
    <div
        class="
        min-h-fit w-full bg-gradient-to-r from-primary to-accent pb-2 relative flex items-center
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
                    bind:value={selectedConnectionValue}
                    on:change={() => selectedConnection.set(selectedConnectionValue)}
                >
                    {#each connections as connection}
                        <option value={connection} class="rounded-full">
                            {connection}
                        </option>
                    {/each}
                </select>

                <label
                    for="user-selection"
                    class="px-3 text-background text-xl font-extrabold"
                >
                    User
                </label>
                <select
                    title="user-selection"
                    class="bg-background p-5 rounded-full"
                    bind:value={selectedUserValue}
                    on:change={() => selectedUser.set(selectedUserValue)}
                >
                    {#each users as user}
                        <option value={user.Name} class="rounded-full">
                            {user.Name}
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
