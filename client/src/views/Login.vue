<template>
	<div>
		<h1>Login</h1>
		<hr>
		<ErrorBlock :err="error"/>
		<b-form @submit.prevent="login()">
      <b-form-group label="Username:">
        <b-form-input
          v-model="loginData.username"
          required
          placeholder="Enter username"
        ></b-form-input>
      </b-form-group>

      <b-form-group label="Password:">
        <b-form-input
          v-model="loginData.password"
          type="password"
          required
          placeholder="Enter password"
        ></b-form-input>
      </b-form-group>

      <b-button type="submit" variant="primary" :disabled="busy">
      	<font-awesome-icon icon="sign-in-alt"/>
	      Login
	    </b-button>
    </b-form>
	</div>
</template>

<script>
import axios from 'axios'

export default {
		data() {
			return {
				loginData: {
					username: null,
					password: null
				},
				error: null,
				busy: false
			}
		},
		methods: {
			login() {
				this.error = null;
				this.busy = true;
				axios.post('/api/auth/login', {
					username: this.loginData.username,
					password: this.loginData.password
				}).then(response => {
					this.busy = false;

					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					this.$session.whoami();
					this.$router.push('/');
				}).catch(error => {
					this.error = error;
					this.busy = false;
				});
			}
		},
		mounted() {
			if (this.$session.isLoggedIn()) {
				this.$router.push('/');
			}
		}
	}
</script>