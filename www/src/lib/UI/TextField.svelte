<script>
	export let label;
	export let type;
	export let required;
	export let name;
</script>

{#if type == "textarea"}
	<div class="relative w-full mb-8">
		<textarea
			class="w-full h-96 px-4 py-5 border border-disabled rounded-xl outline-none text-black peer focus:border-primary resize-none"
			{required}
			{name}
		/>
		<span
			class="bg-white px-1 absolute -top-2 left-3 leading-4 text-disabled peer-focus:text-primary"
		>
			{label}
		</span>
	</div>
{:else}
	<div class="relative w-full focus-within:text-primary mb-8">
		<span class="bg-white text-disabled px-1 absolute -top-2 left-3 leading-4">
			{label}
		</span>
		<input
			{type}
			{name}
			class="w-full px-4 py-5 border border-disabled rounded-xl outline-none text-black focus:border-primary"
			{required}
		/>
	</div>
{/if}
