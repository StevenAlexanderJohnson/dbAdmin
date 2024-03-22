<script>
    import {push} from 'svelte-spa-router';
    let dbName = "";
    let searching = false;

    let searchInput;

    function loadData() {
        searching = true
        setTimeout(() => {
            searching = false;
        }, 3000)
    }
</script>

<div class="flex flex-col justify-center items-center h-full">
    {#if searching}
        <span class="
            h-16 w-96 text-text text-2xl bg-background rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:from-primary before:to-accent before:blur-xl before:animate-ping
            "
        >
            One second while I look up your database info.
        </span>
    {:else}
    <form class="flex flex-col justify-center items-center gap-10 h-full">
        <h1 class="text-6xl font-bold text-text">Db Admin</h1>
        <div class="relative">
            <div class="relative">
                <select
                    title='server-select'
                    class="w-80 h-16 text-center text-2xl outline-none bg-transparent z-10 border-b-2 border-primary text-text invalid:border-red-600 peer"
                    value="-1"
                    required
                >
                    <option value="1" class="bg-background outline-none border-none">
                        JgCustFaceWebDev
                    </option>
                </select>
                <label
                    for='server-select'
                    class="hidden absolute -z-10 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-2xl text-text min-w-max text-center opacity-60 peer-invalid:block"
                >
                    Select Server
                </label>
            </div>
        </div>
        <input
            bind:this={searchInput}
            title="db-name-input"
            type="text"
            class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
            placeholder="Database Name"
            pattern="[A-Za-z0-9\.]*"
            required
            bind:value={dbName}
        />
        <div class="
            w-80 h-16 bg-gradient-to-r from-primary to-accent rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
            hover:before:animate-pulse
        ">
            <input
                type="submit"
                class="h-full w-full rounded-full bg-background text-2xl outline-none flex justify-center items-center relative text-text hover:cursor-pointer"
                value="search"
                on:click={() => loadData()}
            />
        </div>
        <button class="text-primary" on:click={() => push('/register')}>Need to connect to a server?</button>
    </form>
    {/if}
</div>