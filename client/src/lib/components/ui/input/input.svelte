<script lang="ts">
	import type { HTMLInputAttributes, HTMLInputTypeAttribute } from 'svelte/elements'
	import { cn, type WithElementRef } from '$lib/utils.js'

	type InputType = Exclude<HTMLInputTypeAttribute, 'file'>

	type Props = WithElementRef<
		Omit<HTMLInputAttributes, 'type'> &
			({ type: 'file'; files?: FileList } | { type?: InputType; files?: undefined })
	> & {
		label?: string
	}

	let {
		ref = $bindable(null),
		value = $bindable(),
		type,
		files = $bindable(),
		class: className,
		label,
		...restProps
	}: Props = $props()
</script>

{#snippet createFile()}
	{#if type === 'file'}
		<input
			bind:this={ref}
			data-slot="input"
			class={cn(
				'selection:bg-primary dark:bg-input/30 selection:text-primary-foreground border-input ring-offset-background placeholder:text-muted-foreground shadow-xs flex h-9 w-full min-w-0 rounded-md border bg-transparent px-3 pt-1.5 text-sm font-medium outline-none transition-[color,box-shadow] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
				'focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]',
				'aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive',
				className
			)}
			type="file"
			bind:files
			bind:value
			{...restProps}
		/>
	{:else}
		<input
			bind:this={ref}
			data-slot="input"
			class={cn(
				'border-input bg-background selection:bg-primary dark:bg-input/30 selection:text-primary-foreground ring-offset-background placeholder:text-muted-foreground shadow-xs flex h-9 w-full min-w-0 rounded-md border px-3 py-1 text-base outline-none transition-[color,box-shadow] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
				'focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]',
				'aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive',
				className
			)}
			{type}
			bind:value
			{...restProps}
		/>
	{/if}
{/snippet}

{#if label}
	<label>
		<span class="block mb-1 text-sm font-medium text-muted-foreground">
			{label}
		</span>
		{@render createFile()}
	</label>
{:else}
	{@render createFile()}
{/if}
