<script lang="ts">
	import { onMount } from 'svelte';
	// import { kanjiList } from '../../store/store.js';

	let kanjiList: string[] = [];
	const url = "https://kanjiapi.dev/v1/kanji/grade-1";

	onMount(async() => {
		fetch(url)
			.then(response => response.json())
			.then(data => {
				console.log(data); 
				// kanjiList.set(data);
				kanjiList = data; 
			})
			.catch(error => {
				console.error(error);
				console.log(`[ERROR] Fetching API Data error.`);
			})
	})
</script>

<main>
	<div>
		<ul>
			{#each kanjiList as kanji}
				<li>{kanji}</li> 
			{/each}
		</ul> 
	</div>
</main>

<style>
	li {
		margin: 10px auto; 
		font-size: 28px;
		text-align: center; 
	}
</style>