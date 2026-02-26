<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import LogoLink from '$lib/components/LogoLink.svelte';

	let { children } = $props();
</script>

<div class="logo">
	<LogoLink />
</div>

<div class="centered h-full">
	<Card shadow={true}>
		<div class="centered card-content">
			{@render children()}
		</div>
	</Card>
</div>

<style>
	.centered {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.h-full {
		height: 100vh;
	}

	.card-content {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-4);
		inline-size: clamp(180px, 90vw, 320px);
	}

	.logo {
		position: absolute;
		top: var(--spacing-9);
		left: var(--spacing-8);
	}

	@media (max-width: 480px) {
		.logo {
			top: var(--spacing-5);
			left: var(--spacing-4);
		}
	}
</style>
