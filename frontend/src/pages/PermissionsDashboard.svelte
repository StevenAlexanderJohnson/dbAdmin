<script>
    import { link } from "svelte-spa-router";
    import {location} from "svelte-spa-router";
</script>

<div class="h-full w-full flex flex-col">
    <div class="flex-1 flex flex-col gap-20 p-10">
        <a
            use:link={{ href: `${$location}/users`, disabled: false }}
            class="rounded-2xl bg-gradient-to-br from-accent to-text w-1/12 aspect-square flex justify-center items-center drop-shadow-md"
        >
            <span class="text-background text-center font-extrabold">Find Users</span>
        </a>
    </div>
</div>
