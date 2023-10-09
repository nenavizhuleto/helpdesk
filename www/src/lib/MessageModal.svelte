<script>
  import Button from '$lib/UI/Button.svelte'
  export let title;
  export let body;
  export let open;
  export let handleShowModal;
</script>
{#if open}
<div class="absolute top-0 left-0 w-screen h-screen bg-black bg-opacity-50 flex justify-center items-center">
  <div class="w-[506px] bg-white px-12 py-10 flex flex-col rounded-3xl gap-8">
    <h2 class="text-2xl font-bold">{title}</h2>
    <p>{body}</p>
    <slot>
      <Button on:click={handleShowModal}>ОК</Button>
    </slot>
  </div>
</div>
{/if}