<script>
	import TaskTable from "$lib/TaskTable.svelte";
	import TaskTableRow from "$lib/TaskTableRow.svelte";
	import Button from "$lib/UI/Button.svelte";
	import {onMount} from "svelte";
	import TaskCreateModal from "$lib/TaskCreateModal.svelte";
	import MessageModal from "$lib/MessageModal.svelte";
  import TaskDetails from "$lib/TaskDetails.svelte";
  
	let tasks = [];
	let newTask = {};
	let open = false;
	let openInfo = false;
  let openDetails = false;

	onMount(async () => {
		const res = await fetch("/api/tasks");
		tasks = await res.json();
	});

	function handleShowModal() {
		open = !open;
	}
	function handleShowInfoModal() {
		openInfo = !openInfo;
	}
	async function handleSubmitModal() {
		console.log(newTask);
		const res = await fetch("/api/tasks", {
			method: "POST",
			headers: {"Content-Type": "application/json"},
			body: JSON.stringify(newTask),
		});
		if (res.ok) {
			open = false;
			openInfo = true;
		} else {
			return 0;
		}
	}

  let task;

  async function handleShowDetails(_task, close = false) {
      if (close) {
        openDetails = false
        task = undefined
        return
      }
      if (task?.id === _task?.id) {
        openDetails = false
        task = undefined
      } else {
        openDetails = true
        // const res = await fetch("/api/tasks/" + _task.id)
        // task = await res.json()
        task = _task
      }
  }
</script>

<div class="flex justify-between items-center mt-16 mb-10">
	<h1 class="text-4xl font-bold">Обращения</h1>
	<Button on:click={handleShowModal}>Новое обращение</Button>
</div>
<TaskTable>
	{#if tasks}
		{#each tasks as task}
			<TaskTableRow {task} on:click={() => handleShowDetails(task)}/>
		{/each}
	{/if}
</TaskTable>
<TaskDetails open={openDetails} handleShowModal={handleShowDetails} bind:task>

</TaskDetails>
<TaskCreateModal
	bind:task={newTask}
	{open}
	{handleShowModal}
	handleSubmit={handleSubmitModal}
/>
<MessageModal
	title="Спасибо"
	body="Ваше обращение отправлено в техническую поддержку, специалист свяжется с вами"
	open={openInfo}
	handleShowModal={handleShowInfoModal}
/>
