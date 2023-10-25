<script lang="ts">
    import type { PageData } from './$types';
	  import { Button } from "flowbite-svelte";
    import {ArrowRightOutline} from "flowbite-svelte-icons"
    export let data: PageData;
</script>

<div class="max-w-6xl mx-auto flex items-center justify-between my-60 flex-wrap">
	<div class="mb-20 mr-20">
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

    <Button color="blue" size="xl" on:click={() => {}} href="/tasks">
      Начать работу <ArrowRightOutline class="ml-4"/>
    </Button>
		
	</div>
	<img src="/SupportAsset.svg" alt="#"/>
</div>
