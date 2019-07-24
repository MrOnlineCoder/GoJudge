<template>
  <div>
  	<hr>
  	<ErrorBlock :err="error"/>
  	<h3 v-if="busy">
  		<font-awesome-icon icon="spinner" spin/>
  		Loading...
  	</h3>
	  <b-jumbotron :header-tag="'h2'" header-level="5" v-if="!busy">
	  	<template slot="header">
	  		{{ active ? contest.name : 'Contest Status' }}
	  	</template>

	    <template slot="lead">
	      <p class="text-danger" v-if="!active">
	      	There is currently no active contest.
	      </p>

	      <p class="text-success" v-if="active">
	      	Contest active
	      </p>
	    </template>

	    <div class="contest-info" v-if="active">
	    	<p>
	    		<font-awesome-icon icon="play"/>
	    		Contest start time: <b>{{ contest.start_time | formatDatetime}}</b>
	    	</p>
	    	<p>
	    		<font-awesome-icon icon="award"/>
	    		Contest end time: <b>{{ contest.end_time | formatDatetime}}</b>
	    	</p>
	    	<p v-if="contestTimeState == 0">
	    		<font-awesome-icon icon="hourglass-half"/>
	    		Time remaining: <b>{{ remainingTime }}</b>
	    	</p>
	    	<p v-if="contestTimeState < 0">
	    		<font-awesome-icon icon="hourglass-start"/>
	    		Time till start: <b>{{ remainingTime }}</b>
	    	</p>
	    	<p v-if="contestTimeState > 0">
	    		<font-awesome-icon icon="trophy"/>
	    		Contest finished.
	    	</p>
	    	<router-link to="/problems" v-if="contestTimeState == 0">
			    <b-button variant="primary">
			    	<font-awesome-icon icon="code"/>
			    	Go to problemset
			    </b-button>
		    </router-link>
		    <router-link to="/scoreboard" v-if="contestTimeState > 0">
			    <b-button variant="warning">
			    	<font-awesome-icon icon="trophy"/>
			    	View results
			    </b-button>
		    </router-link>
	    </div>
	  </b-jumbotron>
  </div>
</template>

<script>
import axios from 'axios'

import moment from 'moment'

export default {
  name: 'home',
  data() {
  	return {
  		contest: {
  			name: null,
  			start_time: null,
  			end_time: null,
  			problemset: []
  		},
  		active: false,
  		error: null,
  		busy: true,
  		timerID: -1,
  		remainingTime: '',
  		contestTimeState: -1,
  	}
  },
  methods: {
  	fetchData() {
  		axios.get('/api/contest/status').then(response => {
  			this.busy = false;

				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.active = response.data.active;

				if (this.active) {
					this.contest = response.data.contest;
					
					this.startTimer();
					this.timerTick();
				}
			}).catch(error => {
				console.error(`[Contest] Status request failed: ${error}`);
				this.error = error;
			});
  	},
  	startTimer() {
  		this.timerID = setInterval(this.timerTick, 1000);
	  },
	  timerTick() {
	  	this.contestTimeState = this.computeNewTimeState();

	  	let now = moment();
	  	let start = moment(this.contest.start_time);
	  	let end = moment(this.contest.end_time);

	  	if (this.contestTimeState == 0) {
	  		let diff = moment.duration(end.diff(now)).format();
	  		this.remainingTime = diff;
	  	}

	  	if (this.contestTimeState < 0) {
	  		let diff = moment.duration(start.diff(now)).format();
	  		this.remainingTime = diff;
	  	}
	  },
	  //This method just tells at which position in time we are
  	//it returns 0 if contest is running now, we can send submission
  	//it returns -1 if contest has not started yet.
  	//it returns 1 if contest has already finished
	  computeNewTimeState() {
  		let startTime = this.contest.start_time;
  		let endTime = this.contest.end_time;
  		let now = moment().valueOf();

  		if (now < startTime && now < endTime) return -1;
  		if (now > endTime && now > startTime) return 1;
  		if (now > startTime && now < endTime) return 0;

  		//if we reached this code, contest was set up incorrectly
  		//dumb administrator!
  		return -1;
	  }
  },
  beforeDestroy() {
  	clearInterval(this.timerID);
  },
  mounted() {
  	this.fetchData();
  },
}
</script>
