<script lang="ts">
	import Card from '../../components/Card.svelte';
	const languageList = [
		"Chinese",
		"Japanese",
		"Korean"
	];
</script>

<svelte:head>
	<title>Lingo Quest | Select Language</title> 
</svelte:head>

<div>
	<header>
		<div>
			<h1>
				Select Your Language
			</h1>
		</div>
	</header>
	<section id="language-option-section">
		<div>
			{#each languageList as language}
				<Card language={language} />   
			{/each}
		</div>
	</section>
</div>

<style>
	* {
		background-color: #5cc4ff; 
	}
	header {
		text-align: center; 
	}	
</style>