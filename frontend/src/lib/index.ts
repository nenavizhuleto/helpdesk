// place files you want to import through the `$lib` alias in this folder.

import { writable, type Writable } from "svelte/store";
import type { User } from "./api/types";

export function shiftDate(str) {
	if (str.length == 1) {
		return "0" + str;
	}
	return str;
}

export function formatDate(date) {
	let tm = new Date(date);
	let day = shiftDate(tm.getDate().toString());
	let month = shiftDate((tm.getMonth() + 1).toString());
	let year = tm.getFullYear().toString().slice(-2);
	let hour = shiftDate(tm.getHours().toString());
	let minutes = shiftDate(tm.getMinutes().toString());
	return `${day}.${month}.${year} ${hour}:${minutes}`;
}

export const scrollToBottom = async (node) => {
	node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
};


export const userStore: Writable<User> = writable();
