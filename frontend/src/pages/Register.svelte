<script>
  import { push, pop } from "svelte-spa-router";
  import { RegisterDatabase } from "../lib/wailsjs/go/main/App";
  let serverInput = "";
  let databaseInput = "";
  let connectionStringInput = "";
  let usernameInput = "";
  let passwordInput = "";

  let usingConnectionString = false;

  document.addEventListener("submit", async  (e) => {
    e.preventDefault();
    console.log(serverInput);
    console.log(databaseInput);
    console.log(usernameInput);
    console.log(passwordInput);
    console.log(connectionStringInput);
    console.log(usingConnectionString);
    let output = await RegisterDatabase(
      serverInput,
      databaseInput,
      "mssql",
      usernameInput,
      passwordInput
    );
    console.log(output)
  });
</script>

<form class="grid grid-rows-3 gap-8 justify-center h-full">
  <div class="flex flex-col justify-center items-center self-end gap-4 px-4">
    <h1 class="text-6xl font-bold text-text text-center">Register Database</h1>
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
        <span class="z-10 p-2 font-bold text-black text-center"
          >Credential Form</span
        >
        <span class="z-10 p-2 font-bold text-black text-center"
          >Connection String</span
        >
      </div>
    </label>
  </div>
  <div class="flex flex-col justify-center items-center gap-4 mb-4">
    <label for="server-select" class="text-2xl font-bold text-text">
      Driver
    </label>
    <select
      title="server-select"
      class="w-80 h-16 text-center text-2xl outline-none bg-transparent z-10 border-b-2 border-primary text-text invalid:border-red-600 peer"
      value="mssql"
      required
    >
      <option value="mssql" class="bg-background outline-none border-none">
        MS SQL
      </option>
      <option value="mongo" class="bg-background outline-none border-none">
        MongoDB
      </option>
    </select>
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
        pattern="[A-Za-z0-9\.:]*"
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
      <input
        title="db-name-input"
        type="text"
        class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
        placeholder="Username"
        pattern="[A-Za-z0-9\.]*"
        required
        bind:value={usernameInput}
      />
      <input
        title="db-name-input"
        type="password"
        class="w-80 h-16 text-center text-2xl outline-none bg-transparent border-b-2 border-primary text-text invalid:border-red-600"
        placeholder="Password"
        pattern="[A-Za-z0-9\.]*"
        required
        bind:value={passwordInput}
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
