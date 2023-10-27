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
		InboxFullOutline,
		InfoCircleOutline,
		UserCircleOutline,
	} from "flowbite-svelte-icons";
	import type { PageData } from "./$types";
	import { writable } from "svelte/store";
	import { setContext } from "svelte";
	import type { Identity } from "$lib/api/types";

	$: activeUrl = $page.url.pathname;
	export let data: PageData;
	const identity = writable<Identity>();
	$: identity.set(data.identity);

	setContext("identity", identity);
</script>

{#if $identity}
	<div class="flex h-full">
		<div class="bg-gray-50">
			<Sidebar asideClass="" {activeUrl}>
				<SidebarWrapper>
					<SidebarBrand
						site={{ href: "/", name: "", img: "/Logotype.svg" }}
					/>
					<SidebarGroup>
						<SidebarItem href="/profile" label="Профиль">
							<svelte:fragment slot="icon">
								<UserCircleOutline />
							</svelte:fragment>
						</SidebarItem>
						<SidebarItem href="/tasks" label="Обращения">
							<svelte:fragment slot="icon">
								<ClipboardListOutline />
							</svelte:fragment>
						</SidebarItem>
						<SidebarItem href="/messages" label="Сообщения">
							<svelte:fragment slot="icon">
								<InboxFullOutline />
							</svelte:fragment>
							<svelte:fragment slot="subtext">
								<span
									class="inline-flex justify-center items-center p-3 ml-3 w-3 h-3 text-sm bg-gray-200 rounded-full"
								>
									1
								</span>
							</svelte:fragment>
						</SidebarItem>
					</SidebarGroup>
				</SidebarWrapper>
			</Sidebar>
		</div>
		<slot />
	</div>
{/if}
<slot />
