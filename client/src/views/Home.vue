<template>
  <div>
  	<hr>
  	<ErrorBlock :err="error"/>
  	<h3 v-if="busy">
  		<font-awesome-icon icon="spinner" spin/>
  		Loading...
  	</h3>
	  <b-jumbotron :header-tag="'h2'" header-level="4" v-if="!busy">
	  	<template slot="header">
	  		{{ active ? contest.name : 'Contest Status' }}
	  	</template>

	    <template slot="lead">
	      <p class="text-danger" v-if="!active">
	      	There is currently no active contest.
	      </p>

	      <p class="text-success" v-if="active">
	      	Contest in progress
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
	    	<p>
	    		<font-awesome-icon icon="clock"/>
	    		Time remaining: <b>{{ remainingTime }}</b>
	    	</p>
	    	<router-link to="/problems" v-if="contestStarted">
			    <b-button variant="primary">
			    	<font-awesome-icon icon="code"/>
			    	Go to problemset
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
  		contestStarted: false,
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
	  	let now = moment();
	  	let then = moment(this.contest.end_time);
	  	let start = moment(this.contest.start_time);
	  	this.remainingTime = moment.duration(then.diff(now)).format('HH:mm:ss')

	  	this.contestStarted = now.isAfter(start);
	  },
  },
  beforeDestroy() {
  	clearInterval(this.timerID);
  },
  mounted() {
  	this.fetchData();
  },
}
</script>
