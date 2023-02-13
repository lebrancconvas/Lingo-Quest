<script lang="ts">
	import Card from '../../components/Card.svelte';
	const languageList = [
		{
			"title": "Chinese",
			"imageURL": "https://cdn-icons-png.flaticon.com/512/4008/4008919.png"
		},
		{
			"title": "Japanese",
			"imageURL": "https://cdn-icons-png.flaticon.com/512/3054/3054008.png"
		},
		{
			"title": "Korean",
			"imageURL": "https://cdn-icons-png.flaticon.com/512/3054/3054070.png"
		}
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
				<Card language={language.title} imageURL={language.imageURL} />   
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