<script lang="ts">
	export let language = "";
</script>

<main>
	<div id="card">
		{language}
	</div>
</main>

<style>
	#card {
		border: 1px solid black; 
		width: 30%;
		height: 200px;  
		margin: 10px auto; 
		box-shadow: 1px 1px 1px 1px black;  
		cursor: pointer; 
	}

</style>