<script>
    import { GetUsers } from "../lib/wailsjs/go/main/App.js";
    let users = [];
    const loadData = async () => {
        try {
            let data = await GetUsers("localhost:27017:mongo", "admin");
            users = JSON.parse(data)['Data'];
        } catch (err) {
            console.error(err);
        }
    };
    loadData();
</script>

<div>
    <h1 class="text-4xl">Users</h1>
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
