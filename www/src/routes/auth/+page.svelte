<div class=" ">
  <div class="mb-12">
    <h1 class="mb-8 font-bold text-4xl">Вы почти у цели!</h1>
    <p class="text-xl w-[425px]">
      Чтобы начать пользоваться приложением, пожалуйста, пройдите быструю и
      простую процедуру регистрации
    </p>
  </div>
</div>
<form
  hx-post="/register"
  hx-swap="none"
  class="w-[423px] flex shrink-0 flex-col gap-12 bg-white px-12 py-10 rounded-[20px]"
>
  <h2 class="text-2xl font-bold">Регистрация в системе</h2>
  <div>
    <div class="relative w-full focus-within:text-primary mb-8">
      <label
        class="bg-white text-disabled px-1 absolute -top-2 left-3 leading-4"
      >
        Полное Имя
      </label>
      <input
        type="text"
        name="name"
        class="w-full px-4 py-5 border border-disabled rounded-xl outline-none text-black focus:border-primary"
        required
      />
    </div>
    <div class="relative w-full focus-within:text-primary">
      <label
        class="bg-white text-disabled px-1 absolute -top-2 left-3 leading-4"
      >
        Внутренний номер
      </label>
      <input
        type="text"
        name="phone"
        class="w-full px-4 py-5 border border-disabled rounded-xl outline-none text-black focus:border-primary"
        required
      />
    </div>
  </div>
  <button
    class="inline-flex justify-center bg-primary text-white py-5 px-8 rounded-xl items-center gap-4 font-semibold transition-all hover:bg-hover"
  >
    Продолжить
  </button>
</form>
