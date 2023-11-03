<!-- CHAT GOES HERE -->
<script lang="ts">
	// --- Utils ---
	import { scrollToBottom } from "$lib";
	import { enhance } from "$app/forms";

	// --- Types ---
	import type { PageData } from "./$types";

	// --- Components ---
	import { Button, Alert, Input, ButtonGroup } from "flowbite-svelte";
	import Message from "./Message.svelte";
	import Details from "./Details.svelte";
	import { newTaskComment } from "$lib/api";

	export let data: PageData;
	// TODO: Why we expect that task cannot be undefined?
	let task = data.task;
	let comments = task.comments ?? [];
	let sending = false;
	let message = "";

	async function sendMessage(content: string) {
		sending = true;
		const response = await newTaskComment(task.id, content);
		if (!response.status) {
			// do something
			return;
		}
		message = "";
		sending = false;
	}
</script>

<div class="flex w-full">
	<!-- Chat Container -->
	<div class="grow h-screen flex flex-col">
		<!-- Chat Header -->
		<div class="px-6 py-6 border-b border-gray-300">
			<h2 class="text-lg font-bold text-center">{task.name}</h2>
		</div>

		<!-- Chat Body -->
		<div class="flex flex-col flex-grow w-full bg-white overflow-hidden">
			<!-- Chat Flow -->
			<div
				class="flex flex-col flex-grow h-0 p-4 overflow-auto scrollbar-hide"
			>
				<div class="flex flex-col gap-4 mb-6">
					<Alert color="blue">
						<div class="font-medium mb-2">
							Добро пожаловать в чат обращения!
						</div>
						Здесь вы можете уточнять подробности задачи и получать обратную
						связь от специалистов
					</Alert>
					<Alert color="yellow">
						<div class="font-medium mb-2">Важно!</div>
						Старайтесь отправлять только важную и сформулированную информацию
						по теме обращения.
					</Alert>
				</div>
				{#if comments}
					{#each comments as comment}
						<Message {comment} />
					{/each}
				{/if}
			</div>
			<!-- Chat Input -->
			<form
				class="p-6 border-t border-gray-300"
				on:submit|preventDefault={() => sendMessage(message)}
			>
				<ButtonGroup class="w-full">
					<Input
						defaultClass="outline-none w-full"
						id="input-addon"
						size="lg"
						type="text"
						placeholder="Ваше сообщение..."
						name="message"
						bind:value={message}
					/>
					<Button type="submit" color="blue">Отправить</Button>
				</ButtonGroup>
			</form>
		</div>
	</div>
	<Details {task} />
</div>

<style>
	.scrollbar-hide::-webkit-scrollbar {
		display: none;
	}

	/* For IE, Edge and Firefox */
	.scrollbar-hide {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
</style>
