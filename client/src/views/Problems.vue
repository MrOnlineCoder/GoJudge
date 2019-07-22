<template>
  <div>
  	<h1>Problemset</h1>
		<hr>
    <div v-if="busy">
      <h2>
        <font-awesome-icon icon="spinner"/>
        Loading...
      </h2>
    </div>
		<ErrorBlock :err="error"/>
    <b-alert variant="warning" :show="!contestActive && !busy">
      There is no running contest right now.
    </b-alert>
    <b-alert variant="warning" :show="contestActive && !contestStarted && !busy">
      Contest has not started yet, please wait.
    </b-alert>
    <div v-if="!busy && contestActive && contestStarted">
      <div class="problemset-container" v-for="p,idx in problemset" > 
        <b-card  
          :title="`${shortnames[idx]} - ${p.name}`">
          <hr>
          <p>
            <font-awesome-icon icon="stopwatch"/> Time limit:
            <b>{{ p.timelimit | formatTimelimit }}</b> 
          </p>
          <p>
            <font-awesome-icon icon="database"/> Memory limit:
            <b>{{ p.memlimit | formatMemlimit }}</b> 
          </p>
          <hr>
          <router-link :to="'/problem/'+idx">
            <b-button variant="primary">
              <font-awesome-icon icon="file"/>
              View problem {{shortnames[idx]}}
            </b-button>
          </router-link> 

          <router-link :to="'/submit/'+idx">
            <b-button variant="warning" class="submit-btn">
              <font-awesome-icon icon="paper-plane"/>
              Submit solution
            </b-button>
          </router-link> 
        </b-card> 
        <hr>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'problems',
  data() {
  	return {
      busy: true,
  		problemset: [],
  		contestActive: false,
      contestStarted: true,
  		error: null,
      shortnames: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
  	}
  },
  methods: {
  	fetchProblemset() {
  		axios.get('/api/contest/problemset').then(response => {
  			this.busy = false;

        if (!response.data.success) {
  				this.error = response.data.message;
  				return;
  			}

  			this.contestActive = response.data.active;

        if (response.data.not_started) {
          this.contestStarted = false;
          return;
        }

  			if (this.contestActive) {
  				this.problemset = response.data.problemset;
  			}
  		}).catch(error => {
  			this.error = error;
  		});
  	}
  },
  mounted() {
  	this.fetchProblemset();
  }
}
</script>

<style scoped>
.submit-btn {
  margin-left: 5px;
}
</style>