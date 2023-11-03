<script lang="ts">
	import { getTelegram, newTelegram } from "$lib/api";
	import type { Telegram } from "$lib/api/types";
	import { Button, Modal } from "flowbite-svelte";
	import { onMount } from "svelte";

	let telegram: Telegram | undefined;

	onMount(async () => {
		const response = await getTelegram();
		if (!response.status) {
			telegram = undefined;
			return;
		}

		telegram = response.data;
	});

	async function createTelegram() {
		const res = await newTelegram();
		if (!res.status) {
			return;
		}

		telegram = res.data;
		show = true;
	}

	let show = false;
</script>

<Modal bind:open={show} autoclose>
	<div>{telegram?.pass}</div>
</Modal>

<div class="flex flex-col gap-2 p-4 bg-cyan-100">
	<div>Уведомления Телеграм</div>
	<a
		class="p-1 rounded bg-sky-400 hover:bg-sky-600 text-white"
		target="_blank"
		href="https://t.me/HDNotificator_bot">HDNotificator</a
	>
	{#if telegram}
		<Button on:click={() => (show = true)} color="blue"
			>Показать пароль</Button
		>
	{:else}
		<Button on:click={() => createTelegram()} color="blue"
			>Подключить</Button
		>
	{/if}
</div>
