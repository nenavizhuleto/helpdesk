<script>
	import Table from "../Table.svelte";
	export let data;
	let companies = data.companies;
</script>

{#if companies}
	<Table items={companies} />
{/if}
