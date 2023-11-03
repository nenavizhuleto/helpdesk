<script>
    //  --- Utils ---
    import { formatDate } from "$lib";
    // --- Components ---
    import { Badge } from "flowbite-svelte";
    export let task;
</script>

<div
    class="w-full max-w-xs h-screen flex flex-col border-l border-gray-300 bg-white transition-all duration-300"
  >
    <div class="px-6 py-6 border-b border-gray-300">
      <h2 class="text-lg font-bold">Подробности обращения</h2>
    </div>
    <div class="flex justify-between px-6 py-3 text-sm">
      <div class="text-zinc-500">Номер задачи:</div>
      <div class="">{task.id}</div>
    </div>
    <!-- Date -->
    <div class="flex justify-between px-6 py-3 text-sm">
      <div class="text-zinc-500">Дата постановки:</div>
      <div class="">{formatDate(task.created_at)}</div>
    </div>
    <div class="flex justify-between px-6 py-3 text-sm">
      <div class="text-zinc-500">Статус:</div>
      {#if task.status == "created"}
        <Badge large color="dark">Создано</Badge>
      {:else if task.status == "assigned"}
        <Badge large color="blue">Назначена</Badge>
      {:else if task.status == "accepted"}
        <Badge large color="yellow">В работе</Badge>
      {:else if task.status == "completed"}
        <Badge large color="green">Решено</Badge>
      {:else if task.status == "cancelled"}
        <Badge large color="red">Отклонено</Badge>
      {/if}
    </div>
    <div class="flex flex-col justify-between px-6 py-3 text-sm">
      <div class="text-zinc-500 mb-2">Тема:</div>
      <div class="text-base">{task.name}</div>
    </div>
    <div class="flex flex-col justify-between px-6 py-3 text-sm">
      <div class="text-zinc-500 mb-2">Описание:</div>
      <div class="text-base">{task.subject}</div>
    </div>
  </div>