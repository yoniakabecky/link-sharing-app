<script lang="ts">
	let { children } = $props();
</script>

<header>
	<nav>
		<a href="/">TOP PAGE</a>
		<a href="/links">LINKS PAGE</a>
		<a href="/profile">PROFILE PAGE</a>
		<a href="/preview">PREVIEW PAGE</a>
	</nav>
</header>

<main>
	{@render children()}
</main>
