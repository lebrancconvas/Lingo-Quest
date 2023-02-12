<script lang="ts">
	import axios from 'axios';
	let username: string;
	let password: string;
	let confirmpassword: string;
	let displayname: string; 

	const handleSignup = () => {
		if(username === "" || password === "") {
			alert("Please filled in all the field.");
			return; 
		}

		if(displayname === "") {
			displayname = username; 
		}

		if(username.length > 20 || username.length <= 0) {
			alert("Username not valid.");
			return;
		}
		
		if(password !== confirmpassword) {
			alert("Please Confirm Password Again."); 
			return; 
		}

		if(displayname.length > 50 || displayname.length <= 0) {
			alert("Display Name not valid.");
			return; 
		}

		axios.post("http://localhost:8009/api/v1/signup", {
			"username": username,
			"password": password,
			"displayname": displayname
		}).then(res => {
			console.log(res);
		}).catch(err => {
			console.error(err);
		})
	}

</script>

<svelte:head>
	<title>Lingo Quest | Register</title>
</svelte:head>

<div>
	<section id="register-form-canvas">
		<section id="register-form">
			<div id="register-header">
				<h2>Join us and Learn together</h2>
			</div>
			<form action="POST">
				<div>
					<label for="username">Username * (Max 20 Characters)</label>
					<br>
					<input type="text" bind:value={username} max=20 required/>
				</div>	
				<div>
					<label for="password">Password *</label>
					<br>
					<input type="password" bind:value={password} required/>
				</div>	
				<div>
					<label for="confirm-password">Confirm Password *</label>
					<br>
					<input type="password" bind:value={confirmpassword} required/>
				</div>	
				<div>
					<label for="display-name">Display Name (Max 50 Characters)</label> 
					<br> 
					<input type="text" max=50 bind:value={displayname} />
				</div>	
				<div id="action">
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<button id="submit-button" on:click={handleSignup}>
						<h3>Submit</h3>
					</button>
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

	#register-header > h2 {
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