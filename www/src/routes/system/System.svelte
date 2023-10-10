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
    let showTaskCreateModal = false;
    let taskCreateModalIsSubmittable = true;
    let showTaskDetails = false;
    let showMessageModal = false;

	onMount(async () => {
        await getTasks()
	});

    async function getTasks() {
		const res = await fetch("/api/tasks");
		tasks = await res.json();
    }

	async function handleSubmitModal() {
        taskCreateModalIsSubmittable = false;

		const res = await fetch("/api/tasks", {
			method: "POST",
			headers: {"Content-Type": "application/json"},
			body: JSON.stringify(newTask),
		});
		if (res.ok) {
			showTaskCreateModal = false;
			showMessageModal = true;
		} else {
			return 0;
		}
	}

    let task;
    $: showTaskDetails = task ? true : false
    function rowClicked(event) {
        const clickedTask = event.detail.task

        if (!task) {
            task = clickedTask
            return
        }

        if (task.id === clickedTask.id) {
            task = undefined
        } else {
            task = clickedTask
        }
    }

    let clearForm;

</script>

<div class="flex justify-between items-center mt-16 mb-10">
	<h1 class="text-4xl font-bold">Обращения</h1>
	<Button on:click={() => showTaskCreateModal = true}>Новое обращение</Button>
</div>
<TaskTable>
	{#if tasks}
		{#each tasks as task}
			<TaskTableRow {task} on:click={rowClicked}/>
		{/each}
	{/if}
</TaskTable>
<TaskDetails show={showTaskDetails} on:close={() => showTaskDetails = false} bind:task></TaskDetails>
<TaskCreateModal
	bind:task={newTask}
    bind:clearForm={clearForm}
	show={showTaskCreateModal}
    on:close={() => showTaskCreateModal = false}
    on:submit={handleSubmitModal}
    isSubmittable={taskCreateModalIsSubmittable}
/>
<MessageModal
	title="Спасибо"
	body="Ваше обращение отправлено в техническую поддержку, специалист свяжется с вами"
	show={showMessageModal}
    on:close={async () => { 
        showMessageModal = false
        taskCreateModalIsSubmittable = true;
        clearForm()
        await getTasks()
    }}
/>
