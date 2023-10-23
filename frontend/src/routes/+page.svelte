<script lang="ts">
	import type { PageData } from "./$types";

	export let data: PageData;
	let identity = data.identity;
</script>

<div class="flex flex-col m-4">
	<p>Company: {identity.company.name}</p>
	<p>Branch: {identity.branch.name}</p>
	<p>Username: {identity.user.name}</p>
</div>
