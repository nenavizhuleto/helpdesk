<script lang="ts">
	import { getContext } from "svelte";
	import type { PageData } from "./$types";

	import { Card, Button, Label, Input } from "flowbite-svelte";
	import type { Writable } from "svelte/store";
	import type { Identity } from "$lib/api/types";
	import type { ActionData } from "./$types";
	export let data: PageData;

	const identity: Writable<Identity> = getContext("identity");
	export let form: ActionData;
</script>

{#if $identity}
	<p>Already authorized</p>
{/if}

{#if form?.success}
	<p>Successfully registered</p>
{:else if !form?.success}
	<p>{form?.field} is empty</p>
{/if}

<Card class="w-full max-w-md">
	<form class="flex flex-col space-y-6" action="/" method="post">
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
		<Button type="submit" color="blue" class="w-full"
			>Login to your account</Button
		>
	</form>
</Card>

