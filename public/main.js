"use strict" 

const phoneInputs = document.querySelectorAll('input[data-tel-number]')
console.log(phoneInputs);
const formBtn = document.querySelector(".form__submit");
const successModalMsg = document.getElementById("success-modal");
const modal = document.querySelector(".modal");

function checkValidity(event) {
	const formNode = event.target.form;
	const isValid = formNode.checkValidity();
	if (isValid) {
		formBtn.classList.remove("disabled");
	} else {
		formBtn.classList.add("disabled");
	}
	return isValid;
}

function handleFormSubmit(e) {
  e.preventDefault()
	openModal(modal);
}

for (let form of document.forms) {
	form.addEventListener("submit", handleFormSubmit);
	form.addEventListener("input", checkValidity);
}

function openModal(modalNode) {
  console.log(modalNode);
	if (!modalNode.classList.contains("open")) {
		const closeModalBtn = document.querySelector(".modal__close");
		closeModalBtn.addEventListener("click", closeModal(modalNode));
		closeModalBtn.focus();
		modalNode.classList.add("open");
		document.body.classList.add("_scroll-lock");

		modalNode.addEventListener("click", function (e) {
			if (
				!e.target.closest(".modal__body") ||
				e.target.classList.contains("modal__close")
			) {
				closeModal(e.target.closest(".modal"));
			}
		});
	}
}
function closeModal(modalNode) {
	modalNode.classList.remove("open");
	document.body.classList.remove("_scroll-lock");
}

function getInputNumbersValue(input) {
	return input.value.replace(/\D/g, "");
}

function onPhoneInput(e) {
	let input = e.target;
	let inputNumbersValue = getInputNumbersValue(input);
	let formattedInputValue = "";

	if (!inputNumbersValue) return (input.value = "");

	if (["7", "8", "9"].indexOf(inputNumbersValue[0]) > -1) {
		if (inputNumbersValue[0] == "9")
			return (input.value = `+7 ${inputNumbersValue}`);

		let firstSymbols = inputNumbersValue[0] == "8" ? "8" : "+7";
		formattedInputValue = firstSymbols;

		if (inputNumbersValue.length > 1)
			formattedInputValue += ` (${inputNumbersValue.substring(1, 4)}`;

		if (inputNumbersValue.length >= 5)
			formattedInputValue += `) ${inputNumbersValue.substring(4, 7)}`;

		if (inputNumbersValue.length >= 8)
			formattedInputValue += `-${inputNumbersValue.substring(7, 9)}`;

		if (inputNumbersValue.length >= 10)
			formattedInputValue += `-${inputNumbersValue.substring(9, 11)}`;
	} else {
		formattedInputValue = `+${inputNumbersValue.substring(0, 16)}`;
	}
	input.value = formattedInputValue;
}
for (let i = 0; i < phoneInputs.length; i++) {
	let input = phoneInputs[i];
	input.addEventListener("input", onPhoneInput);
}