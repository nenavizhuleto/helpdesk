<script>
	import TextField from "$lib/UI/TextField.svelte";
	import Button from "$lib/UI/Button.svelte";

    let name = '';
    let phone = '';
    let isSubmittable = false;

    async function handleSubmit() {
        const res = await fetch("/api/register", { method: "POST", headers: { 'Content-Type': "application/json" }, body: JSON.stringify({ "name": name, "phone": phone})})
        if (res.ok) {
            window.location.href = "/system"
        }
    }

    function validate() {

        const validateName = () => {
            let isValid = false

            if (name.length > 2)
                isValid = true
            // validate name
            return isValid
        }
        const validatePhone = () => {
            let isValid = false
            // validate phone
            if (phone.length > 2 && /^\d+$/.test(phone))
                isValid = true

            return isValid
        }

        isSubmittable = validateName() && validatePhone()
    }
</script>

<div class="flex justify-between items-center mt-32 max-w-5xl mx-auto">
	<div class="mb-12 flex-shrink max-w-md">
		<h1 class="mb-8 font-bold text-4xl">Вы почти у цели!</h1>
		<p class="text-xl ">
			Чтобы начать пользоваться приложением, пожалуйста, пройдите быструю и
			простую процедуру регистрации
		</p>
	</div>

	<form
		class="w-[423px] flex shrink-0 flex-col gap-12 bg-white px-12 py-10 rounded-[20px]"
        on:submit={handleSubmit}
	>
		<h2 class="text-2xl font-bold">Регистрация в системе</h2>
		<div>
            <TextField bind:value={name} on:change={validate} label="Полное Имя" type="text" required name="name" />
            <TextField bind:value={phone} on:change={validate} label="Внутренний номер" type="text" required name="phone" />
		</div>
        <Button on:click={handleSubmit} disabled={!isSubmittable}>Продолжить</Button>
	</form>
</div>
