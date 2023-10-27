import { getDevices } from "$lib"

export async function load() {
	const devices = await getDevices()

	console.log(devices)

	return {
		devices
	}
}
