<script>
	import AuthForm from '../AuthForm.svelte';
</script>

<h1>Register</h1>

<AuthForm buttonText="Register" />

<div class="login-link">
	<span>Already have an account?</span>
	<a href="/login">Login</a>
</div>

<style>
	h1 {
		font-size: var(--font-size-xl);
	}

	.login-link {
		display: flex;
		justify-content: center;
		gap: var(--spacing-2);
		margin-block: var(--spacing-4);
		color: var(--color-text-secondary);
		font-size: var(--font-size-xs);
	}

	.login-link a {
		color: var(--color-dark-purple);
		text-decoration: none;
	}

	.login-link a:hover {
		text-decoration: underline;
	}
</style>
