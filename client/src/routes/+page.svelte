<script lang="ts">
	import type { ButtonSize } from '$lib/components/ui/button/button.svelte';
	import { Button, type ButtonVariant } from '$lib/components/ui/button/index.js';

	const variants = ['default', 'ghost', 'outline', 'secondary', 'link'] as ButtonVariant[];
	const sizes = ['sm', 'md', 'lg', 'icon'] as ButtonSize[];
</script>

<Button variant="ghost">Tests</Button>

<div class="grid grid-cols-8 justify-items-center text-center gap-4 bg-foreground/20 p-4 max-w-4xl">
	{#each variants as variant (variant)}
		{#each sizes as size (size)}
			<Button {variant} {size}>GO</Button>

			<span>
				{variant}
				{size}
			</span>
		{/each}
	{/each}
</div>
