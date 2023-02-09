<script lang="ts">
	// import { Button } from 'carbon-components-svelte';
</script>

<div>
	<section id="register-form-canvas">
		<section id="register-form">
			<div id="register-header">
				<h2>Join us and Learn together</h2>
			</div>
			<form action="">
				<div>
					<label for="username">Username *</label>
					<input type="text" />
				</div>	
				<div>
					<label for="password">Password *</label>
					<input type="password" />
				</div>	
				<div>
					<label for="confirm-password">Confirm Password *</label>
					<input type="password" />
				</div>	
				<div>
					<label for="display-name">Display Name</label>
					<input type="text" />
				</div>	
			</form>
		</section>
	</section>
</div>

<style>
	#register-form {
		width: 44%;
		height: 600px; 
		border-radius: 10px; 
		background-color: white; 
		margin: 0 auto; 
	}

	#register-header {
		text-align: center; 
	}

	input {
		outline: none; 
	}
</style>