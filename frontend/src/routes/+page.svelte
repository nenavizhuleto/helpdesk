<script lang="ts">
	import { Button } from "flowbite-svelte";
	import { goto } from "$app/navigation";
</script>

<div
	class="max-w-6xl mx-auto flex items-center justify-between my-60 flex-wrap"
>
	<div class="mb-20">
		<div class="mb-12">
			<h1 class="mb-8 font-bold text-4xl">
				Добро пожаловать <br />
				в help-desk!
			</h1>
			<p class="text-xl">
				Это приложение для решения ваших <br />
				технических проблем в офисе
			</p>
		</div>

		<Button color="blue" on:click={() => goto("/system")}>
			Начать работу
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="24"
				height="25"
				viewBox="0 0 24 25"
				fill="none"
			>
				<path
					d="M5 12.5H19"
					stroke="#F7F7F7"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
				<path
					d="M12 19.5L19 12.5L12 5.5"
					stroke="#F7F7F7"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				/>
			</svg>
		</Button>
	</div>
	<img src="/SupportAsset.svg" alt="#" />
</div>
