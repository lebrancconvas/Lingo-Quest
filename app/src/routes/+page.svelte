<script>
</script>

<svelte:head>
	<title>Lingo Quest | Learning language can be fun</title> 
</svelte:head>

<div>
	<header>
		<div>
			<h1>Lingo Quest !</h1>
		</div>
		<div>
			<h2>Learning Language can be fun</h2>
		</div>
	</header>
	<section id="action-section">
		<a href="/register">
			<div id="join-button">
				<h3>Join us !</h3> 
			</div>
		</a>
	</section>
</div>

<style>
	header {
		text-align: center; 
		color: white;
	}

	header > div > h1 {
		font-size: 128px;
	}

	header > div > h2 {
		font-size: 50px; 
	}

	#join-button {
		width: 16%;
		height: 25%;
		border-radius: 10px; 
		background-color: #edfc91; 
		margin: 10px auto; 
		cursor: pointer; 
	}

	#join-button > h3 {
		text-align: center; 
		color: #487bff; 
	}

	#join-button:active {
		transform: scale(0.98); 
	}
</style>