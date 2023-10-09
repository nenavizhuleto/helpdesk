<script>
	import TextField from "$lib/UI/TextField.svelte";
	import PrimaryButton from "$lib/UI/PrimaryButton.svelte";
</script>

<div class="flex justify-between items-center mt-32">
	<div class="mb-12 flex-shrink max-w-md">
		<h1 class="mb-8 font-bold text-4xl">Вы почти у цели!</h1>
		<p class="text-xl ">
			Чтобы начать пользоваться приложением, пожалуйста, пройдите быструю и
			простую процедуру регистрации
		</p>
	</div>

	<form
		class="w-[423px] flex shrink-0 flex-col gap-12 bg-white px-12 py-10 rounded-[20px]"
	>
		<h2 class="text-2xl font-bold">Регистрация в системе</h2>
		<div>
			<TextField label="Полное Имя" type="text" required name="name" />
			<TextField label="Внутренний номер" type="text" required name="phone" />
		</div>
		<PrimaryButton>Продолжить</PrimaryButton>
	</form>
</div>
