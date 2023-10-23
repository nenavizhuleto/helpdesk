<script>
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
	} from "flowbite-svelte-icons";
</script>

<div class="flex">
	<Sidebar>
		<SidebarWrapper>
			<SidebarBrand
				site={{ href: "/", name: "", img: "/Logotype.svg" }}
			/>
			<SidebarGroup>
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
	<slot />
</div>
