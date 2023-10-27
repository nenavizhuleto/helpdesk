<!-- CHAT GOES HERE -->
<script lang="ts">
	// --- Types ---
	import type { PageData } from "./$types";
  import { Button, Badge, Alert } from "flowbite-svelte";

	export let data: PageData;
	// TODO: Why we expect that task cannot be undefined?
	let task = data?.task!;
</script>
<div class="flex w-full">
  <div class="grow">
    <div class="px-6 py-6 border-b border-gray-300">
      <h2 class="text-lg font-bold text-center">{task.name}</h2>
    </div>
  </div>
  <div class=" w-[380px] h-screen   flex flex-col border-l border-gray-300 bg-white transition-all duration-300">
  <div class="px-6 py-6 border-b border-gray-300">
  	<h2 class="text-lg font-bold">Подробности обращения</h2>
  </div>
  	<div class="flex justify-between px-6 py-3 text-sm">
  		<div class="text-zinc-500">Номер задачи:</div>
  		<div class="">1312</div>
  	</div>
  	<!-- Date -->
  	<div class="flex justify-between px-6 py-3 text-sm">
  		<div class="text-zinc-500">Дата постановки:</div>
  		<div class="">Сегодня</div>
  	</div>
  	<div class="flex justify-between px-6 py-3 text-sm">
  		<div class="text-zinc-500">Статус:</div>
  		<Badge large color="dark">Создано</Badge>
  	</div>
  	<div class="flex flex-col justify-between px-6 py-3 text-sm">
  		<div class="text-zinc-500 mb-2">Тема:</div>
  		<div class="text-base">Тема</div>
  	</div>
  	<div class="flex flex-col justify-between px-6 py-3 text-sm">
  		<div class="text-zinc-500 mb-2">Описание:</div>
  		<div class="text-base">Описание</div>
  	</div>
  </div>
</div>
