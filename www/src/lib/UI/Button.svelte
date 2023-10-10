<script>

    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher()

    export let type = "primary";
    export let disabled = false;

    const primaryStyles = "inline-flex  bg-primary  text-white py-4  px-8 rounded-xl items-center justify-center gap-4 font-semibold transition-all hover:bg-hover hover:cursor-pointer disabled:bg-disabled"
    const secondaryStyles = "inline-flex border-primary border text-primary py-4 px-8 rounded-xl items-center gap-4 font-semibold transition-all hover:border-hover hover:text-hover disabled:bg-disabled"
</script>

<button
	on:click={() => dispatch("click")}
	class={type === "primary" ? primaryStyles : secondaryStyles}
    type="button"
    {disabled}
>
	<slot />
</button>
