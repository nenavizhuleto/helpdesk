<script>
	import {
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
	let items = data.branches;
</script>

<TableSearch
	placeholder="Search by maker name"
	hoverable={true}
	bind:inputValue={searchTerm}
>
	<TableHead>
		<TableHeadCell>ID</TableHeadCell>
		<TableHeadCell>Company ITN</TableHeadCell>
		<TableHeadCell>Name</TableHeadCell>
		<TableHeadCell>Description</TableHeadCell>
		<TableHeadCell>Address</TableHeadCell>
		<TableHeadCell>Contacts</TableHeadCell>
	</TableHead>
	<TableBody class="divide-y">
		{#if items}
			{#each items as item}
				<TableBodyRow>
					<TableBodyCell>{item.id}</TableBodyCell>
					<TableBodyCell>{item.company_id}</TableBodyCell>
					<TableBodyCell>{item.name}</TableBodyCell>
					<TableBodyCell>{item.description}</TableBodyCell>
					<TableBodyCell>{item.address}</TableBodyCell>
					<TableBodyCell>{item.contacts}</TableBodyCell>
				</TableBodyRow>
			{/each}
		{/if}
	</TableBody>
</TableSearch>
