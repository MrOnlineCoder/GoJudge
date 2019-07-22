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
			<b-button variant="success" @click="saveContest()" v-if="!active">
				<font-awesome-icon icon="rocket"/>
				Activate contest
			</b-button> 
			<b-button variant="danger" v-if="active">
				<font-awesome-icon icon="stop"/>
				Deactivate contest
			</b-button>
		</b-card>
		<hr>
		<b-card>
			<h4>
				Contest
			</h4>
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
		      		<th>Delete</th>
	      		</thead>
	      		<tbody>
	      			<tr v-for="problemID,idx in contest.problemset">
	      				<td>{{ problemID }}</td>
	      				<td>{{ shortnames[idx] }} </td>
	      				<td>{{ findProblem(problemID).name }}</td>
	      				<td>
	      				<b-button variant="danger" @click="contest.problemset.splice(idx, 1)">
									<font-awesome-icon icon="trash"/>
				      	</b-button>
	      				</td>
	      			</tr>
	      		</tbody>
	      	</table>
	      </b-form-group>
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
				problemset: []
			},
			problems: [],
			addSearch: null,
			error: null,
			shortnames: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
		}
	},
	methods: {
		openAddProblemModal(){
			this.$refs.addProblemModal.show();
		},
		selectProblem(p) {
			this.$refs.addProblemModal.hide();
			this.contest.problemset.push(p.id);
			this.$set(this.problemNamesCache, p.id, p);
		},
		setStartTimePreset(hours, mins) {
			this.contest.start_time = moment().add(hours,'hours').add(mins, 'minutes').format('DD.MM.YYYY HH:mm:ss');
		},
		setEndTimePreset(hours, mins) {
			this.contest.end_time = moment(this.contest.start_time, 'DD.MM.YYYY HH:mm:ss').add(hours,'hours').add(mins, 'minutes').format('DD.MM.YYYY HH:mm:ss');
		},
		fetchStatus() {
			axios.get('/api/contest/status').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.active = response.data.active;

				if (this.active) {
					this.contest = response.data.contest;
					this.contest.start_time = moment(this.contest.start_time).format('DD.MM.YYYY HH:mm:ss');
					this.contest.end_time = moment(this.contest.end_time).format('DD.MM.YYYY HH:mm:ss');
				}
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
			});
		},
		findProblem(id) {
			return this.problems.find(item => { 
				return item.id === id
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

			axios.post('/api/admin/contest/activate', {
				contest: data
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.active = true;
				this.error = null;
			}).catch(error => {
				this.error = error;
			});
		}
	},
	computed: {
		filteredProblems() {
			let list = this.problems.filter((item) => {
				return !this.contest.problemset.includes(item.id);
			});

			if (!this.addSearch) return list.slice(0,30);

			return this.problems.filter((item) => {
				return item.name.includes(this.addSearch);
			});
		},
	},
	created() {
		this.fetchProblems();
	}
}
</script>