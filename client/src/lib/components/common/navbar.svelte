<script lang="ts">
	import SunIcon from '@lucide/svelte/icons/sun'
	import MoonIcon from '@lucide/svelte/icons/moon'
	import GearIcon from '@lucide/svelte/icons/git-compare-arrows'

	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import { resetMode, setMode } from 'mode-watcher'
	import { buttonVariants } from '$lib/components/ui/button/index.js'

	import BadgeCheckIcon from '@lucide/svelte/icons/badge-check'
	import BellIcon from '@lucide/svelte/icons/bell'
	import CreditCardIcon from '@lucide/svelte/icons/credit-card'
	import LogOutIcon from '@lucide/svelte/icons/log-out'
	import SparklesIcon from '@lucide/svelte/icons/sparkles'
	import type { User } from '$lib/types'
	import * as Avatar from '../ui/avatar'
	import { enhance } from '$app/forms'

	interface Props {
		user: User | null
	}

	let formRef = $state<HTMLFormElement | null>(null)
	let { user }: Props = $props()

	function handleLogOut() {
		if (formRef) {
			formRef.submit()
		}
	}
</script>

{#if user}
	<nav
		class="h-16 border-b border-b-border bg-background flex justify-between w-full items-center px-4"
	>
		<span>Effective Octo Memory</span>

		<div class="flex gap-4 items-center">
			<DropdownMenu.Root>
				<DropdownMenu.Trigger
					class={`${buttonVariants({ variant: 'outline', size: 'icon' })} ml-auto`}
				>
					<SunIcon
						class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 !transition-all dark:-rotate-90 dark:scale-0"
					/>
					<MoonIcon
						class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 !transition-all dark:rotate-0 dark:scale-100"
					/>
					<span class="sr-only">Toggle theme</span>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content align="end">
					<DropdownMenu.Item class="gap-2" onclick={() => setMode('light')}>
						<SunIcon class="size-4 text-foreground" />
						Light
					</DropdownMenu.Item>
					<DropdownMenu.Item class="gap-2" onclick={() => setMode('dark')}>
						<MoonIcon class="size-4 text-foreground" />
						Dark
					</DropdownMenu.Item>
					<DropdownMenu.Item onclick={() => resetMode()}>
						<GearIcon class="size-4 text-foreground" />
						System
					</DropdownMenu.Item>
				</DropdownMenu.Content>
			</DropdownMenu.Root>

			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="cursor-pointer">
					<Avatar.Root class="size-8 rounded-lg">
						<Avatar.Image
							src="https://shadcn-svelte.com/avatars/shadcn.jpg"
							alt={`${user.firstName} ${user.lastName}`}
						/>
						<Avatar.Fallback class="rounded-lg">
							{user.firstName[0]}{user.lastName[0]}
						</Avatar.Fallback>
					</Avatar.Root>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content
					class="w-(--bits-dropdown-menu-anchor-width) min-w-56 rounded-lg"
					side="bottom"
					align="end"
					sideOffset={4}
				>
					<DropdownMenu.Label class="p-0 font-normal">
						<div class="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
							<Avatar.Root class="size-8 rounded-lg">
								<Avatar.Image
									src="https://shadcn-svelte.com/avatars/shadcn.jpg"
									alt={`${user.firstName} ${user.lastName}`}
								/>
								<Avatar.Fallback class="rounded-lg">CN</Avatar.Fallback>
							</Avatar.Root>
							<div class="grid flex-1 text-left text-sm leading-tight">
								<span class="truncate font-medium">{`${user.firstName} ${user.lastName}`}</span>
								<span class="truncate text-xs">{user.email}</span>
							</div>
						</div>
					</DropdownMenu.Label>
					<DropdownMenu.Separator />
					<DropdownMenu.Group>
						<DropdownMenu.Item>
							<SparklesIcon />
							Upgrade to Pro
						</DropdownMenu.Item>
					</DropdownMenu.Group>
					<DropdownMenu.Separator />
					<DropdownMenu.Group>
						<DropdownMenu.Item>
							<BadgeCheckIcon />
							Account
						</DropdownMenu.Item>
						<DropdownMenu.Item>
							<CreditCardIcon />
							Billing
						</DropdownMenu.Item>
						<DropdownMenu.Item>
							<BellIcon />
							Notifications
						</DropdownMenu.Item>
					</DropdownMenu.Group>
					<DropdownMenu.Separator />
					<form bind:this={formRef} action="/auth?/logOut" method="post" use:enhance>
						<DropdownMenu.Item onclick={handleLogOut}>
							<LogOutIcon />
							Log out
						</DropdownMenu.Item>
					</form>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
	</nav>
{/if}
