<script lang="ts">
	export let language = "";
	export let imageURL = "";
</script>

<main>
	<div id="card-canvas">
		<div id="card">
			<div id="card-image">
				<img src={imageURL} alt={language}> 
			</div>
			<div id="card-text">
				{language}
			</div>
		</div>
	</div>
</main>

<style>

	#card-canvas {
    display: flex;
    flex-wrap: wrap;
    padding: 0 2px;
    margin: 0 16px 16px 16px;
	}	
	#card {
		border: 1px solid black; 
		padding: 10px;
		display: grid; 
		width: 25%;
		height: 380px;  
		margin: 10px auto; 
		box-shadow: 1px 1px 1px 1px black;  
		border-radius: 10px; 
		background-color: white; 
		cursor: pointer; 
	}

	#card:active {
		transform: scale(0.98); 
	}

	#card-image > img {
		width: 256px;
		height: 256px; 
		margin: 5px auto; 
		text-align: center; 
	}

	#card-text {
		font-size: 64px; 
		text-align: center;  
	}

</style>