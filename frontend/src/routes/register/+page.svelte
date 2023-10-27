<script lang="ts">
	import { getContext } from "svelte";
	import type { PageData } from "./$types";
	import { enhance } from "$app/forms";

	import { Card, Button, Label, Input } from "flowbite-svelte";
	import type { Writable } from "svelte/store";
	import type { Identity } from "$lib/api/types";
	import type { ActionData } from "./$types";
	import { goto } from "$app/navigation";
	export let data: PageData;

	const identity: Writable<Identity> = getContext("identity");
	export let form: ActionData;
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
		<form
			class="flex flex-col space-y-6"
			action="?/register"
			method="post"
			use:enhance={() => {
				return async ({ result }) => {
					if (result.data.success) {
						goto("/system");
					}
				};
			}}
		>
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
					required
				/>
			</Label>
			<Button type="submit" color="blue" class="w-full">Продолжить</Button
			>
		</form>
	</Card>
</div>
