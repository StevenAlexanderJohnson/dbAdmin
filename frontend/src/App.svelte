<script>
  import Home from './pages/Home.svelte';
  import Register from './pages/Register.svelte';
  import Router from 'svelte-spa-router';
  import { Greet } from '../wailsjs/go/main/App';

  let name = "";
  let greeting = "";

  function greet() {
    Greet(name).then((result) => {
      greeting = result;
    });
  }

  const routes = {
    '/': Home,
    '/register': Register
  }
</script>

<main>
  <Router {routes} />
</main>

<style>
  :global(html, body, #app, #app main) {
    height: 100%;
    padding: 0;
    margin: 0;
    box-sizing: border-box;
    font-family: 'Courier Prime';
  }
</style>
