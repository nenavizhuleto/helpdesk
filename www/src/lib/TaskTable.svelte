
<div class="flex flex-col" 
    hx-get="/system/tasks" 
    hx-target="#tableHeader" 
    hx-swap="afterend" 
    hx-trigger="task-created from:body, load" 
    hx-indicator="#loader">
  <!-- Table Header -->
<div
  class="py-5 mb-5 border-b border-neutral-300 text-zinc-500 font-semibold justify-start items-start inline-flex transition-all"
  id="tableHeader"
>
  <div class="w-16 self-stretch px-4 justify-start items-center gap-2.5 flex">
    <!-- ID -->
    <div class="leading-tight">№</div>
  </div>
  <div
    class="grow shrink basis-0 self-stretch px-4 justify-between items-center flex"
  >
    <!-- Theme -->
    <div class="grow shrink basis-0 leading-tight">Тема</div>
  </div>
  <div class="w-40 self-stretch px-4 justify-end items-center gap-2.5 flex">
    <!-- Status -->
    <div class="text-base leading-tight">Статус</div>
  </div>
  <div class="w-48 self-stretch px-4 justify-between items-center flex">
    <!-- Date -->
    <div class="grow shrink basis-0 text-right leading-tight">
      Дата постановки
    </div>
  </div>
</div>
<!-- <img class="htmx-indicator" id="loader" src="assets/oval.svg" alt=""> -->
</div>
