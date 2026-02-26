<script lang="ts">
	import logo from '$lib/assets/logo.svg';
</script>

<a href="/" class="logo-link">
	<img src={logo} alt="Logo" /><span class="logo-text desktop-only">devlinks</span>
</a>

<style>
	.logo-link {
		display: flex;
		align-items: center;
		gap: var(--spacing-1);
		text-decoration: none;
		color: inherit;
	}

	.logo-text {
		font-size: 1.375rem;
		font-weight: 700;
		letter-spacing: -0.03em;
		color: var(--color-text-primary);
	}
</style>
