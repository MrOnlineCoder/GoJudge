<template>
	<div>
		<h1>Submissions</h1>
		<hr>
		<ErrorBlock :err="error"/>
		<h2 v-if="busy && !error">
			<font-awesome-icon icon="spinner" spin/>
			Loading...
		</h2>
		<table class="table table-striped" v-if="!busy && !error">
				<thead>
					<th>#</th>
					<th>Problem</th>
					<th>Date</th>
					<th>Language</th>
					<th>Status</th>
				</thead>
				<tbody>
					<tr v-for="s in submissions">
						<td>{{ s.id }}</td>
						<td>{{ s.problem_name }}</td>
						<td>{{ s.time | formatDatetime	}}</td>
						<td>{{ s.lang }}</td>
						<td :class="getVerdictClass(s.verdict)">
							{{ s.verdict }}
						</td>
					</tr>
				</tbody>
		</table>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			submissions: [],
			problemset: [],
			error: null,
			busy: true,
			shortnames: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
		}
	},
	methods: {
		setNames() {
			this.submissions.forEach((sub, index) => {
				let problem = this.problemset.find(p => p.id === sub.problem_id);

				if (problem) {
					this.$set(this.submissions[index], 'problem_name', problem.name);
				}
			});
		},
		fetchProblemset() {
			axios.get('/api/contest/problemset').then(response => {
  			this.busy = false;

        if (!response.data.success) {
  				this.error = response.data.message;
  				return;
  			}

  			this.problemset = response.data.problemset;
  			this.setNames();
  		}).catch(error => {
  			this.error = error;
  			this.busy = false;
  		});
		},
		fetchSubmissions() {
			axios.get('/api/contest/submissions').then(response => {
        if (!response.data.success) {
        	this.busy = false;
  				this.error = response.data.message;
  				return;
  			}

  			this.submissions = response.data.submissions;
  			this.fetchProblemset();
  		}).catch(error => {
  			this.error = error;
  			this.busy = false;
  		});
		},
		getVerdictClass(v) {
			const Verdicts = {
				"PENDING": "text-muted",
				"QUEUED": "text-dark",
				"COMPILING": "text-primary",
				"COMPILATION_ERROR": "text-danger",
				"CHECKING": "text-warning",
				"WRONG_ANSWER": "text-danger",
				"PRESENTATION_ERROR": "text-danger",
				"FAIL": "text-danger",
				"OK": "text-success"
			};

			return Verdicts[v];
		}
	},
	mounted() {
		this.fetchSubmissions();
	}
}
</script>