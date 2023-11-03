<script lang="ts">
	// --- Utils ---
	import { getTasks, newTask } from "$lib/api";

	// --- Types ---
	import type { PageData } from "./$types";

	// --- Components ---
	import { Button, Tabs, TabItem } from "flowbite-svelte";
	import UserTaskTable from "./UserTaskTable.svelte";
	import BranchTaskTable from "./BranchTaskTable.svelte";
	import CompanyTaskTable from "./CompanyTaskTable.svelte";
	import NewTaskModal from "./NewTaskModal.svelte";

	// --- Icons ---
	import { PlusSolid } from "flowbite-svelte-icons";

	export let data: PageData;
	let tasks = data.tasks;
	let taskFilter: "branch" | "company" | undefined;

	$: fetchTasks(taskFilter);

	async function fetchTasks(filter?: "branch" | "company") {
		const res = await getTasks(filter);
		if (!res.status) {
			return;
		}
		tasks = res.data;
	}

	let taskModal = false;
	async function createTask(name: string, subject: string) {
		const task = await newTask(name, subject);
		if (!task.status) {
			return "";
		}
		let task_id = task.data;
		fetchTasks(taskFilter);
		return task_id;
	}
</script>

<NewTaskModal bind:open={taskModal} onSubmit={createTask} />

<div class="flex flex-col w-full p-6">
	<div class="flex w-full justify-between pb-8">
		<div class="text-2xl font-semibold">Обращения</div>

		<Button color="blue" on:click={() => (taskModal = true)}>
			<PlusSolid class="mr-2 w-3 h-3" />
			Новое обращение
		</Button>
	</div>
	<Tabs
		contentClass="bg-white pt-4"
		activeClasses="p-4 text-blue-600 bg-gray-100 rounded-t-lg dark:bg-gray-800 dark:text-blue-500"
	>
		{#if tasks}
			<TabItem open title="Мои" on:click={() => (taskFilter = undefined)}>
				<UserTaskTable {tasks} />
			</TabItem>
			<TabItem title="Отдел" on:click={() => (taskFilter = "branch")}>
				<BranchTaskTable {tasks} />
			</TabItem>
			<TabItem title="Компания" on:click={() => (taskFilter = "company")}>
				<CompanyTaskTable {tasks} />
			</TabItem>
		{/if}
	</Tabs>
	{#if tasks?.length == 0}
		<div class="w-[440px] mx-auto py-40">
			<div
				class="w-[268px] mb-10 mx-auto text-center text-zinc-500 text-xl font-medium"
			>
				Здесь ещё нет обращений в техническую поддержку
			</div>

			<img src="/EmptyTickets.svg" alt="Empty Tasks" />
		</div>
	{/if}
</div>
