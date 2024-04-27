<<<<<<< HEAD
<script>
    import { GetUserPermissions } from "../lib/wailsjs/go/main/App.js";
    import {selectedConnection, selectedUser} from '../store.js';

    let selectedConnectionValue = ""
    selectedConnection.subscribe((value) => selectedConnectionValue = value);
    let selectedUserValue = "";
    selectedUser.subscribe((value) => selectedUserValue = value);

    let userPermissions = [];
    $: GetUserPermissions(selectedConnectionValue, selectedUserValue, 'admin').then((data) => userPermissions = JSON.parse(data)['Data']).catch((err) => console.error(err));
</script>

<div>
    {#if userPermissions}
        {#each userPermissions as permission}
            <div>
                <span>{permission.Name}</span>
                <span>{permission.PermissionName}</span>
                <span>{permission.ObjectName}</span>
            </div>
        {/each}
    {:else}
        <p>Please wait while we look up that user's permissions.</p>
    {/if}
</div>
=======
<div>
    <h1 class="text-4xl">Users</h1>
</div>
>>>>>>> c4c1b26fd7d0469b658074c0b37f9421346e5d22
