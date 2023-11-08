<script lang="ts">
	import {
		Table,
		TableHead,
		TableHeadCell,
		TableBody,
		TableBodyRow,
		TableBodyCell,
	} from "flowbite-svelte";

	export let items: any[];
</script>

<Table hoverable={true}>
	<TableHead>
		{#each Object.keys(items[0]) as key}
			<TableHeadCell>{key}</TableHeadCell>
		{/each}
	</TableHead>
	<TableBody class="divide-y">
		{#each items as item}
			<TableBodyRow>
				{#each Object.keys(item) as key}
					<TableBodyCell>{item[key]}</TableBodyCell>
				{/each}
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
