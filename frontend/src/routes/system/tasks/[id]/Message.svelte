<script lang="ts">
	import type { Comment } from "$lib/api/types";
	// --- Utils ---
	import { formatDate } from "$lib";
	export let comment: Comment;
</script>

{#if comment.direction == "to"}
	<!-- START Message from HelpDesk -->
	<div class="flex w-full mt-2 space-x-3 max-w-xs">
		<!-- Message Avatar -->
		<div
			class="flex-shrink-0 flex items-center justify-center h-10 w-10 rounded-full bg-gradient-to-tr from-indigo-700 to-indigo-400"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="18"
				height="18"
				viewBox="0 0 18 18"
				fill="none"
			>
				<path
					d="M0.923618 6.77019C-0.307874 8.00168 -0.307873 9.99832 0.92362 11.2298L6.77019 17.0764C8.00168 18.3079 9.99832 18.3079 11.2298 17.0764L17.0764 11.2298C18.2474 10.0588 18.3049 8.19598 17.2489 6.95701C17.2079 6.90885 17.068 6.77988 17.0285 6.74046L16.8357 6.57136C16.5963 6.34042 16.2111 6.36137 15.9982 6.61691L10.795 11.8595L9.81792 12.7369C9.28422 13.2162 8.46778 13.1902 7.96557 12.6781L5.21967 9.87773C4.71785 9.36596 4.72188 8.5455 5.2287 8.03869C5.73905 7.52833 6.5665 7.52833 7.07685 8.03869L8.23981 9.20164C8.6407 9.60253 9.29089 9.60178 9.69086 9.19997L14.1062 4.76427C14.3334 4.46884 14.3127 4.0522 14.0574 3.7807L13.9184 3.63291C13.8831 3.59536 13.8514 3.55667 13.8233 3.51709L11.2298 0.923619C9.99832 -0.307873 8.00168 -0.307873 6.77019 0.923619L0.923618 6.77019Z"
					fill="white"
				/>
			</svg>
		</div>
		<div>
			<!-- Message Text-->
			<div class="bg-indigo-200 p-3 rounded-r-lg rounded-bl-lg text-sm">
				{comment.content}
			</div>
			<!-- Message Time-->
			<span class="text-xs text-gray-500 leading-none"
				>{formatDate(comment.created_at)}</span
			>
		</div>
	</div>
	<!-- END Message from HelpDesk -->
{:else}
	<!-- START Message from User -->
	<div class="flex w-full mt-2 space-x-3 max-w-xs ml-auto justify-end">
		<div>
			<!-- Message Text-->
			<div
				class="bg-blue-600 text-white p-3 rounded-l-lg rounded-br-lg text-sm"
			>
				{comment.content}
			</div>
			<!-- Message Time -->
			<span class="text-xs text-gray-500 leading-none"
				>{formatDate(comment.created_at)}</span
			>
		</div>
	</div>
	<!-- END Message from User -->
{/if}

