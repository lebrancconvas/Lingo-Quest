<script lang="ts">
	import { page } from '$app/stores';

	const capitalized = (word: string): string => {
		return word[0].toUpperCase() + word.slice(1, word.length);
	}
	
	const heading = capitalized($page.params.language);

	let isValid: boolean = false; 

	if(heading === "Chinese" || heading === "Japanese" || heading === "Korean") {
		isValid = true;
	}
</script>

<svelte:head>
	<title>Lingo Quest | Stage ({heading})</title>
</svelte:head>

{#if isValid}
<div>
	<header>
		<div>
			<h1>{heading}</h1>
		</div>
	</header>
</div>
{:else}
<div>
	<header>
		<div>
			<h1>Content Not Found</h1>
		</div>
	</header>
</div>
{/if}

<style>
	header {
		text-align: center; 
	}
</style>