<script lang="ts">
	// --- Components ---
	import {
		Modal,
		Button,
		Alert,
		Label,
		Input,
		Textarea,
	} from "flowbite-svelte";

	// --- Icons ---
	import { ExclamationCircleOutline } from "flowbite-svelte-icons";

	let popup = false;
	let popupType: "exit" | "success" | "fail" = "exit";
	export let open = false;

	let name = "";
	let subject = "";

	function reset() {
		name = "";
		subject = "";
		popup = false;
		popupType = "exit";
		open = false;
	}

	type SubmitFunc = (name: string, subject: string) => Promise<string>; // returns created task's id
	export let onSubmit: SubmitFunc;
</script>

<Modal bind:open={popup} dismissable={false} size="xs" autoclose>
	<div class="text-center">
		{#if popupType == "exit"}
			<ExclamationCircleOutline
				class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
			/>
			<h3
				class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400"
			>
				Закрыть создание обращения? Введенные данные не сохранятся.
			</h3>
			<Button color="blue" class="mr-2" on:click={reset}>Закрыть</Button>
			<Button color="alternative" on:click={() => (open = true)}
				>Отмена</Button
			>
		{:else if popupType == "success"}
			<h3
				class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400"
			>
				Обращение успешно создано.
			</h3>
			<Button color="blue" class="mr-2" on:click={reset}>Закрыть</Button>
		{:else if popupType == "fail"}
			<ExclamationCircleOutline
				class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
			/>
			<h3
				class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400"
			>
				Не удалось создать обращение.
			</h3>
			<Button color="blue" class="mr-2" on:click={reset}>Закрыть</Button>
			<Button color="alternative" on:click={() => (open = true)}
				>Отмена</Button
			>
		{/if}
	</div>
</Modal>
<Modal
	bind:open
	on:close={() => (popup = true)}
	size="xs"
	autoclose={false}
	class="w-full"
>
	<form
		class="flex flex-col space-y-6"
		on:submit={async () => {
			let id = await onSubmit(name, subject);
			if (id === "") {
				popupType = "fail";
			} else {
				popupType = "success";
				setTimeout(() => {
					popup = false;
				}, 1000);
			}

			open = false;
		}}
	>
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
				bind:value={name}
				placeholder="Поломка принтера..."
				required
				size="md"
			/>
		</Label>
		<Label class="space-y-2">
			<span>Суть обращения</span>
			<Textarea
				name="subject"
				rows={4}
				placeholder="Опишите подробнее"
				bind:value={subject}
				required
			/>
		</Label>
		<Button color="blue" type="submit" class="w-full gap-4" size="lg"
			>Создать обращение</Button
		>
	</form>
</Modal>
