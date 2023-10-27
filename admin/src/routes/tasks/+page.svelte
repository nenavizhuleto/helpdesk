<script>
	import {
		Button,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		TableSearch,
	} from "flowbite-svelte";
	let searchTerm = "";
	export let data;
	let items = data.tasks;

	function deleteTask(task_id) {
		fetch("http://127.0.0.1:3000/api/v3/tasks/" + task_id, {
			method: "DELETE",
		});
	}
</script>

<TableSearch
	placeholder="Search by maker name"
	hoverable={true}
	bind:inputValue={searchTerm}
>
	<TableHead>
		<TableHeadCell>TaskID</TableHeadCell>
		<TableHeadCell>Name</TableHeadCell>
		<TableHeadCell>User</TableHeadCell>
		<TableHeadCell>Branch</TableHeadCell>
		<TableHeadCell>Company</TableHeadCell>
		<TableHeadCell>Action</TableHeadCell>
	</TableHead>
	<TableBody class="divide-y">
		{#if items}
			{#each items as item}
				<TableBodyRow>
					<TableBodyCell>{item.id}</TableBodyCell>
					<TableBodyCell>{item.name}</TableBodyCell>
					<TableBodyCell>{item.user.name}</TableBodyCell>
					<TableBodyCell>{item.branch.name}</TableBodyCell>
					<TableBodyCell>{item.company.name}</TableBodyCell>
					<TableBodyCell>
						<Button on:click={() => deleteTask(item.id)} color="red"
							>Delete</Button
						>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		{/if}
	</TableBody>
</TableSearch>
