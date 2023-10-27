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

export interface Network {
	netmask: string
}

export interface User {
	id: string,
	name: string,
	phone: string,
	network: Network,
	branch: Branch,
	company: Company,
	devices: Device[],
}

export type DeviceType = "PC" | "Unknown"

export interface Device {
	ip: string,
	type: DeviceType
	owner_id: string,
}

export interface Comment {
	id: string,
	content: string,
	user: User | undefined,
	direction: "to" | "from",
	timeCreated: Date,
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
