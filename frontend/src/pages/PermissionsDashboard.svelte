<script>
    import { GetUserPermissions } from "../lib/wailsjs/go/main/App";
    import { selectedConnection, selectedUser } from "../store";

    let selectedConnectionValue = "";
    selectedConnection.subscribe((value) => (selectedConnectionValue = value));
    let selectedUserValue = "";
    selectedUser.subscribe((value) => (selectedUserValue = value));

    let selectedDatabaseValue = "";
    let connectionArr = selectedConnectionValue.split(":");
    selectedDatabaseValue = connectionArr[connectionArr.length - 1];

    let userPermissions = [];
    $: GetUserPermissions(
        selectedConnectionValue,
        selectedUserValue,
        selectedDatabaseValue,
    )
        .then((data) => (userPermissions = JSON.parse(data)["Data"]))
        .catch((err) => console.error(err));

    let showModal = false;
    let addingPermission = false;
    let permissionName = "";
</script>

{#if showModal}
    <div
        class="fixed border border-text px-20 py-20 bg-background top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 "
    >
        <form on:submit|preventDefault={() => {}} class="flex flex-col justify-center gap-2">
            {#if addingPermission}
                <label for="add-permission">
                    Permission to add to {selectedUserValue}
                </label>
                <input
                    type="text"
                    class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
                    name="add-permission"
                    title="add-permission"
                    required
                    pattern="*"
                />
            {:else}
                <h2>
                    You are about to remove the {permissionName} permission from
                    {selectedUserValue}.
                </h2>
            {/if}
            <div class="flex justify-around">
                <button
                    on:click={() => {
                        showModal = false;
                    }}>Cancel</button
                >
                <button>Submit</button>
            </div>
        </form>
    </div>
{/if}
<div class="h-full relative">
    <div
        class="mt-10 grid grid-cols-4 w-1/2 min-w-max
               mx-auto border border-e-red-50
               p-5"
    >
        <span>User</span>
        <span>Role</span>
        <span>Database</span>
        <span>
            <button
                type="button"
                class="underline underline-offset-4 hover:text-primary"
                on:click={() => {
                    showModal = true;
                    addingPermission = true;
                }}
            >
                Grant Permssion
            </button>
        </span>
        {#each userPermissions as permission}
            <span>{permission.Name}</span>
            <span>{permission.PermissionName}</span>
            <span>{permission.ObjectName}</span>
            <div>
                <button
                    class="underline underline-offset-4 hover:text-primary"
                    on:click={() => {
                        showModal = true;
                        addingPermission = false;
                        permissionName = permission.PermissionName;
                    }}>Remove</button
                >
            </div>
        {/each}
    </div>
</div>
