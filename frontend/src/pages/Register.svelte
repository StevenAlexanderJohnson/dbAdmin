<script>
    import { push, pop } from "svelte-spa-router";
    let serverInput = "";
    let databaseInput = "";
    let connectionStringInput = "";

    let usingConnectionString = false;

    document.addEventListener('submit', (e) => {
        e.preventDefault();
        console.log(serverInput)
        console.log(databaseInput)
        console.log(connectionStringInput)
        console.log(usingConnectionString)
    })
</script>

<form class="grid grid-rows-3 justify-center h-full">
    <div class="flex flex-col justify-center items-center self-end gap-4 ">
        <h1 class="text-6xl font-bold text-text">Register Database</h1>
        <label class="cursor-pointer select-none">
            <input
                type="checkbox"
                class="peer sr-only"
                bind:checked={usingConnectionString}
            />
            <div
                class="
                    relative p-4 inline-flex items-center gap-8 transition-all duration-300 rounded-full bg-gradient-to-br from-primary to-accent
                    after:absolute after:top-0 after:left-0 after:transition-all after:duration-300 after:h-full after:w-1/2 after:content-[''] after:bg-white after:rounded-full after:border-4 after:border-transparent
                    peer-checked:after:translate-x-full
                    "
            >
                <span class="z-10 p-2">Credential Form</span>
                <span class="z-10 p-2">Connection String</span>
            </div>
        </label>
    </div>
    <div class="flex flex-col justify-center items-center gap-4">
        {#if usingConnectionString}
            <input
                title="connection-string-input"
                type="text"
                class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
                placeholder="Connection String"
                required
                bind:value={connectionStringInput}
            />
        {:else}
            <input
                title="db-name-input"
                type="text"
                class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
                placeholder="Server Name"
                pattern="[A-Za-z0-9\.]*"
                required
                bind:value={serverInput}
            />
            <input
                title="db-name-input"
                type="text"
                class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
                placeholder="Database Name"
                pattern="[A-Za-z0-9\.]*"
                required
                bind:value={databaseInput}
            />
        {/if}
    </div>
    <div class="flex flex-col justify-center items-center gap-4 self-start">
        <div
            class="
            w-80 h-16 bg-gradient-to-r from-primary to-accent rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
            hover:before:animate-pulse
        "
        >
            <button
                class="h-full w-full rounded-full bg-background text-2xl outline-none flex justify-center items-center relative text-text hover:cursor-pointer"
                on:click={() => pop()}
            >
                back
            </button>
        </div>
        <div
            class="
            w-80 h-16 bg-gradient-to-r from-primary to-accent rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
            hover:before:animate-pulse
        "
        >
            <input
                type="submit"
                class="h-full w-full rounded-full bg-background text-2xl outline-none flex justify-center items-center relative text-text hover:cursor-pointer"
                value="register"
            />
        </div>
    </div>
</form>
