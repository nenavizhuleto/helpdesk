<script>
  import TextField from '$lib/UI/TextField.svelte'
  import Button from '$lib/UI/Button.svelte';
  import { createEventDispatcher } from 'svelte';
  export let show;
  export let isSubmittable;
  export let task = {
    "name": '',
    "subject": ''
  }

  export const clearForm = () => {
      task = {
        "name": '',
        "subject": ''
      }  
  }

  const dispatch = createEventDispatcher();


</script>
{#if show}
<!-- Taks Create Modal -->
<div
  class="absolute top-0 left-0 w-screen h-screen bg-black bg-opacity-50 flex justify-center items-center"
>
  <!-- Modal Body -->
  <div class="w-[506px] bg-white px-12 py-10 flex flex-col rounded-3xl gap-10">
    <div>
      <h2 class="text-2xl font-bold mb-8">Новое обращение</h2>
      <!-- Message -->
      <div
        class="w-full px-4 py-3 bg-orange-50 rounded-lg justify-start items-center gap-4 inline-flex"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
        >
          <path
            d="M12 22C17.5 22 22 17.5 22 12C22 6.5 17.5 2 12 2C6.5 2 2 6.5 2 12C2 17.5 6.5 22 12 22Z"
            stroke="#A87E28"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M12 8V13"
            stroke="#A87E28"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M11.9946 16H12.0036"
            stroke="#A87E28"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <div
          class="grow shrink basis-0 text-yellow-700 text-sm font-normal leading-tight"
        >
          Пожалуйста, постарайтесь сформулировать суть обращения как можно
          точнее, это значительно ускорит процесс решения проблемы
        </div>
      </div>
    </div>

    <form>
      <!-- Theme Field -->
      <div class="relative w-full mb-8">
        <TextField disabled={!isSubmittable} bind:value={task.name} label="Тема обращения" type="text" required ></TextField>
      </div>
      <div class="relative w-full mb-8">
        <TextField disabled={!isSubmittable} bind:value={task.subject} type="textarea" label="Суть обращения" required ></TextField>
      </div>
      <div class="flex justify-between">
        <Button on:click={() => dispatch("close")} type="secondary">
          Отмена
        </Button>
        <Button disabled={!isSubmittable} on:click={() => dispatch("submit")}>
          Отправить
        </Button>
      </div>
    </form>
  </div>
</div>
{/if}
