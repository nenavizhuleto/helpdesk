<script lang="ts">
	import { Button, Badge, Alert } from "flowbite-svelte";
	import { Modal, Label, Input, Checkbox } from "flowbite-svelte";
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Textarea,
	} from "flowbite-svelte";
	import { ExclamationCircleOutline } from "flowbite-svelte-icons";

	let taskTextareaProps = {
		id: "subject",
		name: "subject",
		rows: 4,
		placeholder: "Опишите подробнее",
	};

	import { PlusSolid } from "flowbite-svelte-icons";
	import type { PageData } from "./$types";
	import { redirect } from "@sveltejs/kit";
	import { goto } from "$app/navigation";

	let formData = {
		name: "",
		subject: "",
	};

	let formModal = false;
	let popupModal = false;

	export let data: PageData;
	let tasks = data.tasks!;
</script>

<Modal bind:open={popupModal} dismissable={false} size="xs" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline
			class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
		/>
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Закрыть создание обращения? Введенные данные не сохранятся.
		</h3>
		<Button
			color="blue"
			class="mr-2"
			on:click={() => {
				formData = {
					name: "",
					subject: "",
				};
			}}>Закрыть</Button
		>
		<Button color="alternative" on:click={() => (formModal = true)}
			>Отмена</Button
		>
	</div>
</Modal>
<Modal
	bind:open={formModal}
	on:close={() => (popupModal = true)}
	size="xs"
	autoclose={false}
	class="w-full"
>
	<form class="flex flex-col space-y-6" method="POST" action="?/task">
		<h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
			Создание обращения
		</h3>
		<Alert color="blue">
			Опишите в поле «Тема обращения» краткое содержание обращения, в поле
			«Суть обращения» опишите проблему подробно.
		</Alert>
		<Label class="space-y-2">
			<span>Тема обращения</span>
			<Input
				type="text"
				name="name"
				bind:value={formData.name}
				placeholder="Поломка принтера..."
				required
				size="md"
			/>
		</Label>
		<Label class="space-y-2">
			<span>Суть обращения</span>
			<Textarea
				{...taskTextareaProps}
				bind:value={formData.subject}
				required
			/>
		</Label>
		<Button color="blue" type="submit" class="w-full" size="lg"
			>Создать обращение</Button
		>
	</form>
</Modal>

<div class="flex flex-col w-full p-6">
	<div class="flex w-full justify-between pb-8">
		<div class="text-2xl font-semibold">Мои обращения</div>

		<Button color="blue" on:click={() => (formModal = true)}>
			<PlusSolid class="mr-2 w-3 h-3" />
			Новое обращение
		</Button>
	</div>
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
						console.log(task);
						goto(`/system/tasks/${task.id}`);
					}}
				>
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
					<TableBodyCell
						>{task.created_at.toLocaleString()}</TableBodyCell
					>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
{#if tasks.length == 0}
	<div class="w-[440px] mx-auto py-40">
		<div
			class="w-[268px] mb-10 mx-auto text-center text-zinc-500 text-xl font-medium"
		>
			У вас ещё нет обращений в техническую поддержку
		</div>

		<img src="/EmptyTickets.svg" alt="Empty Tasks" />
	</div>
{/if}
</div>

<div
	class="hidden w-[400px] h-screen shadow-lg absolute flex flex-col top-0 right-0 bg-white transition-all duration-300"
>
	<div class="px-6 pt-8 pb-2">
		<!-- Close Button -->
		<button class="flex items-center cursor-pointer">
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
			<span
				class="ml-1 text-zinc-500 text-sm font-semibold leading-tight"
			>
				Свернуть
			</span>
		</button>
	</div>
	<h2 class="text-2xl font-bold px-6 py-6">Подробности обращения</h2>
	<!-- ID -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Номер задачи:</div>
		<div class="">1312</div>
	</div>
	<!-- Date -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Дата постановки:</div>
		<div class="">Сегодня</div>
	</div>
	<!-- Status -->
	<div class="flex justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500">Статус:</div>
		<Badge large color="dark">Создано</Badge>
	</div>
	<!-- Subject -->
	<div class="flex flex-col justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500 mb-2">Тема:</div>
		<div class="text-base">Тема</div>
	</div>
	<!-- Description -->
	<div class="flex flex-col justify-between px-6 py-3 text-sm">
		<div class="text-zinc-500 mb-2">Описание:</div>
		<div class="text-base">Описание</div>
	</div>
</div>
