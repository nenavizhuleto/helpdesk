import type { Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as api from '$lib/api'

export const load = (async () => {
	return {
	};
}) satisfies PageServerLoad;


export const actions: Actions = {
	register: async ({ request }) => {
		const err = {
			success: false,
			field: ""
		}
		const data = await request.formData()
		const firstName = data.get("firstName")
		if (!firstName) {
			err.field = "firstName"
			return err
		}
		const lastName = data.get("lastName")
		if (!lastName) {
			err.field = "lastName"
			return err
		}
		const phone = data.get("phone")
		if (!phone) {
			err.field = "phone"
			return err
		}
		const [user, error] = await api.RegisterUser(firstName.toString() + " " + lastName.toString(), phone.toString())
		if (error) {
			return {
				success: false,
				error: error,
			}
		}
		return {
			success: true,
			user: user,
		}
	}
}
