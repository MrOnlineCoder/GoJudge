<template>
	<div>
		<h1>Submit solution</h1>
		<hr>
		<ErrorBlock :err="error"/>
		<b-form @submit.prevent="submitSolution()">
			<b-form-group label="Problem:">
				<b-form-select v-model="submitData.problem_index" required>
					<option :value="null" disabled="">-- Choose a problem ---</option>
					<option v-for="p,idx in problemset" :key="p.id" :value="idx">
						{{ `${shortnames[idx]} - ${p.name}` }}
					</option>
				</b-form-select>
			</b-form-group>
			<b-form-group label="Language:">
					<b-form-select v-model="submitData.language" required>
					<option :value="null" disabled="">-- Choose a language ---</option>
					<option :value="'cpp'">C++</option>
				</b-form-select>
			</b-form-group>
			<b-form-group label="Source code:">
				<b-form-textarea
						      v-model="submitData.sourcecode"
						      placeholder="Paste / type your solution here"
						      rows="5"
						      required
				></b-form-textarea>
			</b-form-group>
			<b-button variant="warning" size="lg" type="submit" :disabled="busy">
				<font-awesome-icon icon="paper-plane"/>
				Submit!
			</b-button>
		</b-form>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			submitData: {
				problem_index: null,
				language: null,
				sourcecode: null
			},
			problemset: [],
			busy: true,
			error: null,
			shortnames: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
		}
	},
	methods: {
		fetchData() {
			axios.get('/api/contest/problemset').then(response => {
  			this.busy = false;

        if (!response.data.success) {
  				this.error = response.data.message;
  				return;
  			}

        if (response.data.not_started || !response.data.active) {
          this.error = "Contest has not started yet!";
          return;
        }

  			this.problemset = response.data.problemset;
  			this.checkRoute();
  		}).catch(error => {
  			this.error = error;
  		});
		},
		checkRoute() {
			if (this.$route.params.idx) {
				this.submitData.problem_index = parseInt(this.$route.params.idx);
			}
		},
		submitSolution() {
			this.busy = true;
			axios.post('/api/contest/submit', this.submitData).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					this.busy = false;
					return;
				}

				this.$router.push('/submissions');
			}).catch(error => {
				this.error = error;
				this.busy = false;
			});
		}
	},
	mounted() {
		this.fetchData();
	}
}
</script>