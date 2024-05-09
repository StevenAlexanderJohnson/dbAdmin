<script>
    import { GetUserPermissions } from "../lib/wailsjs/go/main/App";
    import { selectedConnection, selectedUser } from "../store";

    let selectedConnectionValue = "";
    selectedConnection.subscribe((value) => (selectedConnectionValue = value));
    let selectedUserValue = "";
    selectedUser.subscribe((value) => (selectedUserValue = value));

    let selectedDatabaseValue = "";

    let userPermissions = [];
    $: GetUserPermissions(
        selectedConnectionValue,
        selectedUserValue,
        selectedDatabaseValue,
    )
        .then((data) => (userPermissions = JSON.parse(data)["Data"]))
        .catch((err) => console.error(err));
</script>

<div class="h-full">
    <div
        class="mt-10 grid grid-cols-4 w-1/2 min-w-max
               mx-auto border border-e-red-50
               p-5"
    >
        <span>User</span>
        <span>Role</span>
        <span>Database</span>
        <span>
            <button type="button" class="underline underline-offset-4 hover:text-primary">
                Grant Permssion
            </button>
        </span>
        {#each userPermissions as permission}
            <span>{permission.Name}</span>
            <span>{permission.PermissionName}</span>
            <span>{permission.ObjectName}</span>
            <div>
                <button class="underline underline-offset-4 hover:text-primary">Remove</button>
            </div>
        {/each}
    </div>
</div>