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
	let items = data.users;

	function deleteUser(user_id) {
		fetch("http://127.0.0.1:3000/api/v3/users/" + user_id, {
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
		<TableHeadCell>UserID</TableHeadCell>
		<TableHeadCell>Name</TableHeadCell>
		<TableHeadCell>Phone</TableHeadCell>
		<TableHeadCell>Network</TableHeadCell>
		<TableHeadCell>Branch</TableHeadCell>
		<TableHeadCell>Company</TableHeadCell>
		<TableHeadCell>Devices</TableHeadCell>
		<TableHeadCell>Action</TableHeadCell>
	</TableHead>
	<TableBody class="divide-y">
		{#if items}
			{#each items as item}
				<TableBodyRow>
					<TableBodyCell>{item.id}</TableBodyCell>
					<TableBodyCell>{item.name}</TableBodyCell>
					<TableBodyCell>{item.phone}</TableBodyCell>
					<TableBodyCell>{item.network.netmask}</TableBodyCell>
					<TableBodyCell>{item.branch.name}</TableBodyCell>
					<TableBodyCell>{item.company.name}</TableBodyCell>
					<TableBodyCell>
						{#each item.devices as dev}
							<a>{dev.ip}</a>
						{/each}
					</TableBodyCell>
					<TableBodyCell>
						<Button on:click={() => deleteUser(item.id)} color="red"
							>Delete</Button
						>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		{/if}
	</TableBody>
</TableSearch>
