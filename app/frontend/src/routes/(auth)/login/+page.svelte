<script>
	import AuthForm from '../AuthForm.svelte';
</script>

<h1>Login</h1>

<AuthForm buttonText="Login" />

<div class="register-link">
	<span>Don't have an account?</span>
	<a href="/register">Create an account</a>
</div>

<style>
	h1 {
		font-size: var(--font-size-xl);
	}

	.register-link {
		display: flex;
		justify-content: center;
		gap: var(--spacing-2);
		margin-block: var(--spacing-4);
		color: var(--color-text-secondary);
		font-size: var(--font-size-xs);
	}

	.register-link a {
		color: var(--color-dark-purple);
		text-decoration: none;
	}

	.register-link a:hover {
		text-decoration: underline;
	}
</style>
