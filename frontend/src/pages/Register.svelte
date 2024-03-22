<script>
    import {push, pop} from 'svelte-spa-router';
    let serverInput = "";
    let databaseInput = "";
</script>

<div class="flex flex-col justify-center items-center h-full">
    <form class="flex flex-col justify-center items-center gap-10 h-full">
        <h1 class="text-6xl font-bold text-text">Register Database</h1>

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
        <div class="
            w-80 h-16 bg-gradient-to-r from-primary to-accent rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
            hover:before:animate-pulse
        ">
            <button
                class="h-full w-full rounded-full bg-background text-2xl outline-none flex justify-center items-center relative text-text hover:cursor-pointer"
                on:click={() => pop()}
            >
                back
            </button>
        </div>
        <div class="
            w-80 h-16 bg-gradient-to-r from-primary to-accent rounded-full flex items-center justify-center p-1 relative
            before:absolute before:top-0 before:left-0 before:content-[''] before:w-full before:h-full before:bg-gradient-to-r before:blur-xl
            hover:before:animate-pulse
        ">
            <input
                type="submit"
                class="h-full w-full rounded-full bg-background text-2xl outline-none flex justify-center items-center relative text-text hover:cursor-pointer"
                value="register"
            />
        </div>
    </form>
</div>