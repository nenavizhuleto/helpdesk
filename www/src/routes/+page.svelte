<script>
    import TaskTable from '$lib/TaskTable.svelte';
</script>
<div class="flex justify-between items-center mt-16 mb-10">
  <h1 class="text-4xl font-bold">Обращения</h1>
  <a class="inline-flex bg-primary text-white py-4 px-8 rounded-xl items-center gap-4 font-semibold transition-all hover:bg-hover">
    Новое обращение
  </a>
</div>
<TaskTable />
