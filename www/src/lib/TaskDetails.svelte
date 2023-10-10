<script>
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
  export let show;
  export let task;
   import TaskStatus from "$lib/TaskStatus.svelte";
</script>
{#if show}
<div class="w-[400px] h-screen shadow-lg absolute flex flex-col top-0 right-0 bg-white transition-all duration-300">
  <div class="px-6 pt-8 pb-2">
		<!-- Close Button -->
		<button class="flex items-center cursor-pointer" on:click={() => dispatch("close")} >
			<svg
				class="shrink-0"
				xmlns="http://www.w3.org/2000/svg"
				width="16"
				height="16"
				viewBox="0 0 16 16"
				fill="none"
			>
				<path
					d="M12.6666 8H3.33325"
					stroke="#747380"
					stroke-width="1.5"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					d="M7.99992 12.6666L3.33325 7.99992L7.99992 3.33325"
					stroke="#747380"
					stroke-width="1.5"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
			</svg>
			<span class="ml-1 text-zinc-500 text-sm font-semibold leading-tight">
				Свернуть
			</span>
		</button>
	</div>
	<h2 class="text-2xl font-bold px-6 py-6">Подробности обращения</h2>
	<!-- ID -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Номер задачи:</div>
		<div class="">{ task?.humanNumber }</div>
	</div>
	<!-- Date -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Дата постановки:</div>
		<div class="">{ task?.timeCreated?.value }</div>
	</div>
	<!-- TODO: Add Close Date -->
	<!-- Status -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Статус:</div>

      <TaskStatus status={task?.status} />
	</div>
	<!-- Subject -->
	<div class="flex flex-col justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500 mb-2">Тема:</div>
		<div class="text-base">{ task?.name }</div>
	</div>
	<!-- Description -->
	<div class="flex flex-col justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500 mb-2">Описание:</div>
		<div class="text-base">
			{ task?.subject }
		</div>
	</div>
</div>
{/if}
