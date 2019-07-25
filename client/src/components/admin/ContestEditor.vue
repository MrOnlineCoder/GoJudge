<template>
	<div>
		<ErrorBlock :err="error"/>
		<b-card>
			<h4>
				Status:
				<small class="text-danger" v-if="!active">
					Not active
				</small>
				<small class="text-success" v-if="active">
					Active and running
				</small>
			</h4>
			<hr>
			<p v-if="!active">
				Before acitvating the contest, please set contest name, start and end times and problemset in the form below.
			</p>
			<p v-if="active">
				Contest is active now.
			</p>
			<b-button variant="success" @click="setContestActive(true)" v-if="!active">
				<font-awesome-icon icon="rocket"/>
				Activate contest
			</b-button> 
			<b-button variant="danger" @click="setContestActive(false)" v-if="active">
				<font-awesome-icon icon="stop"/>
				Deactivate contest
			</b-button>
		</b-card>
		<hr>
		<b-card>
			<h4>
				Contest
			</h4>
			<small>Latest contest data is loaded.</small>
			<hr>
			<b-form>
				<b-form-group label="Name (title):">
	        <b-form-input v-model="contest.name" placeholder="Enter contest name"/>
	      </b-form-group>
	      <b-form-group label="Contest start time:" description="In format DD.MM.YYYY HH:mm:ss">
	        <b-form-input v-model="contest.start_time" placeholder="Enter start time in format"/>
	        Use preset value for start time:
				  <b-button-group size="sm">
				    <b-button @click="setStartTimePreset(0, 1)">Start in 1 minute</b-button>
				    <b-button @click="setStartTimePreset(0, 10)">Start in 10 minutes</b-button>
				    <b-button @click="setStartTimePreset(0, 30)">Start in 30 minutes</b-button>
				    <b-button @click="setStartTimePreset(1, 0)">Start in 1 hour</b-button>
				    <b-button @click="setStartTimePreset(3, 0)">Start in 3 hours</b-button>
				  </b-button-group>
	      </b-form-group>
	       <b-form-group label="Contest end time:" description="In format DD.MM.YYYY HH:mm:ss">
	        <b-form-input v-model="contest.end_time" placeholder="Enter end time"/>
	        Use preset value for end time:
	        <b-button-group size="sm">
				    <b-button @click="setEndTimePreset(1, 0)">End in 1 hour after start</b-button>
				    <b-button @click="setEndTimePreset(2, 0)">End in 2 hours after start</b-button>
				    <b-button @click="setEndTimePreset(4, 0)">End in 4 hours after start</b-button>
				    <b-button @click="setEndTimePreset(5, 0)">End in 5 hours after start</b-button>
				  </b-button-group>
	      </b-form-group>
	      <b-form-group label="Mode:">
	      	<b-form-radio-group v-model="contest.mode">
	      		<b-form-radio value="time">Time based ranking (ACM ICPC)</b-form-radio>
	      		<b-form-radio value="tests">Test based ranking (allows Partial Solutions)</b-form-radio>
	      		<p v-if="contest.mode === 'time'">
	      			ACM ICPC-based ranking. Scoreboard is filtered by amount of solved problems by a participant. If the amount of solved tasks is equal, sorting by penalty time is used. Each successfully solved problem by a participant recevies penalty time, which equals the time since the contest start. Each failed submission also increases the penalty time by 20 minutes. 

	      			Submission is run through tests until it passes all tests or until the first failed test.
	      		</p>
	      		<p v-if="contest.mode === 'tests'">
	      			Each submission is tested against all test cases. If the solution passed all tests, user receives maximum points for a given problem. However, if solution failed to pass some tests, submission receives status Partial Solution and participant gets points based on amount of passed tests. 
	      		</p>
	      	</b-form-radio-group>
	      </b-form-group>
	      <b-form-group label="Problemset:">
	      	<b-button variant="success" @click="openAddProblemModal()">
						<font-awesome-icon icon="plus"/>
						Add Problem
	      	</b-button>
	      	<br>
	      	<br>
	      	<table class="table table-hover">
	      		<thead>
	      			<th>Problem ID</th>
	      			<th>Shortname</th>
		      		<th>Name</th>
		      		<th>Points</th>
		      		<th>Delete</th>
	      		</thead>
	      		<tbody>
	      			<tr v-for="item,idx in contest.problemset">
	      				<td>{{ item.problem.id }}</td>
	      				<td>{{ idx | problemShortname }} </td>
	      				<td>{{ item.problem.name }}</td>
	      				<td>
	      					<b-form-input v-model.number="item.points"/>
	      				</td>
	      				<td>
	      				<b-button variant="danger" @click="contest.problemset.splice(idx, 1)">
									<font-awesome-icon icon="trash"/>
				      	</b-button>
	      				</td>
	      			</tr>
	      		</tbody>
	      	</table>
	      </b-form-group>

	      <b-button variant="success" @click="saveContest()">
	      	<font-awesome-icon icon="save"/>
	      	Save contest
	      </b-button>
	      <br>
	      <br>
	      <b-button variant="warning" @click="resetContest()">
	      	<font-awesome-icon icon="sync-alt"/>
	      	Reset
	      </b-button>
			</b-form>
		</b-card>
		<br>

		<b-modal
			header-bg-variant="primary"
			title="Add problem to contest problemset"
			ok-only
			ok-title="Close"
			ref="addProblemModal"
			size="lg"
			lazy>
			<b-form-input v-model="addSearch" placeholder="Search problems..."/>
			<table class="table table-hover">
	      <thead>
	      	<th>ID</th>
	      	<th>Name</th>
	      	<th>Timelimit</th>
	      	<th>Memlimit</th>
	      	<th>Select</th>
	      </thead>
	      <tbody>
	      	<tr v-for="p in filteredProblems">
	      		<td>{{ p.id }}</td>
	      		<td>{{ p.name }}</td>
	      		<td>{{ p.timelimit | formatTimelimit}}</td>
	      		<td>{{ p.memlimit | formatMemlimit}}</td>
	      		<td>
	      			<b-button variant="success" @click="selectProblem(p)">
	      				<font-awesome-icon icon="check"/>
	      			</b-button>	
	      		</td>
	      	</tr>
	      </tbody>
	    </table>
		</b-modal>
		<ErrorBlock :err="error"/>
	</div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'

export default {
	data() {
		return {
			active: false,
			contest: {
				name: null,
				start_time: null,
				end_time: null,
				mode: 'time',
				problemset: []
			},
			problems: [],
			addSearch: null,
			error: null,
		}
	},
	methods: {
		openAddProblemModal(){
			this.$refs.addProblemModal.show();
		},
		selectProblem(p) {
			this.$refs.addProblemModal.hide();
			this.contest.problemset.push({
				problem: p,
				points: 0
			});
		},
		setStartTimePreset(hours, mins) {
			this.contest.start_time = moment().add(hours,'hours').add(mins, 'minutes').format('DD.MM.YYYY HH:mm:ss');
		},
		setEndTimePreset(hours, mins) {
			this.contest.end_time = moment(this.contest.start_time, 'DD.MM.YYYY HH:mm:ss').add(hours,'hours').add(mins, 'minutes').format('DD.MM.YYYY HH:mm:ss');
		},
		resetContest() {
			this.contest = {
				name: null,
				start_time: moment().format('DD.MM.YYYY HH:mm:ss'),
				end_time: moment().add(1,'hours').format('DD.MM.YYYY HH:mm:ss'),
				problemset: []
			};
		},
		fetchStatus() {
			axios.get('/api/contest/status').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.active = response.data.active;

			}).catch(error => {
				this.error = error;
			});
		},
		fetchContest() {
			axios.get('/api/admin/contest/load').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.contest = response.data.contest;
				this.contest.start_time = moment(this.contest.start_time).format('DD.MM.YYYY HH:mm:ss');
				this.contest.end_time = moment(this.contest.end_time).format('DD.MM.YYYY HH:mm:ss');
			}).catch(error => {
				this.error = error;
			});
		},
		fetchProblems() {
			axios.get('/api/admin/problems').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.problems = response.data.problems;
				this.fetchStatus();
				this.fetchContest();
			});
		},
		findProblem(id) {
			return this.problems.find(item => { 
				return item.id === id
			});
		},
		setContestActive(mode) {
			axios.post('/api/admin/contest/setActive', {
				active: mode
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.active = mode;
				this.error = null;
			}).catch(error => {
				this.error = error;
			});
		},
		saveContest() {
			if (!this.contest.name || !this.contest.start_time || !this.contest.end_time) {
				this.error = 'You have to fill all contest data before activating it!';
				return;
			}

			if (!this.contest.problemset || this.contest.problemset.length == 0) {
				this.error = 'Your contest must have at least 1 problem in problemset!';
				return;
			}

			let data = Object.assign({}, this.contest);

			//Convert dates to unix timestamps

			data.start_time = moment(data.start_time, 'DD.MM.YYYY HH:mm:ss').valueOf();
			data.end_time = moment(data.end_time, 'DD.MM.YYYY HH:mm:ss').valueOf();

			axios.post('/api/admin/contest/save', {
				contest: data
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.error = null;
			}).catch(error => {
				this.error = error;
			});
		}
	},
	computed: {
		filteredProblems() {
			if (!this.problems || !this.contest || !this.contest.problemset) return [];

			let _problemset = this.contest.problemset.map(item => item.problem);

			let list = this.problems.filter((item) => {
				return !_problemset.includes(item);
			});

			if (!this.addSearch) return list.slice(0,30);

			return this.problems.filter((item) => {
				return item.name.includes(this.addSearch);
			});
		}
	},
	created() {
		this.fetchProblems();
	}
}
</script>