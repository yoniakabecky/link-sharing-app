<script lang="ts">
	import { icons } from './icons';

	export type IconName = keyof typeof icons;

	type Props = {
		name: IconName;
		size?: number;
	} & svelteHTML.IntrinsicElements['svg'];

	let { name, size = 24, ...props }: Props = $props();
</script>

<svg
	xmlns="http://www.w3.org/2000/svg"
	width={size}
	height={size}
	viewBox="0 0 24 24"
	fill="currentColor"
	{...props}
>
	<title>{name}</title>
	<path d={icons[name]}></path>
</svg>
