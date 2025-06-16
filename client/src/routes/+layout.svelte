<script lang="ts">
	import Navbar from '$lib/components/common/navbar.svelte'
	import '../app.css'
	import { ModeWatcher } from 'mode-watcher'

	let { data, children } = $props()
</script>

<ModeWatcher />

<Navbar user={data.user} />

{@render children()}
