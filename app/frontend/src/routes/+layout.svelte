<script lang="ts">
	import { Toaster } from 'svelte-sonner';
	import favicon from '$lib/assets/favicon.svg';
	import './global.css';
	import './normalize.css';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>devlinks</title>
	<meta name="description" content="A platform for sharing your links with the world" />
</svelte:head>

{@render children()}

<Toaster richColors position="bottom-center" />
