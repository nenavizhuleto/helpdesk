<script lang="ts">
  import { Drawer, Button, CloseButton } from 'flowbite-svelte';
  import { InfoCircleSolid, ArrowRightOutline } from 'flowbite-svelte-icons';
  import { sineIn } from 'svelte/easing';

  let hidden6 = true;
  let transitionParamsRight = {
    x: 320,
    duration: 200,
    easing: sineIn
  };

</script>

<div class="text-center">
  <Button on:click={() => (hidden6 = false)}>Show drawer</Button>
</div>

<Drawer placement="right" transitionType="fly" transitionParams={transitionParamsRight} bind:hidden={hidden6} id="sidebar6">
  <div class="flex items-center">
    <h5 id="drawer-label" class="inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400">
      <InfoCircleSolid class="w-4 h-4 mr-2.5" />Info
    </h5>
    <CloseButton on:click={() => (hidden6 = true)} class="mb-4 dark:text-white" />
  </div>
  <p class="mb-6 text-sm text-gray-500 dark:text-gray-400">
    Supercharge your hiring by taking advantage of our <a href="/" class="text-primary-600 underline dark:text-primary-500 hover:no-underline"> limited-time sale </a>
    for Flowbite Docs + Job Board. Unlimited access to over 190K top-ranked candidates and the #1 design job board.
  </p>
  <div class="grid grid-cols-2 gap-4">
    <Button color="light" href="/">Learn more</Button>
    <Button href="/" class="px-4">Get access <ArrowRightOutline class="w-3.5 h-3.5 ml-2" /></Button>
  </div>
</Drawer>