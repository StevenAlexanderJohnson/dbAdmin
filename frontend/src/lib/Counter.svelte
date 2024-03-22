<script>
  let count = 0
  const increment = () => {
    count += 1
  }
</script>

<button class="px-14 py-5 bg-orange-500/75 rounded-full focus:ring focus:ring-orange-500" on:click={increment}>
  Clicks: {count}
</button>
