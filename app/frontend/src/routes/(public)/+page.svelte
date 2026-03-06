<script>
	import { resolve } from '$app/paths';
	import mainDesign from '$lib/assets/lp_image.webp';
	import Button from '$lib/components/Button.svelte';
	import PublicHeader from './PublicHeader.svelte';
</script>

<PublicHeader />

<main class="landing">
	<h1>Share your links in one place</h1>
	<section class="hero">
		<figure class="preview">
			<img src={mainDesign} alt="Links page preview" />
		</figure>
		<div class="contents">
			<p class="tagline">
				Linkdev is a link-sharing app for developers. Add your social and project links, customize
				your profile, and get a single page to share—with live preview and drag-and-drop ordering.
			</p>
			<ul class="features">
				<li>
					<strong>Add and manage links</strong> — Add as many links as you want. Each link shows a preview
					and is validated so URLs match the platform (GitHub, Twitter, YouTube, etc.).
				</li>
				<li>
					<strong>Reorder with drag and drop</strong> — Arrange your links in any order. What you see
					in the editor is what visitors see on your profile.
				</li>
				<li>
					<strong>Profile details</strong> — Upload a profile picture and set your first name, last name,
					and email. Your devlinks page looks like a real portfolio.
				</li>
				<li>
					<strong>Preview and share</strong> — See exactly how your page looks before sharing. Copy your
					profile link to your clipboard and use it anywhere—resumes, bios, social.
				</li>
			</ul>
			<Button href={resolve('/register')}>Create Account</Button>
		</div>
	</section>
</main>

<style>
	.landing {
		display: flex;
		flex-direction: column;
	}

	h1 {
		font-size: clamp(2.5rem, 10vw, 3.5rem);
		font-weight: 700;
		margin: var(--spacing-6);
		margin-block-start: var(--spacing-4);
		color: var(--color-text-primary);
	}

	.hero {
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		justify-content: center;
		gap: var(--spacing-6);
		padding-block-end: var(--spacing-6);
		padding-inline: var(--spacing-6);
		margin-inline: auto;
	}

	.contents {
		flex: 1 1 20rem;
	}

	.tagline {
		font-size: var(--font-size-md);
		color: var(--color-text-secondary);
		line-height: 1.6;
		margin-block-end: var(--spacing-5);
	}

	.features {
		list-style: none;
		padding: 0;
		margin: 0 0 var(--spacing-6);
	}

	.features li {
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
		line-height: 1.6;
		margin-block-end: var(--spacing-3);
		padding-inline-start: var(--spacing-4);
		position: relative;
	}

	.features li::before {
		content: '';
		position: absolute;
		left: 0;
		top: 0.5em;
		inline-size: 6px;
		block-size: 6px;
		background-color: var(--color-dark-purple);
		border-radius: var(--radius-full);
	}

	.features strong {
		color: var(--color-text-primary);
	}
	.preview {
		flex: 0 0 auto;
		margin: 0;
	}

	.preview img {
		display: block;
		inline-size: 90vw;
		block-size: auto;
		border-radius: var(--radius-md);
	}

	@media (min-width: 1024px) {
		h1 {
			margin: var(--spacing-10);
			margin-block-start: var(--spacing-6);
		}

		.hero {
			padding-block-end: var(--spacing-10);
			padding-inline: var(--spacing-10);
		}

		.preview img {
			inline-size: 45vw;
			block-size: auto;
		}
	}
</style>
