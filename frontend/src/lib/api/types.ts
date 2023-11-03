export interface Token {
	token: string,
	refresh_token: string,
}

export interface User {
	name: string,
	phone: string,
}

export interface Telegram {
	pass: string,
}

export interface Profile {
	name: string,
	phone: string,
	company: {
		name: string,
	},
	branch: {
		name: string,
		description: string,
		address: string,
		contacts: string,
	}
}

export type TaskStatus = "created" | "assigned" | "accepted" | "done" | "completed" | "rejected" | "cancelled" | "expired" | "delayed" | "template" | "overdue";

export interface Task {
	id: string,
	name: string,
	subject: string,
	status: TaskStatus,
	created_at: Date,
	activity_at: Date,
	branch: {
		name: string,
		description: string,
		address: string,
		contacts: string,
	},
	user: {
		name: string,
		phone: string,
	}
	comments?: Comment[],
}

export interface Comment {
	id: string,
	content: string,
	user: User,
	direction: "to" | "from",
	created_at: Date,
}
