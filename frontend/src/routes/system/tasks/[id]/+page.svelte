<!-- CHAT GOES HERE -->
<script lang="ts">
	// --- Types ---
	import type { PageData } from "./$types";

	export let data: PageData;
	// TODO: Why we expect that task cannot be undefined?
	let task = data?.task!;
</script>

<h1>Task {task.name}</h1>
