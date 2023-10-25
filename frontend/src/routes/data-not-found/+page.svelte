<script lang="ts">
    import type { PageData } from './$types';
    
    export let data: PageData;
</script>



<div class="h-screen w-full flex items-center justify-center">
  <div class="flex">  
    <div class="max-w-2xl mr-24">
      <div class="w-80 text-black text-4xl font-bold leading-10 mb-8">Не можем найти ваши данные</div>
      <div class="w-96 text-xl">
        Пожалуйста, обратитесь в техническую поддержку по телефону
        <span class="text-indigo-700  text-xl font-semibold">+7 (351) 2002 123</span></div>
    </div>
    <img src="NoDataAsset.svg" alt="">
  </div>
</div>