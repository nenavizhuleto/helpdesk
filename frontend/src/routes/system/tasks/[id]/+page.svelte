<!-- CHAT GOES HERE -->
<script lang="ts">
	// --- Utils ---

	// --- Types ---
	import type { PageData } from "./$types";

	// --- Components ---
	import { Button, Alert, Input, ButtonGroup } from "flowbite-svelte";
	import Message from "./Message.svelte";
	import Details from "./Details.svelte";
	import { getTaskComments, newTaskComment } from "$lib/api";
	import { scrollToBottom } from "$lib";
	import { afterUpdate } from "svelte";

	afterUpdate(() => {
		if (comments && sending) {
			scrollToBottom(chat);
			sending = false;
		}
	});

	export let data: PageData;
	let task = data.task;
	let comments = data.comments!;
	let sending = true;
	let message = "";
	let chat: Element;
	let messageInput: HTMLInputElement;

	async function sendMessage(content: string) {
		sending = true;
		const com = await newTaskComment(task.id, content);
		if (!com.status) {
			return;
		}
		const coms = await getTaskComments(task.id);
		if (!coms.status) {
			return;
		}
		comments = coms.data;
		message = "";
		messageInput.focus();
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
				class="flex flex-col flex-grow h-0 p-4 overflow-y-scroll scrollbar-hide"
				bind:this={chat}
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
						let:props
						defaultClass="outline-none w-full"
						id="input-addon"
						size="lg"
						type="text"
						placeholder="Ваше сообщение..."
						required
						name="message"
					>
						<input
							{...props}
							bind:value={message}
							bind:this={messageInput}
						/>
					</Input>
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
