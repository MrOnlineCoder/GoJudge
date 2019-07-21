<template>
	<div>
		<div class="add-container">
			<b-button variant="success" size="lg" @click="openNewUserDialog()">
				<font-awesome-icon icon="plus"/>
				Create User
			</b-button>
		</div>
		<hr>
		<ErrorBlock :err="error"/>
		<table class="table">
			<thead>
				<th>ID</th>
				<th>Username</th>
				<th>Access Level</th>
				<!--<th>Edit</th>-->
				<th>Delete</th>
			</thead>
			<tbody>
				<tr v-for="user in users" :key="user.id">
					<td>{{user.id}}</td>
					<td>{{user.username}}</td>
					<td :class="accessClases[user.access]">{{ accessLevels[user.access] }}</td>
					<!--<td>
						<b-button variant="warning">
							<font-awesome-icon icon="pencil-alt"/>
						</b-button>
					</td>-->
					<td>
						<b-button variant="danger" @click="openDeleteUserDialog(user.id)">
							<font-awesome-icon icon="trash"/>
						</b-button>
					</td>
				</tr>
			</tbody>
		</table>

		<b-modal
			header-bg-variant="success"
			title="Create new user"
			ref="createUserModal"
			ok-title="Cancel"
			ok-only
			no-close-on-backdrop>
			<ErrorBlock :err="error"/>
			<b-form @submit.prevent="createUser">
	      <b-form-group label="Username:">
	        <b-form-input
	          v-model="userData.username"
	          required
	          placeholder="Enter username for new user"
	        ></b-form-input>
	      </b-form-group>

	      <b-form-group label="Full name:">
	        <b-form-input
	          v-model="userData.fullname"
	          required
	          placeholder="Enter fullname for new user"
	        ></b-form-input>
	      </b-form-group>

	      <b-form-group label="Password:">
	        <b-form-input
	          v-model="userData.password"
	          required
	          placeholder="Enter password for new user. Don't forget to remember it!"
	        ></b-form-input>
	        <br>
	        <b-button variant="info" size="sm" @click="generatePassword()">
	        	<font-awesome-icon icon="lock"/>
	        	Generate random password
	        </b-button>
	      </b-form-group>

	      <b-form-group label="Access level:">
	        <b-form-select v-model="userData.access">
			      <option :value="0">Participant</option>
			      <option :value="1">Jury</option>
			      <option :value="2">Admin</option>
			    </b-form-select>

			    <p>
			    	{{ accessDescriptions[userData.access] }}
			    </p>
	      </b-form-group>

	      <b-button variant="success" type="submit">
	      	<font-awesome-icon icon="save"/>
	      	Create
	      </b-button>
	    </b-form>
		</b-modal>

		<b-modal
			header-bg-variant="danger"
			title="Delete user"
			ok-only
			ok-title="Cancel"
			ref="deleteUserModal"
			no-close-on-backdrop>
			<ErrorBlock :err="error"/>
			<p>Are you sure you want to delete this user?</p>
			<b>This action is irreversible.</b>
			<hr>
			<b-button variant="danger" @click="deleteUser()">
				<font-awesome-icon icon="trash"/>
				Yes, delete this user.
			</b-button>
		</b-modal>
	</div>
</template>

<script>
import axios from 'axios'

	export default {
		data() {
			return {
				users: [],
				error: null,
				accessLevels: {
					0: 'Participant',
					1: 'Jury',
					2: 'Admin'
				},
				accessClases: {
					0: '',
					1: 'text-primary',
					2: 'text-danger'
				},
				accessDescriptions: {
					0: 'Normal contest participant. Can send problem submissions and clarifications.',
					1: 'Jury user. Can see all submissions and answer clarifications.',
					2: 'GoJudge Administrator. Can access admin panel and do pretty much anything they want.'
				},
				userData: {
					username: null,
					fullname: null,
					password: null,
					access: -1
				},
				deleteData: {
					id: -1
				}
			}
		},
		methods: {
			fetchUsers() {
				axios.get('/api/admin/users').then(response => {
					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					this.users = response.data.users;
				});
			},
			openNewUserDialog() {
				this.userData = {
					username: null,
					fullname: null,
					password: null,
					access: 0
				}
				this.$refs.createUserModal.show();
			},
			openDeleteUserDialog(id) {
				this.deleteData.id = id;
				this.$refs.deleteUserModal.show();
			},
			createUser() {
				if (this.userData.password.length < 8) {
					this.error = 'For security purposes, user password must be at least 8 characters in length';
					return;
				}

				this.error = null;

				axios.post('/api/admin/createUser', {
					user: this.userData
				}).then(response => {
					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					this.fetchUsers();

					this.$refs.createUserModal.hide();

					this.error = null;
				}).catch(error => {
					this.error = `Create User Request failed: ${error}`
				});
			},
			deleteUser() {
				this.error = null;

				axios.post('/api/admin/deleteUser', {
					user_id: this.deleteData.id
				}).then(response => {
					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					this.fetchUsers();

					this.$refs.deleteUserModal.hide();

					this.error = null;
				}).catch(error => {
					this.error = `Delete User Request failed: ${error}`
				});
			},
			generatePassword() {
				/*
					SECURITY: This is not the most secure method to generate password.
					But in case when it is needed for a 1-2 day contest in a limited LAN network, it's okay.
				*/
				this.userData.password = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
			}
		},
		mounted() {
			this.fetchUsers();
		}
	}
</script>

<style scoped>
.add-container {
	text-align: center;
}
</style>