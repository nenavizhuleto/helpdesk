<script lang="ts">
	import { getContext } from "svelte";
	import type { Identity } from "$lib/api/types";
	import { Button, Badge, Alert } from "flowbite-svelte";
  import { Modal, Label, Input, Checkbox } from 'flowbite-svelte';
	import {
    Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
    Textarea,
	} from "flowbite-svelte";
  import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

  let taskTextareaProps = {
    id: 'subject',
    name: 'subject',
    rows: 4,
    placeholder: 'Опишите подробнее'
  };

	import { PlusSolid } from "flowbite-svelte-icons";
	import type { PageData } from "./$types";

  let formData = {
    name: "",
    subject: ""
  }
  
  let formModal = false;
  let popupModal = false;


	export let data: PageData;
	let tasks = data.tasks;
</script>

<Modal bind:open={popupModal} dismissable={false} size="xs" autoclose>
  <div class="text-center">
    <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">Закрыть создание обращения? Введенные данные не сохранятся.</h3>
    <Button color="blue" class="mr-2" on:click={() => {
      formData = {
      name: "",
      subject: ""
    }
    }}>Закрыть</Button>
    <Button color="alternative" on:click={() => formModal = true}>Отмена</Button>
  </div>
</Modal>
<Modal bind:open={formModal} on:close={() => popupModal = true} size="xs" autoclose={false} class="w-full">
  <form class="flex flex-col space-y-6" action="#">
    <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">Создание обращения</h3>
    <Alert color="blue">
      Опишите в поле «Тема обращения» краткое содержание обращения, в поле «Суть обращения» опишите проблему подробно.
    </Alert>
    <Label class="space-y-2" >
      <span>Тема обращения</span>
      <Input type="text" name="name" bind:value={formData.name}  placeholder="Поломка принтера..." required size="md"/>
    </Label>
    <Label class="space-y-2">
      <span>Суть обращения</span>
      <Textarea {...taskTextareaProps} bind:value={formData.subject} required></Textarea>
    </Label>
    <Button color="blue" type="submit" class="w-full" size="lg">Создать обращение</Button>

  </form>
</Modal>

<div class="flex flex-col w-full px-6">
	<div class="flex w-full justify-between pb-8">

		<div class="text-2xl font-semibold">Мои обращения</div>
		
    <Button color="blue"  on:click={() => (formModal = true)}>
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
