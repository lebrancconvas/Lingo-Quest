<script lang="ts">
	// import { Button } from 'carbon-components-svelte';
</script>

<div>
	<header>
		<div>
			<h1>Register</h1> 
		</div>
	</header>
	<section id="register-section">
		<form method="POST">
			<label for="Username"></label>
			<input type="text" name="Username" id="username">
			<label for="Password"></label>
			<input type="password" name="Password" id="password"> 
			<div>
				<button>Register</button> 
			</div>
		</form>
	</section>
</div>

<style>
	header {
		text-align: center; 
	}
</style>