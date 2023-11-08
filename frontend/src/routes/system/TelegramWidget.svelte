<script lang="ts">
  import { getTelegram, newTelegram } from "$lib/api";
  import type { Telegram } from "$lib/api/types";
  import { Button, Modal, Tooltip } from "flowbite-svelte";
  import { FileCopyOutline } from "flowbite-svelte-icons";
  import { onMount } from "svelte";

  let telegram: Telegram | undefined;

  onMount(async () => {
    const response = await getTelegram();
    if (!response.status) {
      telegram = undefined;
      return;
    }

    telegram = response.data;
  });

  async function createTelegram() {
    const res = await newTelegram();
    if (!res.status) {
      return;
    }

    telegram = res.data;
    show = true;
  }

  let show = false;

  function clickToCopy(node, target) {
    async function copyText() {
      let text = target
        ? document.querySelector(target).innerText
        : node.innerText;

      try {
        await navigator.clipboard.writeText(text);

        node.dispatchEvent(
          new CustomEvent("copysuccess", {
            bubbles: true,
          })
        );
      } catch (error) {
        node.dispatchEvent(
          new CustomEvent("copyerror", {
            bubbles: true,
            detail: error,
          })
        );
      }
    }

    node.addEventListener("click", copyText);

    return {
      destroy() {
        node.removeEventListener("click", copyText);
      },
    };
  }
  let showTooltip = false;
</script>

<Modal title="Доступ к телеграм-боту" bind:open={show} size="sm" autoclose>
  <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
    Для получения уведомений по обращениям скопируйте и отправьте этот код
    доступа телеграм-боту:
  </p>
  <div>
    <button
	  
	  on:click|stopPropagation={()=>(showTooltip=true)}
      id="copy"
      class="inline-flex items-center self-start gap-2 cursor-pointer text-lg font-semibold leading-relaxed text-gray-500 dark:text-gray-400"
      use:clickToCopy>{telegram?.pass} 
	  <FileCopyOutline />
  </button>
	{#if !showTooltip}
	<Tooltip trigger="hover" triggeredBy="#copy" placement="right" type="light"
      >Скопировать код</Tooltip>
	{/if}
    <Tooltip trigger="click" triggeredBy="#copy" placement="right" type="light"
      >Скопировано</Tooltip>
  </div>
  <Button color="blue" href="https://t.me/HDNotificator_bot"
    >Перейти в телеграм-бота</Button
  >
</Modal>

<h2 class="text-black text-2xl font-semibold mb-6">Уведомления в Телеграм</h2>

<div class="flex flex-col gap-6">
  {#if telegram}
    <div class="w-96 text-gray-500 text-base font-normal leading-normal">
      У вас уже подключен телеграм-бот для уведомений по обращениям
    </div>
    <div class="flex gap-4">
      <Button color="blue" href="https://t.me/HDNotificator_bot">Перейти в телеграм-бота</Button>
      <Button on:click={() => (show = true)} color="blue" outline>Показать пароль</Button>
    </div>
  {:else}
  <div class="w-96 text-gray-500 text-base font-normal leading-normal">
	Активируйте телеграм-бота и получайте уведомления по обращениям прямо в
	приложение телеграм
  </div>
  <Button class="self-start" on:click={() => createTelegram()} color="blue">Активировать телеграм-бота</Button>
  {/if}
</div>
