<template>
	<div class="problem-container">
		<div class="problem-header">
			<h2>{{ problem.name }}</h2>
			<small>Time Limit: <b>{{ problem.timelimit | formatTimelimit}}</b> </small>
			<br>
			<small>Memory Limit: <b>{{ problem.memlimit | formatMemlimit}}</b> </small>
			<hr>
		</div>
		<ErrorBlock :err="error"/>
		<div class="problem-text">
			<div v-html="renderedProblemText"></div>
		</div>
		<hr>
	</div>
</template>

<script>
import marked from 'marked'
import axios from 'axios'

export default {
	data() {
		return {
			problem: {
				name: "Loading...",
				timelimit: 0,
				memlimit: 0,
				text: 'Loading...'
			},
			error: null
		}
	},
	methods: {
		fetchProblem() {
			axios.get(`/api/contest/problemset/${this.$route.params.idx}`)
				.then(response => {
					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					if (!response.data.active) {
						this.$router.push('/');
						return;
					}

					this.problem = response.data.problem;
				})
				.catch(error => {
					this.error = error;
				})
		}
	},
	computed: {
		renderedProblemText() {
			return this.problem.text ? marked(this.problem.text) : '';
		}
	},
	mounted() {
		this.fetchProblem();
	}
}
</script>

<style scoped>
.problem-container {
	margin-top: 25px;
}

.problem-header {
	text-align: center;
}
</style>