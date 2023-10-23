<script lang="ts">
	import type {Property} from "$lib";

	const types = {};
</script>

<div class="relative w-full">
	<input
		type="email"
		class="
    peer
    w-full
    px-4
    py-5
    border-2
    border-gray-400
    rounded-lg
    outline-none
    text-black
    focus:invalid:border-red-500
    focus:border-indigo-700
    "
	/>
	<span
		class="
  bg-white
    font-medium
    px-1
    absolute
    -top-[0.565rem]
    left-3
    text-sm
    leading-tight
    text-gray-400
    peer-focus:text-indigo-700
    peer-focus:peer-invalid:text-red-500
    "
	>
		Label
	</span>
	<span
		class="text-red-500
    px-1 absolute
    -bottom-6
    left-3 text-sm leading-tight invisible peer-invalid:visible"
	>
		Supporting Text
	</span>
</div>
