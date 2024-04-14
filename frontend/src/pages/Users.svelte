<script>
    import {GetUsers} from "../lib/wailsjs/go/main/App.js";
    const loadData = async () => {
        try {
        let data = await GetUsers("localhost:27017:a", "admin")
        console.log(data)
        }
        catch(err) {
            console.error(err)
        }
    }
    loadData();
</script>
<div>
    <h1 class="text-4xl">Users</h1>
</div>