<script lang="ts">
	import Navbar from '$lib/components/navbar.svelte'
	import '../app.css'
	import { ModeWatcher } from 'mode-watcher'
	import { initUserStore } from '$lib/stores/user.svelte'
	import { page } from '$app/state'
	import paths from '$lib/utils/paths'

	const user = initUserStore()
	const canSeeContent = user.user != null || page.url.pathname === paths.login()

	let { children } = $props()
</script>

<ModeWatcher />

<Navbar />

{#if user.isLoading || !canSeeContent}
	<p>Loading...</p>
{:else}
	{@render children()}
{/if}
