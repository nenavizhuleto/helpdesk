<script>
    import TaskTable from '$lib/TaskTable.svelte';
    import TaskTableRow from '$lib/TaskTableRow.svelte';
    import PrimaryButton from '$lib/UI/PrimaryButton.svelte';
    import {  onMount } from 'svelte';

    let tasks = [];
    export let identity;

    onMount(async () => {
        const res = await fetch("http://localhost:3000/api/tasks");
        tasks = await res.json();
    })
</script>

<div class="flex justify-between items-center mt-16 mb-10">
  <h1 class="text-4xl font-bold">Обращения</h1>
  <PrimaryButton>
    Новое обращение
  </PrimaryButton>
</div>
<TaskTable>
{#if tasks}
  {#each tasks as task}
  <TaskTableRow task={task} />
  {/each}
{/if}
</TaskTable>
