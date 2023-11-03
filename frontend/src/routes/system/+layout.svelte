<script lang="ts">
	// --- Utils ---
	import { page } from "$app/stores";

	// --- Components ---
	import {
		Sidebar,
		SidebarWrapper,
		SidebarGroup,
		SidebarBrand,
		SidebarItem,
	} from "flowbite-svelte";

	// -- Icons ---
	import {
		ClipboardListOutline,
		UserCircleOutline,
	} from "flowbite-svelte-icons";

	// highlight corresponding SideBarItem in relation of current url
	$: activeUrl = $page.url.pathname;
</script>

<div class="flex h-full">
	<div class="bg-gray-50 border-r border-gray-300 px-3 pt-3">
		<Sidebar asideClass="" {activeUrl}>
			<SidebarWrapper>
				<SidebarBrand
					site={{ href: "/", name: "", img: "/Logotype.svg" }}
				/>
				<SidebarGroup>
					<SidebarItem href="/system" label="Профиль">
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
