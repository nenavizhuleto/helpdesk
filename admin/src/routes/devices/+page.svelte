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
	let items = data.devices;

	function deleteDevice(device_id) {
		fetch("http://127.0.0.1:3000/api/v3/devices/" + device_id, {
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
		<TableHeadCell>DeviceID</TableHeadCell>
		<TableHeadCell>Type</TableHeadCell>
		<TableHeadCell>Action</TableHeadCell>
	</TableHead>
	<TableBody class="divide-y">
		{#if items}
			{#each items as item}
				<TableBodyRow>
					<TableBodyCell>{item.ip}</TableBodyCell>
					<TableBodyCell>{item.type}</TableBodyCell>
					<TableBodyCell>
						<Button on:click={() => deleteDevice(item.ip)} color="red"
							>Delete</Button
						>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		{/if}
	</TableBody>
</TableSearch>
