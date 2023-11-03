<script lang="ts">
	import { goto } from "$app/navigation";
	import { register } from "$lib/api";

	async function onRegister(name: string, phone: string) {
		const response = await register(name, phone);
		if (!response.status) {
			return await goto("/data-not-found");
		}
		window.localStorage.setItem("token", response.data.token);
		return await goto("/system");
	}

	// --- Components ---
	import { Card, Button, Label, Input } from "flowbite-svelte";

	let firstName = "";
	let lastName = "";
	let phone = "";
</script>

<div class="flex gap-24 justify-center max-w-4xl mx-auto mt-56">
	<div class="self-center">
		<div class="text-4xl font-bold leading-10 mb-8">Вы почти у цели!</div>
		<div class="text-xl font-normal leading-tight">
			Чтобы начать пользоваться приложением, пожалуйста, пройдите быструю
			и простую процедуру регистрации
		</div>
	</div>
	<Card class="w-full max-w-md">
		<form class="flex flex-col space-y-6">
			<h3 class="text-xl font-medium text-gray-900 dark:text-white">
				Регистрация в системе
			</h3>
			<Label class="space-y-2">
				<span>Имя</span>
				<Input
					type="text"
					size="lg"
					name="firstName"
					placeholder="Нателла"
					bind:value={firstName}
					required
				/>
			</Label>
			<Label class="space-y-2">
				<span>Фамилия</span>
				<Input
					type="text"
					size="lg"
					name="lastName"
					placeholder="Наумова"
					bind:value={lastName}
					required
				/>
			</Label>
			<Label class="space-y-2">
				<span>Внутренний номер</span>
				<Input
					type="text"
					size="lg"
					name="phone"
					placeholder="1234"
					bind:value={phone}
					required
				/>
			</Label>
			<Button
				on:click={() => onRegister(firstName + " " + lastName, phone)}
				color="blue"
				class="w-full">Продолжить</Button
			>
		</form>
	</Card>
</div>
