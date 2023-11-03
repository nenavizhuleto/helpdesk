<script lang="ts">
	import type { TaskStatus } from "$lib/api/types";
	import { Badge } from "flowbite-svelte";
	export let status: TaskStatus;
</script>

{#if status == "assigned"}
	<Badge large color="blue">Назначена</Badge>
{:else if status == "accepted"}
	<Badge large color="yellow">В работе</Badge>
{:else if status == "completed"}
	<Badge large color="green">Решено</Badge>
{:else if status == "cancelled"}
	<Badge large color="red">Отклонено</Badge>
{:else}
	<Badge large color="dark">Создано</Badge>
{/if}
