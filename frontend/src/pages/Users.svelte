<script>
    import { onMount } from "svelte";
    import { GetUsers } from "../lib/wailsjs/go/main/App.js";
    import {selectedConnection} from '../store.js';

    let users = [];
    onMount(async () => {
        try {
            let data = await GetUsers("localhost:27017:mongo", "admin");
            users = JSON.parse(data)['Data'];
        } catch (err) {
            console.error(err);
        }
    });

    let connection = '';
    selectedConnection.subscribe((value) => {
        connection = value;
    })
</script>

<div>
    <h1 class="text-4xl">Users</h1>
    <h2>{connection}</h2>
    <table>
        <thead>
            <tr>
                <th>User</th>
                <th>Role</th>
                <th>Database</th>
            </tr>
        </thead>
        <tbody>
            {#each users as user}
                <tr>
                    <th>{user.Name}</th>
                    <th>{user.PermissionName}</th>
                    <th>{user.ObjectName}</th>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
