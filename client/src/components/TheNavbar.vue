<template>
 <div>
  <b-navbar toggleable="lg" type="dark" variant="dark">
     <b-navbar-brand href="#">GoJudge</b-navbar-brand>

     <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

     <b-collapse id="nav-collapse" is-nav>
       <b-navbar-nav>
         <b-nav-item to="/home">
          <font-awesome-icon icon="home"/>
          Home
        </b-nav-item>
         <b-nav-item to="/problems">
          <font-awesome-icon icon="question-circle"/>
          Problems
         </b-nav-item>
         <b-nav-item to="/scoreboard">
          <font-awesome-icon icon="list-ol"/>
          Scoreboard
        </b-nav-item>
         <b-nav-item to="/submissions">
          <font-awesome-icon icon="check-double"/>
          Submissions
        </b-nav-item>
       </b-navbar-nav>

       <b-navbar-nav class="ml-auto">
         <b-nav-item to="/login" v-if="!loggedIn" right>
            <font-awesome-icon icon="sign-in-alt"/>
            Login
         </b-nav-item>
         <b-nav-item-dropdown v-if="loggedIn" right>
          <template slot="button-content">
            <font-awesome-icon icon="user"/>
            {{ user.fullname }}
          </template>
          <b-dropdown-item to="/settings">
            <font-awesome-icon icon="user-cog"/>
            Settings
          </b-dropdown-item>
          <b-dropdown-item @click="logout()">
            <font-awesome-icon icon="sign-out-alt"/>
            Logout
          </b-dropdown-item>
        </b-nav-item-dropdown>
       </b-navbar-nav>
     </b-collapse>
   </b-navbar>
 </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      loggedIn: false,
      user: {
        username: null,
        fullname: null,
        access: null
      }
    }
  },
  methods: {
    checkSession() {
      this.$session.subscribe(this.sessionChangeEvent);
      this.$session.whoami();
    },
    sessionChangeEvent() {
      this.loggedIn = this.$session.isLoggedIn();
      if (this.loggedIn) this.user = this.$session.getCachedUser();
    },
    logout() {
      axios.post('/api/auth/logout').then(response => {
        if (!response.data.success) {
          alert('Logout error.');
          return;
        }

        this.$session.logout();
        this.$router.push('/login');
      }).catch(error => {
        alert('Logout request error.');
      });
    }
  },
  mounted() { 
    this.checkSession();
  }
 }
</script>