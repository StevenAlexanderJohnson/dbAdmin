<script>
  import AppRouter from "./routers/AppRouter.svelte";
</script>

<main>
  <AppRouter />
</main>

<style>
  :global(html, body, #app, #app main) {
    height: 100%;
    padding: 0;
    margin: 0;
    box-sizing: border-box;
    font-family: "Courier Prime";
    @apply text-text;
  }
</style>
