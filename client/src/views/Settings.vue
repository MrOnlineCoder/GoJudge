<template>
	<div>
		<h1>User settings</h1>
		<hr>
		<ErrorBlock :err="error"/>
		<b-alert variant="success" :show="passChange.ok">
			<font-awesome-icon icon="check"/>
			Your password has been changed!
		</b-alert>
		<b-card title="Change password">
			<b-form @submit.prevent="changePassword()">
				<b-form-group label="Old password:">
					<b-form-input type="password" v-model="passChange.old_password" placeholder="Enter your old password" required/>
				</b-form-group>
				<b-form-group label="New password:">
					<b-form-input type="password" v-model="passChange.new_password" placeholder="Enter your new password" required/>
				</b-form-group>
				<b-form-group label="Repeat new password:">
					<b-form-input type="password" v-model="passChange.repeat" placeholder="Repeat your new password" required/>
				</b-form-group>

				<b-button variant="success" type="submit">
					<font-awesome-icon icon="key"/>
					Change password
				</b-button>
			</b-form>
		</b-card>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			passChange: {
				new_password: null,
				old_password: null,
				ok: false,
				repeat: null,
			},

			error: null
		}
	},
	methods: {
		changePassword() {
			this.error = null;
			this.passChange.ok = false;

			if (this.passChange.new_password !== this.passChange.repeat) {
				this.error = "Passwords mismatch!";
				return;
			}

			if (this.passChange.new_password.length < 8) {
				this.error = "Your new password must be at least 8 characters long!";
				return;
			}

			axios.post('/api/auth/changePassword', this.passChange).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.passChange.ok = true;
				this.error = null;
			}).catch(error => {
				this.error = error;
			});
		}
	},
	mounted() {
		if (!this.$session.isLoggedIn()) {
			this.$router.push('/login');
		}
	}
}
</script>