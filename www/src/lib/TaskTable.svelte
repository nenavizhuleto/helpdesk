<div class="flex-col flex">
	<!-- Table Header -->
	<div
		class="py-5 mb-5 flex items-stretch transition-all border-b border-neutral-300 text-zinc-500 font-medium"
	>
		<div class="w-16 px-4">
			<!-- ID -->
			№
		</div>

		<div class="grow shrink basis-0 px-4 justify-between items-center flex">
			<!-- Theme -->
			Тема
		</div>

		<div class="w-40 px-4 justify-end items-center gap-2.5 flex">
			<!-- Status -->
			Статус
		</div>

		<div class="w-48 px-4 justify-end items-center flex">
			<!-- Date -->
			Дата постановки
		</div>
	</div>

	<!-- Table Row -->
  <slot>    
    <div class="w-[440px] mx-auto py-40">
      <div class="w-[268px] mb-10 mx-auto text-center text-zinc-500 text-xl font-medium">
        У вас ещё нет обращений в техническую поддержку
      </div>
      <img src="img/empty-tickets.svg" alt="" />
    </div>
  </slot>
	
</div>
