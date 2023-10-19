import type * as types from "$lib/api/types"


const Tasks: types.Task[] = []


const mockCompany: types.Company = {
	id: "mock-company-id",
	name: "Mock Company",
	slug: "mock-company"
}

const mockBranch: types.Branch = {
	id: "mock-branch-id",
	name: "Mock Branch",
	description: "This is fake branch",
	company_id: "mock-company-id",
	address: "Mock st. Fake bld.",
	contacts: "branch@mock.com"
}

const mockUser: types.User = {
	id: "mock-user-id",
	name: "Faker Mockovich",
	phone: "666",
	devices: [
		"172.16.222.31"
	],
}

const mockIdentity: types.Identity = {
	ip: "172.16.222.31",
	type: "PC",
	company: mockCompany,
	branch: mockBranch,
	user: mockUser,
}

const TIMEOUT = 500

function mock(data: any, timeout: number) {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			if(!!data) {
				resolve(data)
			} else {
				reject({ message: "error" })
			}
		}, timeout)
	})
}

async function GetIdentity() {
	return mock(mockIdentity, TIMEOUT)
}

async function CreateTask(name: string, subject: string, status: types.TaskStatus = "created"): Promise<types.Task> {
	const task: types.Task = {
		id: crypto.randomUUID(),
		user: mockUser,
		name: name,
		subject: subject,
		branch: mockBranch,
		company: mockCompany,
		status: status,
		comments: [],
		created_at: new Date(),
		activity_at: new Date(),
	}

	Tasks.push(task)
	return mock(task, TIMEOUT) as Promise<types.Task>
}

async function GetTasks(): Promise<types.Task[]> {
	return mock(Tasks, TIMEOUT) as Promise<types.Task[]>
}

export default {
	GetIdentity,
	GetTasks,
	CreateTask,
}




