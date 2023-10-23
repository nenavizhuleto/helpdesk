import type * as types from "$lib/api/types"




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
const Users: types.User[] = [
	mockUser
]

const Tasks: types.Task[] = [
	{
		id: "8412",
		name: "PC Broken",
		subject: "Help me please",
		created_at: new Date(),
		activity_at: new Date(),
		user: mockUser,
		branch: mockBranch,
		company: mockCompany,
		status: "created",
		comments: [],
	},
	{
		id: "8412",
		name: "PC Broken",
		subject: "Help me please",
		created_at: new Date(),
		activity_at: new Date(),
		user: mockUser,
		branch: mockBranch,
		company: mockCompany,
		status: "assigned",
		comments: [],
	},
	{
		id: "8412",
		name: "PC Broken",
		subject: "Help me please",
		created_at: new Date(),
		activity_at: new Date(),
		user: mockUser,
		branch: mockBranch,
		company: mockCompany,
		status: "accepted",
		comments: [],
	},
	{
		id: "8412",
		name: "PC Broken",
		subject: "Help me please",
		created_at: new Date(),
		activity_at: new Date(),
		user: mockUser,
		branch: mockBranch,
		company: mockCompany,
		status: "completed",
		comments: [],
	},
	{
		id: "8412",
		name: "PC Broken",
		subject: "Help me please",
		created_at: new Date(),
		activity_at: new Date(),
		user: mockUser,
		branch: mockBranch,
		company: mockCompany,
		status: "cancelled",
		comments: [],
	}

]


const TIMEOUT = 500

async function RegisterUser(firstName: string, lastName: string, phone: string): Promise<types.User> {
	const user: types.User = {
		id: "1234",
		name: firstName + " " + lastName,
		phone: phone,
		devices: ["127.0.0.1"],
	}

	Users.push(user)
	return mock(user, TIMEOUT)
}

function mock<T>(data: T, timeout: number): Promise<T> {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			if (!!data) {
				resolve(data)
			} else {
				reject({ message: "error" })
			}
		}, timeout)
	})
}

async function GetIdentity(): Promise<types.Identity> {
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
	RegisterUser,
}
