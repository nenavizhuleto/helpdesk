<script>
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
	export let label;
	export let type;
	export let required;
	export let name;
    export let disabled = false;

    export let value = '';
</script>

{#if type == "textarea"}
	<div class="relative w-full mb-8">
		<textarea
            on:keyup={() => dispatch("change")}
			class="w-full h-96 px-4 py-5 border border-disabled rounded-xl outline-none text-black peer focus:border-primary resize-none"
			{required}
			{name}
            bind:value={value}
            {disabled}
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
        <!-- {type} -->
		<input
            type="text"
            on:keyup={() => dispatch("change")}
			{name}
			class="w-full px-4 py-5 border border-disabled rounded-xl outline-none text-black focus:border-primary"
			{required}
            {disabled}
            bind:value={value}
		/>
	</div>
{/if}
