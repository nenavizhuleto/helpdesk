<script lang="ts">
	import { goto } from "$app/navigation";
	import { formatDate } from "$lib";
	import type { Task } from "$lib/api/types";
	import {
		Table,
		TableHead,
		TableHeadCell,
		TableBody,
		TableBodyRow,
		TableBodyCell,
	} from "flowbite-svelte";
	import TaskStatus from "./TaskStatus.svelte";

	export let tasks: Task[];
</script>

<Table hoverable shadow>
	<TableHead>
		<TableHeadCell>Номер</TableHeadCell>
		<TableHeadCell class="w-full">Тема</TableHeadCell>
		<TableHeadCell>Статус</TableHeadCell>
		<TableHeadCell>Создано</TableHeadCell>
	</TableHead>
	<TableBody>
		{#each tasks as task}
			<TableBodyRow
				class="cursor-pointer"
				on:click={() => {
					goto(`/system/tasks/${task.id}`);
				}}
			>
				<TableBodyCell>{task.id}</TableBodyCell>
				<TableBodyCell>{task.name}</TableBodyCell>
				<TableBodyCell>
					<TaskStatus status={task.status} />
				</TableBodyCell>
				<TableBodyCell>{formatDate(task.created_at)}</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
