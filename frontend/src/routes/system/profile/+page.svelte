<script lang="ts">
	// --- Types ---
	import type { PageData } from "./$types";

	export let data: PageData;
	let identity = data.identity;
</script>

<div class="flex items-start gap-4 p-6">
	<div class="shadow-md p-6 bg-white rounded-xl">
		<div class="mb-2 text-xl font-medium">Компания:</div>
		<div class="text-black text-xl font-medium">
			{identity.company.name}
		</div>
	</div>
	<div class="shadow-md p-6 bg-white rounded-xl">
		<div class="mb-2 text-xl font-medium">Отдел:</div>
		<div class="text-black text-xl font-medium">{identity.branch.name}</div>
	</div>
	<div class="shadow-md p-6 bg-white rounded-xl">
		<div class="mb-2 text-xl font-medium">Пользователь:</div>
		<div class="text-black text-xl font-medium">{identity.name}</div>
	</div>
</div>

