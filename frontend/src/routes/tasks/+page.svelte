<script lang="ts">
	import { Button, Badge } from "flowbite-svelte";
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
	} from "flowbite-svelte";

	import { PlusSolid } from "flowbite-svelte-icons";
	import type { PageData } from "./$types";

	export let data: PageData;
	let identity = data.identity;
	let tasks = data.tasks;
</script>

<div class="flex flex-col w-full">
	<div class="flex p-4 my-4 w-full justify-between">
		Мои обращения
		<Button color="blue">
			<PlusSolid class="mr-2 w-3 h-3" />
			Новое обращение
		</Button>
	</div>
	<Table divClass="w-full" hoverable={true}>
		<TableHead>
			<TableHeadCell>Номер</TableHeadCell>
			<TableHeadCell class="w-full">Тема</TableHeadCell>
			<TableHeadCell>Статус</TableHeadCell>
			<TableHeadCell>Создано</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each tasks as task}
				<TableBodyRow>
					<TableBodyCell>{task.id}</TableBodyCell>
					<TableBodyCell>{task.name}</TableBodyCell>
					<TableBodyCell>
						{#if task.status == "created"}
							<Badge large color="dark">Создано</Badge>
						{:else if task.status == "assigned"}
							<Badge large color="blue">Назначена</Badge>
						{:else if task.status == "accepted"}
							<Badge large color="yellow">В работе</Badge>
						{:else if task.status == "completed"}
							<Badge large color="green">Решено</Badge>
						{:else if task.status == "cancelled"}
							<Badge large color="red">Отклонено</Badge>
						{/if}
					</TableBodyCell>
					<TableBodyCell>{task.created_at.toLocaleString()}</TableBodyCell>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
</div>
