import { getNetworks } from "$lib"

export async function load() {
	const networks = await getNetworks()

	console.log(networks)

	return {
		networks
	}
}
