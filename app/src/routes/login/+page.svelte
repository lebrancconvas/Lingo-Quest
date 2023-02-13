<script lang="ts">
	import axios from 'axios';
	let username: string;
	let password: string;

	const handleLogin = () => {
		if(username === "" || password === "") {
			alert("Please filled in all the field");
			return;
		}

		if(username.length > 20 || username.length <= 0) {
			alert("Invalid Username to login, Please check again.");
		}

		axios.post("http://localhost:8009/api/v1/login", {
			"username": username,
			"password": password 
		}).then(res => {
			alert("Login Success");
			console.log(res);
			window.location.reload(); 
		}).catch(err => {
			alert("Login Unsuccess, Please Check your username and password"); 
			window.location.reload(); 
			console.error(err); 
		})
	};
</script>

<div>
	<section id="login-section">
		<div id="login-canvas">
			<div id="login-form">
				<div id="login-header">
					<h2>Login</h2>
				</div>
				<form>
					<div>
						<label for="username">Username</label>
						<input type="text" name="username" id="username" bind:value={username} required/>
					</div>
					<div>
						<label for="password">Password</label>
						<input type="password" name="password" id="password" bind:value={password} required/>
					</div>
					<div id="action">
						<!-- svelte-ignore a11y-click-events-have-key-events -->
						<button id="submit-button" on:click={handleLogin}>
							<h3>Submit</h3>
						</button>
					</div>
				</form>
			</div>
		</div>
	</section>
</div>

<style>
	#login-form {
		width: 44%;
		height: 360px; 
		border-radius: 10px; 
		background-color: white; 
		margin: 0 auto; 
	}

	#login-header {
		text-align: center; 
	}

	#login-header > h2 {
		font-size: 50px;
	}

	input {
		outline: none; 
		width: 100%;
		border-radius: 10px; 
		height: 50px; 
		margin: 5px; 
		font-size: 30px; 
	}

	form {
		margin: 10px; 
	}

	label {
		font-size: 26px; 
	}	

	#submit-button {
		padding: 10px;
		background-color: #00d053; 
		width: 25%;
		height: 10%;
		border-radius: 10px; 
		margin: 10px auto; 
		cursor: pointer; 
	}

	#submit-button:active {
		transform: scale(0.97); 
	}

	#submit-button > h3 {
		text-align: center; 
	}
</style>