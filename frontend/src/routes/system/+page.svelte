<script lang="ts">
	// --- Types ---
	import type { Task, TaskStatus } from "$lib/api/types";
	import { ImagePlaceholder } from "flowbite-svelte";
	import type { PageData } from "./$types";
	import TelegramWidget from "./TelegramWidget.svelte";

	export let data: PageData;
	let profile = data.profile;
	let tasks = data.tasks;

	function tasksByStatus(tasks: Task[], status: TaskStatus) {
		return tasks.filter((task) => task.status == status);
	}
</script>

<div class="p-6">
	<h2 class="text-black text-2xl font-semibold mb-6">Мой профиль</h2>
	{#if profile}
		<div class="justify-start items-center gap-6 flex mb-10">
			<div
				class="h-24 w-24 bg-gradient-to-tr from-indigo-700 to-violet-400 rounded-full"
			/>
			<div>
				<div class="text-gray-900 text-base font-medium mb-4">
					{profile.name}
				</div>
				<div class="text-gray-500 text-sm leading-tight">
					<div class="mb-2">{profile.company.name}</div>
					<div>{profile.branch.name}</div>
				</div>
			</div>
		</div>
	{:else}
		<ImagePlaceholder imgHeight={"24"} class="h-24 mb-10" />
	{/if}

	{#if tasks}
		<div class="mb-10">
			<h2 class="text-black text-2xl font-semibold mb-6">Обращения</h2>
			<!-- Cards List -->
			<div class="flex gap-4">
				<!-- Card -->
				<div
					class="px-10 pt-6 pb-8 bg-blue-50 rounded-xl flex-col items-center gap-4 flex"
				>
					<div class="text-blue-500 text-base font-bold">Создано</div>
					<div class="text-blue-900 text-5xl font-bold leading-10">
						{tasks.length}
					</div>
				</div>
				<!-- Card -->
				<div
					class="px-10 pt-6 pb-8 bg-green-50 rounded-xl flex-col items-center gap-4 flex"
				>
					<div class="text-green-500 text-base font-bold">Решено</div>
					<div class="text-green-900 text-5xl font-bold leading-10">
						{tasksByStatus(tasks, "done").length}
					</div>
				</div>
				<div
					class="px-10 pt-6 pb-8 bg-yellow-50 rounded-xl flex-col items-center gap-4 flex"
				>
					<div class="text-yellow-500 text-base font-bold">
						В работе
					</div>
					<div class="text-yellow-900 text-5xl font-bold leading-10">
						{tasksByStatus(tasks, "accepted").length}
					</div>
				</div>
			</div>
		</div>
	{/if}


	<div class="mt-4">
		<TelegramWidget />
	</div>
</div>
