<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js'
	import * as Card from '$lib/components/ui/card/index.js'
	import { Input } from '$lib/components/ui/input'
	import { cn } from '$lib/utils.js'
	import AppleIcon from '../icons/apple-icon.svelte'
	import GoogleIcon from '../icons/google-icon.svelte'
	import { enhance } from '$app/forms'
	import type { HTMLAttributes } from 'svelte/elements'
	import type { ActionData } from '../../../routes/signup/$types'

	type Props = HTMLAttributes<HTMLDivElement> & {
		form: ActionData
	}

	let { class: className, form, ...restProps }: Props = $props()

	$effect(() => {
		console.log({ form })
	})
</script>

<div class={cn('flex flex-col gap-4', className)} {...restProps}>
	<Card.Root>
		<Card.Header class="text-center">
			<Card.Title class="text-xl">Create an account</Card.Title>
			<Card.Description>Continue with your Apple or Google account</Card.Description>
		</Card.Header>
		<Card.Content>
			<form class="flex flex-col gap-4" action="/signup?" method="POST" use:enhance>
				<div class="flex flex-col gap-4">
					<Button variant="outline" class="w-full">
						<AppleIcon class="size-4" />
						Login with Apple
					</Button>
					<Button variant="outline" class="w-full">
						<GoogleIcon class="size-4" />
						Login with Google
					</Button>
				</div>
				<div
					class="after:border-border relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t"
				>
					<span class="bg-card text-muted-foreground relative z-10 px-2"> Or continue with </span>
				</div>
				<div class="flex flex-col gap-4">
					<div class="flex gap-4">
						<Input name="firstName" label="First Name" type="text" placeholder="John" required />
						<Input name="lastName" label="Last Name" type="terrorext" placeholder="Doe" required />
					</div>
					<Input
						name="email"
						label="Email"
						type="email"
						placeholder="your-email@example.com"
						required
					/>
					<Input
						name="password"
						label="Password"
						type="password"
						placeholder="*********"
						required
					/>
					<Input
						name="confirmPassword"
						label="Confirm Password"
						type="password"
						placeholder="*********"
						required
					/>

					{#if form?.error}
						<div class="text-red-500 text-sm">
							{form?.error || 'An error occurred. Please try again.'}
						</div>
					{/if}

					<Button type="submit" class="w-full">Create account</Button>
				</div>
				<div class="text-center text-sm">
					Already have an account?
					<a href="/login" class="underline underline-offset-4">Log In</a>
				</div>
			</form>
		</Card.Content>
	</Card.Root>
	<div
		class="text-muted-foreground *:[a]:hover:text-primary *:[a]:underline *:[a]:underline-offset-4 text-balance text-center text-xs"
	>
		By clicking continue, you agree to our <a href="##">Terms of Service</a>
		and <a href="##">Privacy Policy</a>.
	</div>
</div>
