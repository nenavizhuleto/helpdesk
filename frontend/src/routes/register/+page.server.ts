import type { Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import mock from '$lib/mock';

export const load = (async () => {
	return {
	};
}) satisfies PageServerLoad;


export const actions: Actions = {
	default: async ({ request }) => {
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
		const user = await mock.RegisterUser(firstName.toString(), lastName.toString(), phone.toString())
		return {
			success: true,
			user,
		}
	}
}
