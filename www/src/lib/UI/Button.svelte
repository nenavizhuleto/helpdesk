<script>

    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher()

    export let type = "primary";

    function handleOnClick() {
      dispatch("click")
    }

    const primaryStyles = "inline-flex  bg-primary  text-white py-4  px-8 rounded-xl items-center justify-center gap-4 font-semibold transition-all hover:bg-hover hover:cursor-pointer"
    const secondaryStyles = "inline-flex border-primary border text-primary py-4 px-8 rounded-xl items-center gap-4 font-semibold transition-all hover:border-hover hover:text-hover"
</script>

<button
	on:click={handleOnClick}
	class={type === "primary" ? primaryStyles : secondaryStyles}
>
	<slot />
</button>
