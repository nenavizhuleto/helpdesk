<script>
  export let status
</script>
{#if status == 'assigned'}
<div class="px-3 py-1.5 bg-neutral-200 rounded-md items-center flex">
  <div class="text-neutral-900 text-base font-medium leading-tight">
    Назначено
  </div>
</div>

{:else if status == 'accepted'}
<div class="px-3 py-1.5 bg-yellow-200 rounded-md items-center flex">
  <div class="text-yellow-900 text-base font-medium leading-tight">
    В работе
  </div>
</div>

{:else if status == 'completed'}
<div class="px-3 py-1.5 bg-green-200 rounded-md items-center flex">
  <div class="text-green-900 mr-2 text-base font-medium leading-tight">
    Решено 
  </div>
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="17" viewBox="0 0 16 17" fill="none">
    <path d="M13.3332 4.5L5.99984 11.8333L2.6665 8.5" stroke="#5C7345" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
  </svg>
</div>
{/if}
