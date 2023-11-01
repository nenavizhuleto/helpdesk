<!-- CHAT GOES HERE -->
<script lang="ts">
  import { enhance } from "$app/forms";

  // --- Types ---
  import type { PageData } from "./$types";
  import {
    Button,
    Badge,
    Alert,
    Label,
    Input,
    InputAddon,
    ButtonGroup,
    Checkbox,
  } from "flowbite-svelte";
  import Message from "./message.svelte";
  import Details from "./details.svelte";
  import { scrollToBottom } from "$lib";

  export let data: PageData;
  // TODO: Why we expect that task cannot be undefined?
  let task = data?.task!;
  console.log(task);
  let comments = task.comments;
  let sending = false;
  let message = "";
  let chatNode;
  let inputNode;
</script>

<div class="flex w-full">
  <!-- Chat Container -->
  <div class="grow h-screen flex flex-col">
    <!-- Chat Header -->
    <div class="px-6 py-6 border-b border-gray-300">
      <h2 class="text-lg font-bold text-center">{task.name}</h2>
    </div>

    <!-- Chat Body -->
    <div class="flex flex-col flex-grow w-full bg-white overflow-hidden">
      <!-- Chat Flow -->
      <div bind:this={chatNode} class="flex flex-col flex-grow h-0 p-4 overflow-auto scrollbar-hide">
        <div class="flex flex-col gap-4 mb-6">
          <Alert color="blue">
            <div class="font-medium mb-2">
              Добро пожаловать в чат обращения!
            </div>
            Здесь вы можете уточнять подробности задачи и получать обратную связь
            от специалистов
          </Alert>
          <Alert color="yellow">
            <div class="font-medium mb-2">Важно!</div>
            Старайтесь отправлять только важную и сформулированную информацию по
            теме обращения.
          </Alert>
        </div>
        {#each comments as comment}
          <Message
            content=   {comment.content}
            date=      {comment.timeCreated}
            direction= {comment.direction}
          />
        {/each}
      </div>
      <!-- Chat Input -->
      <form
        method="post"
        action="?/comment"
        class="p-6 border-t border-gray-300"
        use:enhance={() => {

          sending = true;
          return async ({ result, update }) => {
			console.log(result.data)
			comments = [...comments, result.data.comment]
            sending = false;
			scrollToBottom(chatNode)
			message = ""
			inputNode.focus()
          };
        }}
      >
        <ButtonGroup class="w-full">
          <Input defaultClass="hidden" name="task_id" value={task.id} />
          <Input
			let:props
            defaultClass="outline-none w-full"
            id="input-addon"
            size="lg"
            type="text"
            placeholder="Ваше сообщение..."
            name="message"
          >
		  <input type="text" {...props} bind:value={message} bind:this={inputNode}>
		</Input>
          <Button type="submit" color="blue">Отправить</Button>
        </ButtonGroup>
      </form>
    </div>
  </div>
  <Details {task} />
</div>

<style>
  .scrollbar-hide::-webkit-scrollbar {
    display: none;
  }

  /* For IE, Edge and Firefox */
  .scrollbar-hide {
    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
  }
</style>
