<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js'
	import * as Card from '$lib/components/ui/card/index.js'
	import { Input } from '$lib/components/ui/input'
	import { cn } from '$lib/utils.js'
	import type { HTMLAttributes } from 'svelte/elements'
	import AppleIcon from '../icons/apple-icon.svelte'
	import GoogleIcon from '../icons/google-icon.svelte'
	import { getUserStore } from '$lib/stores/user.svelte'
	import { goto } from '$app/navigation'
	let { class: className, ...restProps }: HTMLAttributes<HTMLDivElement> = $props()

	const userStore = getUserStore()

	let email = $state('')
	let password = $state('')

	async function onSubmit(event: Event) {
		event.preventDefault()

		const res = await userStore.logIn({ email, password })

		console.log({ res })

		if (!res) {
			// Handle login error (e.g., show a notification)
			console.error('Login failed')
			return
		}

		goto('/')
	}
</script>

<div class={cn('flex flex-col gap-4', className)} {...restProps}>
	<Card.Root class="gap-4">
		<Card.Header class="text-center">
			<Card.Title class="text-xl">Welcome back!</Card.Title>
			<Card.Description>Login with your Apple or Google account</Card.Description>
		</Card.Header>
		<Card.Content>
			<form class="flex flex-col gap-4" onsubmit={onSubmit}>
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
				<div class="grid gap-4">
					<Input
						label="Email"
						type="email"
						placeholder="your-email@example.com"
						required
						bind:value={email}
					/>
					<Input
						label="Password"
						type="password"
						placeholder="*********"
						required
						bind:value={password}
					/>
					<a href="##" class="ml-auto text-sm underline-offset-4 hover:underline -my-2">
						Forgot your password?
					</a>
					<Button type="submit" class="w-full">Login</Button>
				</div>
				<div class="text-center text-sm">
					Don&apos;t have an account?
					<a href="/signup" class="underline underline-offset-4"> Sign up </a>
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
