<script lang="ts">
	import { page } from "$app/stores";
	import {
		Sidebar,
		SidebarWrapper,
		SidebarGroup,
		SidebarBrand,
		SidebarItem,
	} from "flowbite-svelte";
	import {
		ClipboardListOutline,
		UserCircleOutline,
	} from "flowbite-svelte-icons";
	import type { PageData } from "./$types";
	import { writable } from "svelte/store";
	import { setContext } from "svelte";
	import type { Identity } from "$lib/api/types";
	import { goto } from "$app/navigation";

	export let data: PageData;
	const identity = writable<Identity>();
	$: () => {
		if (!data.identity) {
			goto("/register");
		} else {
			identity.set(data.identity);
		}
	};

	$: activeUrl = $page.url.pathname;
	setContext("identity", identity);
</script>

<div class="flex h-full">
	<div class="bg-gray-50">
		<Sidebar asideClass="" {activeUrl}>
			<SidebarWrapper>
				<SidebarBrand
					site={{ href: "/", name: "", img: "/Logotype.svg" }}
				/>
				<SidebarGroup>
					<SidebarItem href="/system/profile" label="Профиль">
						<svelte:fragment slot="icon">
							<UserCircleOutline />
						</svelte:fragment>
					</SidebarItem>
					<SidebarItem href="/system/tasks" label="Обращения">
						<svelte:fragment slot="icon">
							<ClipboardListOutline />
						</svelte:fragment>
					</SidebarItem>
				</SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>
	<slot />
</div>
