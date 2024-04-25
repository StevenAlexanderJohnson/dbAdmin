<script>
    import { GetUserPermissions } from "../lib/wailsjs/go/main/App.js";
    import {selectedConnection, selectedUser} from '../store.js';

    let selectedConnectionValue = ""
    selectedConnection.subscribe((value) => selectedConnectionValue = value);
    let selectedUserValue = "";
    selectedUser.subscribe((value) => selectedUserValue = value);

    let user = null;
    $: GetUserPermissions(selectedConnectionValue, selectedUserValue, 'admin').then((data) => user = JSON.parse(data)['Data']).catch((err) => console.error(err));
</script>

<div>
    {#if user}
        <p>{user.Name}</p>
    {:else}
        <p>Please wait while we look up that user's permissions.</p>
    {/if}
</div>
