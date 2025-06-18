<script lang="ts">
	import Navbar from '$lib/components/navbar.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';

	let { children } = $props();
</script>

<ModeWatcher />

<Navbar />

{@render children()}
