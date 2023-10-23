<!-- Button.svelte -->
<script lang="ts">
	import type {Property} from "$lib";
	import {createEventDispatcher} from "svelte";
	const dispatch = createEventDispatcher();

	const sizes = {
		small: "px-5 py-3 font-medium",
		medium: "px-6 py-4 font-semibold",
		large: "px-8 py-5 font-semibold",
	};

	const variations = {
		primary: `
      text-white
      bg-indigo-700
      
      hover:bg-indigo-600
      disabled:bg-indigo-200`,

		secondary: `
      bg-transparent
      text-indigo-700
      border-2
      border-indigo-700

      hover:text-indigo-600
      hover:border-indigo-600
      
      disabled:text-indigo-200
      disabled:border-indigo-200
      
      
      `,

		tetriary: `
      bg-transparent  
      text-indigo-700
      hover:bg-indigo-100
      disabled:text-indigo-200
      disabled:bg-transparent
      stroke-transparent
    `,
	};

	export let variation: Property<typeof variations> = "primary";
	export let size: Property<typeof sizes> = "medium";
	export let disabled = false;

	function handleClick() {}

	function message() {
		dispatch("message", {
			text: "Hello",
		});
	}
</script>

<button
	class="
    inline-flex h-fit gap-2 text-base leading-normal rounded-lg
    {sizes[size]}
    {variations[variation]}
  "
	{disabled}
	on:click={message}
>
	<slot name="left-icon" />
	<slot />
	<slot name="right-icon" />
</button>

<style>
	/* Применяем стили Tailwind CSS к кнопке */
</style>
