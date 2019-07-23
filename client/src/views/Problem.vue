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
		<div class="problem-examples" v-if="examples.length > 0">
			<p>Examples:</p>
			<b-card v-for="example in examples">
				<b>Input</b>
				<hr>
				<code v-html="example.input.replace('\n', '<br>')"></code>
				<br>
				<br>
				<b>Output</b>
				<hr>
				<code v-html="example.output.replace('\n', '<br>')"></code>
			</b-card>
			<br>
		</div>
		<router-link :to="'/submit/'+this.$route.params.idx">
      <b-button variant="warning" class="submit-btn">
          <font-awesome-icon icon="paper-plane"/>
          Submit solution to this problem
       </b-button>
    </router-link> 
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
			examples: [],
			error: null
		}
	},
	methods: {
		fetchExamples() {
			axios.get(`/api/contest/problemset/${this.$route.params.idx}/examples`)
				.then(response => {
					if (!response.data.success) {
						this.error = response.data.message;
						return;
					}

					this.examples = response.data.examples;
				})
				.catch(error => {
					this.error = error;
			});
		},
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
					this.fetchExamples();
				})
				.catch(error => {
					this.error = error;
			});
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