export interface Company {
	id: string,
	name: string,
	slug: string,
}

export interface Branch {
	id: string,
	name: string,
	description: string,
	address: string,
	contacts: string,
	company_id: string,
}

export interface User {
	id: string,
	name: string,
	phone: string,
	devices: string[],
}

export type DeviceType = "PC" | "Unknown"

export interface Identity {
	ip: string,
	company: Company,
	branch: Branch,
	user: User,
	type: DeviceType
}

export interface Comment {
	id: string,
	content: string,
}

export type TaskStatus = "created" | "assigned" | "accepted" | "done" | "completed" | "rejected" | "cancelled" | "expired" | "delayed" | "template" | "overdue";

export interface Task {
	id: string,
	name: string,
	subject: string,
	status: TaskStatus,
	created_at: Date,
	activity_at: Date,
	company: Company,
	branch: Branch,
	user: User,
	comments: Comment[],
}

export interface APIError {
	type: string,
	body: {
		action: string,
		entity: string,
		errors: string[]
	}
}
