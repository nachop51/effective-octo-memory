<script lang="ts">
	import SunIcon from '@lucide/svelte/icons/sun'
	import MoonIcon from '@lucide/svelte/icons/moon'
	import GearIcon from '@lucide/svelte/icons/git-compare-arrows'

	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js'
	import { resetMode, setMode } from 'mode-watcher'
	import { buttonVariants } from '$lib/components/ui/button/index.js'
</script>

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
	</div>
</nav>
