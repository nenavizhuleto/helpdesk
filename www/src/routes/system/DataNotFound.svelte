<script>
    import Button from '$lib/UI/Button.svelte';
</script>

<div class="max-w-6xl mx-auto flex flex-col gap-6 items-center justify-between mt-60 sm:flex-row ">
  <div class="max-w-[373px]">
    <div class="mb-12">
      <h1 class="mb-8 font-bold text-4xl">Не можем найти <br> ваши данные</h1>
      <p class="text-xl">Пожалуйста, обратитесь в техническую поддержку по телефону <a class="text-[#3627E7] font-medium hover:underline" href="tel:73512002123">+7 (351) 2002-123</a></p>
    </div>

    <Button on:click={() => { window.location.href = "/" }}>
      Обновить
    </Button>
    
  </div>
  <img class="flex-shrink-1" src="./img/no-data.svg" alt="#">
</div>
