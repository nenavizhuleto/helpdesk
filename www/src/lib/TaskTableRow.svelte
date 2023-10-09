<script>
	import TaskStatus from "./TaskStatus.svelte";
	export let task;
</script>

<div
	class="py-5 rounded-xl items-stretch inline-flex transition-all hover:bg-white hover:cursor-pointer text-neutral-800 text-base font-normal leading-tight"
>
	<!-- ID -->
	<div class="w-16 px-4 justify-start items-center gap-2.5 flex">
		{task.number}
	</div>

	<!-- Theme -->
	<div class="grow shrink self-stretch px-4 items-center flex">
		{task.name}
	</div>

	<!-- Status -->
	<div class="px-4 justify-end items-center gap-2.5 flex">
		<TaskStatus status={task.status} />
	</div>

	<div class="w-48 px-4 justify-end items-center text-right flex">
		<!-- Date -->
		{task.timeCreated.value}
	</div>
</div>
