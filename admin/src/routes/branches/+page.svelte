<script>
	import Table from "../Table.svelte";

	export let data;
	let branches = data.branches;
</script>

{#if branches}
	<Table items={branches} />
{/if}
